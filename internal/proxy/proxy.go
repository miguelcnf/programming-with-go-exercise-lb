package proxy

import (
	"net/http"
)

// Engine represents the proxy engine.
type Engine struct {
	// TODO
}

// ProxyFor serves the request via the backends configured for the received host.
// In case host is unknown or has no backends configured it serves a 404 response.
func (p *Engine) ProxyFor(host string, w http.ResponseWriter, r *http.Request) {
	// TODO
}
