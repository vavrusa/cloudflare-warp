package tunneldns

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

const (
	defaultTimeout = 5 * time.Second
)

// UpstreamHTTPS is the upstream implementation for DNS over HTTPS service
type UpstreamHTTPS struct {
	client    *http.Client
	endpoint  *url.URL
	addresses []string
}

// NewUpstreamHTTPS creates a new DNS over HTTPS upstream from hostname and optional bootstrap addresses
func NewUpstreamHTTPS(endpoint string, bootstrap []string) (Upstream, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	// Update TLS and HTTP client configuration
	tls := &tls.Config{ServerName: u.Hostname()}
	client := &http.Client{
		Timeout:   time.Second * defaultTimeout,
		Transport: &http.Transport{TLSClientConfig: tls},
	}

	return &UpstreamHTTPS{client: client, endpoint: u, addresses: bootstrap}, nil
}

// Exchange provides an implementation for the Upstream interface
func (u *UpstreamHTTPS) Exchange(ctx context.Context, state request.Request) (*dns.Msg, error) {
	queryBuf, err := state.Req.Pack()
	if err != nil {
		return nil, err
	}

	// No content negotiation for now, use DNS wire format
	buf, backendErr := u.exchangeWireformat(queryBuf)
	if backendErr == nil {
		m := &dns.Msg{}
		if err := m.Unpack(buf); err != nil {
			return nil, err
		}

		m.Id = state.Req.Id
		return m, nil
	}

	log.WithError(backendErr).Errorf("Failed to connect to HTTPS backend %q", u.endpoint)
	return nil, backendErr
}

// Perform message exchange with the default UDP wireformat defined in current draft
// https://datatracker.ietf.org/doc/draft-ietf-doh-dns-over-https
func (u *UpstreamHTTPS) exchangeWireformat(msg []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", u.endpoint.String(), bytes.NewBuffer(msg))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/dns-udpwireformat")
	req.Host = u.endpoint.Hostname()

	resp, err := u.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Check response status code
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("returned status code %d", resp.StatusCode)
	}

	// Read wireformat response from the body
	buf, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return buf, nil
}
