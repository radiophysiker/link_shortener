package v1

import (
	"github.com/go-chi/chi"

	"github.com/radiophysiker/link_shortener/internal/config"
	"github.com/radiophysiker/link_shortener/internal/handlers"
)

type URL interface {
	CreateShortURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

// NewRouter creates a new router for the v1 API
func NewRouter(u URL, cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	urlHandler := handlers.NewURLHandler(u, cfg)
	urlHandler.RegisterRoutes(r)
	return r
}
