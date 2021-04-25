package healthcheck

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HttpHandler func(string, func(w http.ResponseWriter, r *http.Request))

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "healthy", "code": 200}`))
}

func RegisterHandlers(router *chi.Mux) {
	router.Get("/api/healthcheck", healthcheck)
}
