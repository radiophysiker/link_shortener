package v1

import (
	"github.com/go-chi/chi"

	"github.com/radiophysiker/link_shortener/internal/handlers"
)

// NewRouter creates a new router for the v1 API.
func NewRouter(h *handlers.URLHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", h.CreateShortURL)
	r.Get("/{id}", h.GetFullURL)
	return r
}
