package handlers

import (
	"github.com/go-chi/chi"

	"github.com/radiophysiker/link_shortener/internal/handlers/url/usecases"
)

type URLHandler struct {
	URLUseCase usecases.URL
}

func NewURLHandler(u usecases.URL) *URLHandler {
	return &URLHandler{
		URLUseCase: u,
	}
}

func (h *URLHandler) RegisterRoutes(r chi.Router) {
	r.Post("/", h.CreateShortURL)
	r.Get("/{id}", h.GetFullURL)
}
