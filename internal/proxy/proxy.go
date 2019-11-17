package proxy

import (
	"net/http"
	"net/http/httputil"
)

type backends struct {
	prev    int
	proxies []*httputil.ReverseProxy
}

// Engine represents the proxy engine.
type Engine struct {
	// Map of Host to backends
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

	// TODO

	return &Engine{
		backends: engineBackends,
	}, nil
}

// ProxyFor serves the request via the backends configured for the received host.
// In case host is unknown or has no backends configured it serves a 404 response.
func (p *Engine) ProxyFor(host string, w http.ResponseWriter, r *http.Request) {
	// TODO
}
