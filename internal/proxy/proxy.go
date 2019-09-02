package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type backends struct {
	prev    int
	next    chan *httputil.ReverseProxy
	append  chan *httputil.ReverseProxy
	proxies []*httputil.ReverseProxy
}

// Engine represents the proxy engine.
type Engine struct {
	backends map[string]*backends
}

type Config struct {
	Backends []Backend
}

type Backend struct {
	Host    string
	RawURLs []string
}

// NewEngine creates a new proxy engine based on the received Config.
func NewEngine(c Config) (*Engine, error) {
	engineBackends := make(map[string]*backends)

	for _, backend := range c.Backends {
		for _, rawURL := range backend.RawURLs {
			URL, err := url.ParseRequestURI(rawURL)
			if err != nil {
				return nil, fmt.Errorf("error parsing URI: %v", err)
			}

			reverseProxy := httputil.NewSingleHostReverseProxy(URL)
			if proxy, exists := engineBackends[backend.Host]; exists {
				proxy.append <- reverseProxy
			} else {
				b := &backends{
					next:    make(chan *httputil.ReverseProxy),
					proxies: []*httputil.ReverseProxy{reverseProxy},
					append:  make(chan *httputil.ReverseProxy),
				}

				go b.manage()

				engineBackends[backend.Host] = b
			}
		}
	}

	return &Engine{
		backends: engineBackends,
	}, nil
}

func (b *backends) manage() {
	for {
		var selected *httputil.ReverseProxy

		if len(b.proxies) == 1 {
			selected = b.proxies[0]
		}

		// Round-Robin

		if b.prev == len(b.proxies)-1 {
			b.prev = 0
			selected = b.proxies[b.prev]
		} else {
			b.prev++
			selected = b.proxies[b.prev]
		}

		select {
		case b.next <- selected:
		case proxy := <-b.append:
			b.proxies = append(b.proxies, proxy)

			// Reset Round-Robin index
			b.prev = 0
		}
	}
}

// ProxyFor serves the request via the backends configured for the received host.
// In case host is unknown or has no backends configured it serves a 404 response.
func (p *Engine) ProxyFor(host string, w http.ResponseWriter, r *http.Request) {
	backends, exists := p.backends[host]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, "unknown host")
		return
	}

	proxy := <-backends.next
	if proxy == nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w, "no backends configured")
		return
	}

	proxy.ServeHTTP(w, r)
}
