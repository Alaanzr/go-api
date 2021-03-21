package healthcheck

import (
	"net/http"

	"github.com/alaanzr/go-api/internals/middleware"
)

type HttpHandler func(string, func(w http.ResponseWriter, r *http.Request))

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	middleware.LogRequestStart(func() {
		w.Write([]byte(`{"status": "healthy", "code": 200}`))
	})
}

func RegisterHandlers(handler HttpHandler) {
	handler("/api/healthcheck", healthcheckHandler)
}
