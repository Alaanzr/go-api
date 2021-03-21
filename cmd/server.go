package main

import (
	"net/http"
	"time"

	"github.com/alaanzr/go-api/internals/healthcheck"
)

func main() {
	healthcheck.RegisterHandlers(http.HandleFunc)

	server := http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
