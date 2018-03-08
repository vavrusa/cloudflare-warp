package tunneldns

import (
	"errors"
	"fmt"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

var errUnreachable = errors.New("unreachable backend")

// Upstream is a simplified interface for proxy destination
type Upstream interface {
	Exchange(ctx context.Context, state request.Request) (*dns.Msg, error)
}

// ProxyPlugin is a simplified DNS proxy using a generic upstream interface
type ProxyPlugin struct {
	Upstreams []Upstream
	Next      plugin.Handler
}

// ServeDNS implements interface for CoreDNS plugin
func (p ProxyPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	var reply *dns.Msg
	var backendErr error

	for _, upstream := range p.Upstreams {
		reply, backendErr = upstream.Exchange(ctx, state)
		if backendErr == nil {
			w.WriteMsg(reply)
			return 0, nil
		}
	}

	return dns.RcodeServerFailure, fmt.Errorf("%s: %s", errUnreachable, backendErr)
}

// Name implements interface for CoreDNS plugin
func (p ProxyPlugin) Name() string { return "proxy" }
