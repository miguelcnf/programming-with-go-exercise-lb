package main

import (
	"log"
	"net/http"
)

func main() {
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
