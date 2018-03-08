package tunneldns

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"gopkg.in/urfave/cli.v2"

	"github.com/cloudflare/cloudflare-warp/metrics"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/cache"
	log "github.com/sirupsen/logrus"
)

var (
	// DefaultHost is the default endpoint for this service.
	DefaultHost = "dns.cloudflare.com."
	// DefaultBootstrapAddress is the fallback list of service addresses
	DefaultBootstrapAddress = []string{
		"2400:cb00:2048:1::6813:c066",
		"104.19.192.102",
	}
)

// Listener is an adapter between CoreDNS server and Warp runnable
type Listener struct {
	server *dnsserver.Server
	wg     sync.WaitGroup
}

// Run implements a foreground runner
func Run(c *cli.Context) error {
	metricsListener, err := net.Listen("tcp", c.String("metrics"))
	if err != nil {
		log.WithError(err).Fatal("Error opening metrics server listener")
	}

	go metrics.ServeMetrics(metricsListener, nil)

	listener, err := CreateListener(c.String("address"), uint16(c.Uint("port")), c.StringSlice("upstream"))
	if err != nil {
		log.Errorf("Failed to create a listener: %v", err)
		return err
	}

	// Try to start the server
	err = listener.Start()
	if err != nil {
		log.WithError(err).Errorf("Failed to start")
		return listener.Stop()
	}

	// Wait for signal
	signals := make(chan os.Signal, 10)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(signals)
	<-signals

	// Shut down server
	err = listener.Stop()
	if err != nil {
		log.WithError(err).Errorf("Failed to stop")
	}
	return err
}

// Create a CoreDNS server plugin from configuration
func createConfig(address string, port uint16, p plugin.Handler) *dnsserver.Config {
	c := &dnsserver.Config{
		Zone:        ".",
		Transport:   "dns",
		ListenHosts: []string{address},
		Port:        strconv.FormatUint(uint64(port), 10),
	}

	c.AddPlugin(func(next plugin.Handler) plugin.Handler { return p })
	return c
}

// Start blocks for serving requests
func (l *Listener) Start() error {
	log.WithField("addr", l.server.Address()).Infof("Starting DNS over HTTPS proxy server")

	// Start UDP listener
	if udp, err := l.server.ListenPacket(); err == nil {
		l.wg.Add(1)
		go func() {
			l.server.ServePacket(udp)
			l.wg.Done()
		}()
	} else {
		return err
	}

	// Start TCP listener
	tcp, err := l.server.Listen()
	if err == nil {
		l.wg.Add(1)
		go func() {
			l.server.Serve(tcp)
			l.wg.Done()
		}()
	}

	return err
}

// Stop signals server shutdown and blocks until completed
func (l *Listener) Stop() error {
	if err := l.server.Stop(); err != nil {
		return err
	}

	l.wg.Wait()
	return nil
}

// CreateListener configures the server and bound sockets
func CreateListener(address string, port uint16, upstreams []string) (*Listener, error) {
	// Build the list of upstreams
	upstreamList := make([]Upstream, 0)
	for _, url := range upstreams {
		log.WithField("url", url).Infof("Adding DNS upstream")
		upstream, err := NewUpstreamHTTPS(url, DefaultBootstrapAddress)
		if err != nil {
			return nil, err
		}
		upstreamList = append(upstreamList, upstream)
	}

	// Create a local cache with HTTPS proxy plugin
	chain := cache.New()
	chain.Next = ProxyPlugin{
		Upstreams: upstreamList,
	}

	// Format an endpoint
	endpoint := fmt.Sprintf("dns://%s:%d", address, port)

	// Create the actual middleware server
	server, err := dnsserver.NewServer(endpoint, []*dnsserver.Config{createConfig(address, port, NewMetricsPlugin(chain))})
	if err != nil {
		return nil, err
	}

	return &Listener{server: server}, nil
}
