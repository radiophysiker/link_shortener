package v1

import (
	"github.com/go-chi/chi"

	"github.com/radiophysiker/link_shortener/internal/handlers"
)

type URL interface {
	CreateShortURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

// NewRouter creates a new router for the v1 API.
func NewRouter(h *handlers.URLHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", h.CreateShortURL)
	r.Get("/{id}", h.GetFullURL)
	return r
}
