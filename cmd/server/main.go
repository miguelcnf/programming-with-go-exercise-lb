package main

import (
	"log"
	"net/http"

	"load-balancer/internal/proxy"
)

func main() {
	proxyConfig := proxy.Config{Backends: []proxy.Backend{{
		Host:    "test",
		RawURLs: []string{"http://localhost:8001", "http://localhost:8002", "http://localhost:8003"},
	}}}

	// TODO: initialise load-balancer engine

	// The "/" string is a special route that matches all requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: serve request through a backend
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error on http server: %v", err)
	}
}
