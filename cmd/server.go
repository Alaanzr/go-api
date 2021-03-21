package main

import (
	"net/http"
	"time"

	"github.com/alaanzr/go-api/internals/healthcheck"
	"github.com/alaanzr/go-api/internals/tweet"
)

func main() {
	healthcheck.RegisterHandlers(http.HandleFunc)
	tweet.RegisterHandlers(http.HandleFunc)

	server := http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
