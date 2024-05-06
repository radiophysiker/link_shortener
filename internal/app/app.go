package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/radiophysiker/link_shortener/internal/config"
	v1 "github.com/radiophysiker/link_shortener/internal/controller/http/v1"
	"github.com/radiophysiker/link_shortener/internal/handlers"
	"github.com/radiophysiker/link_shortener/internal/repository"
	"github.com/radiophysiker/link_shortener/internal/usecases"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("cannot load config: %w", err)
	}
	urlRepository := repository.NewURLRepository()
	useCasesURLShortener := usecases.NewURLShortener(urlRepository, cfg)
	urlHandler := handlers.NewURLHandler(useCasesURLShortener, cfg)

	router := v1.NewRouter(urlHandler)
	err = http.ListenAndServe(cfg.ServerPort, router)
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("HTTP server has encountered an error: %w", err)
		}
	}
	return nil
}
