package handlers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/radiophysiker/link_shortener/internal/usecases"
)

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fullUrl := string(body)
	if fullUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("url is empty"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	shortURL, err := h.URLUseCase.CreateShortURL(fullUrl)
	if err != nil {
		if errors.Is(err, usecases.ErrURLExists) {
			w.WriteHeader(http.StatusConflict)
			_, err := w.Write([]byte("url already exists"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		if errors.Is(err, usecases.ErrEmptyFullURL) {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("url is empty"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	baseURL := h.config.BaseURL
	fmt.Println(baseURL)
	shortURLPath, err := url.JoinPath(baseURL, shortURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(shortURLPath))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
