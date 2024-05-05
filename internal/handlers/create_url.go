package handlers

import (
	"io"
	"net/http"
	"path"
)

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	url := string(body)
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("url is empty"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	shortURL, err := h.URLUseCase.CreateShortURL(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	baseURL := h.config.BaseURL
	shortURLPath := path.Join(baseURL, shortURL)
	_, err = w.Write([]byte(shortURLPath))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
