package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/radiophysiker/link_shortener/internal/utils"
)

type CreateShortURLEntryRequest struct {
	FullURL string `json:"url"`
}

type CreateShortURLEntryResponse struct {
	ShortURL string `json:"result"`
}

func (h *URLHandler) CreateShortURLWithJSON(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("cannot read request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var request CreateShortURLEntryRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.logger.Error("cannot unmarshal request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var fullURL = request.FullURL
	if fullURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("fullURL is empty"))
		if err != nil {
			utils.WriteErrorWithCannotWriteResponse(w, err, h.logger)
		}
		return
	}
	shortURL, err := h.URLUseCase.CreateShortURL(fullURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	baseURL := h.config.BaseURL
	shortURLPath, err := url.JoinPath(baseURL, shortURL)
	if err != nil {
		h.logger.Error("cannot join base URL and short URL: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := CreateShortURLEntryResponse{ShortURL: shortURLPath}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResp)
	if err != nil {
		utils.WriteErrorWithCannotWriteResponse(w, err, h.logger)
	}
}
