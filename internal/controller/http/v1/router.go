package v1

import (
	"github.com/go-chi/chi"
	"net/http"

	"github.com/radiophysiker/link_shortener/internal/handlers"
)

type Logger interface {
	Fatal(format string, args ...any)
	Error(format string, args ...any)
	Info(format string, args ...any)
	CustomMiddlewareLogger(next http.Handler) http.Handler
}

// NewRouter creates a new router for the v1 API.
func NewRouter(h *handlers.URLHandler, log Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(log.CustomMiddlewareLogger)
	r.Post("/", h.CreateShortURL)
	r.Get("/{id}", h.GetFullURL)
	r.Post("/api/shorten", h.CreateShortURLWithJSON)
	return r
}
