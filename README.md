# Programming with Go Load Balancer exercise

Implement the `ProxyFor` method as described in the `internal/proxy/proxy.go` file:

```go
// Engine represents the proxy engine.
type Engine struct {
	// TODO
}

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

Testing:

Use the provided docker-compose setup to start 3 HTTP backends in ports `8001`, `8002` and `8003`. 

Run `docker-compose up` to start the backends: 

```
$ docker-compose up
Creating network "load-balancer-backends_default" with the default driver
Pulling backend-1 (miguelcnf/programming-with-go-exercise-lb-backend:latest)...
latest: Pulling from miguelcnf/programming-with-go-exercise-lb-backend
9d48c3bd43c5: Already exists
d3996683d0c7: Pull complete
c4345ad81850: Pull complete
Digest: sha256:946a13839914c2e3d36a109c5230b2c8c647a70514d263ca9716d4362f7d9a97
Status: Downloaded newer image for miguelcnf/programming-with-go-exercise-lb-backend:latest
Creating load-balancer-backends_backend-3_1 ... done
Creating load-balancer-backends_backend-1_1 ... done
Creating load-balancer-backends_backend-2_1 ... done
...
```

Each backend will log a message to stdout when serving a request:

```
Attaching to load-balancer-backends_backend-2_1, load-balancer-backends_backend-3_1, load-balancer-backends_backend-1_1
backend-1_1  | 2019/09/01 15:27:55 serving request
backend-2_1  | 2019/09/01 15:28:01 serving request
backend-3_1  | 2019/09/01 15:28:03 serving request
```
 
And return a `200 OK` response for any path with a body message stating the port it served the request from:

```
$ curl -v 0:8001
* Rebuilt URL to: 0:8001/
*   Trying 0.0.0.0...
* TCP_NODELAY set
* Connected to 0 (127.0.0.1) port 8001 (#0)
> GET / HTTP/1.1
> Host: 0:8001
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Sun, 01 Sep 2019 15:39:23 GMT
< Content-Length: 15
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 0 left intact
served by: 8001
```
