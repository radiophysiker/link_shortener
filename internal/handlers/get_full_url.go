package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *URLHandler) GetFullURL(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "id")
	fullURL, err := h.URLUseCase.GetFullURL(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("url is not found for " + shortURL))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", fullURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
