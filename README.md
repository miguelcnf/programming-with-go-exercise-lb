# Programming with Go Load Balancer exercise

Implement the `ProxyFor` method as described in the `internal/proxy/proxy.go` file:

```go
// ProxyFor serves the request via the backends configured for the received host.
// In case host is unknown or has no backends configured it serves a 404 response.
func (p *Engine) ProxyFor(host string, w http.ResponseWriter, r *http.Request) {
	// TODO
}
```

References:

* [net/http](https://golang.org/pkg/net/http/)
* [net/http/httputil #ReverseProxy](https://golang.org/pkg/net/http/httputil/#ReverseProxy)
* [net/url](https://golang.org/pkg/net/url/)
* [encoding/json](https://golang.org/pkg/encoding/json/)