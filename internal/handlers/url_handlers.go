package handlers

import (
	"github.com/go-chi/chi"

	"github.com/radiophysiker/link_shortener/internal/config"
	"github.com/radiophysiker/link_shortener/internal/handlers/url/usecases"
)

type URLHandler struct {
	URLUseCase usecases.URL
	config     *config.Config
}

func NewURLHandler(u usecases.URL, cfg *config.Config) *URLHandler {
	return &URLHandler{
		URLUseCase: u,
		config:     cfg,
	}
}

func (h *URLHandler) RegisterRoutes(r chi.Router) {
	r.Post("/", h.CreateShortURL)
	r.Get("/{id}", h.GetFullURL)
}
