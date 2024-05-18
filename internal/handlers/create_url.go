package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/radiophysiker/link_shortener/internal/usecases"
	"github.com/radiophysiker/link_shortener/internal/utils"
)

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.logger.Error("cannot read request body: %v", err)
		return
	}
	fullURL := string(body)
	if fullURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("url is empty"))
		if err != nil {
			utils.WriteErrorWithCannotWriteResponse(w, err, h.logger)
		}
		return
	}
	shortURL, err := h.URLUseCase.CreateShortURL(fullURL)
	if err != nil {
		if errors.Is(err, usecases.ErrURLExists) {
			w.WriteHeader(http.StatusConflict)
			_, err := w.Write([]byte("url already exists"))
			if err != nil {
				utils.WriteErrorWithCannotWriteResponse(w, err, h.logger)
			}
			return
		}
		if errors.Is(err, usecases.ErrEmptyFullURL) {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("url is empty"))
			if err != nil {
				utils.WriteErrorWithCannotWriteResponse(w, err, h.logger)
			}
			return
		}
		h.logger.Error("cannot create short URL: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	baseURL := h.config.BaseURL
	shortURLPath, err := url.JoinPath(baseURL, shortURL)
	if err != nil {
		h.logger.Error("cannot join base URL and short URL: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(shortURLPath))
	if err != nil {
		utils.WriteErrorWithCannotWriteResponse(w, err, h.logger)
	}
}
