package main

import (
	"log"
	"net/http"

	"load-balancer/internal/proxy"
)

func main() {
	engine := &proxy.Engine{}

	// The "/" string is a special route that matches all requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO
		engine.ProxyFor()
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error on http server: %v", err)
	}
}
