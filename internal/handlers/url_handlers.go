package handlers

import (
	"net/http"

	"github.com/radiophysiker/link_shortener/internal/config"
)

type URL interface {
	CreateShortURL(fullURL string) (string, error)
	GetFullURL(shortURL string) (string, error)
}

type Logger interface {
	Fatal(format string, args ...any)
	Error(format string, args ...any)
	Info(format string, args ...any)
	CustomMiddlewareLogger(next http.Handler) http.Handler
}

type URLHandler struct {
	URLUseCase URL
	config     *config.Config
	logger    Logger
}

func NewURLHandler(u URL, cfg *config.Config, log Logger) *URLHandler {
	return &URLHandler{
		URLUseCase: u,
		config:     cfg,
		logger:    log,
	}
}
