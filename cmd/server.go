package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/alaanzr/go-api/internals/healthcheck"
	"github.com/alaanzr/go-api/internals/tweet"
)

func main() {
	router := chi.NewRouter()

  	router.Use(middleware.Logger)
  	router.Use(middleware.Recoverer)

	healthcheck.RegisterHandlers(router)
	tweet.RegisterHandlers(router)

	log.Fatal(http.ListenAndServe(":9000", router))
}