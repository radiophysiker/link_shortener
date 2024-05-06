package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/radiophysiker/link_shortener/internal/usecases"
)

func (h *URLHandler) GetFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "id")
	fullURL, err := h.URLUseCase.GetFullURL(shortURL)
	if err != nil {
		if errors.Is(err, usecases.ErrEmptyShortURL) {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("short url is empty"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		if errors.Is(err, usecases.ErrURLNotFound) {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("url is not found for " + shortURL))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", fullURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
