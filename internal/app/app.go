package app

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/radiophysiker/link_shortener/internal/config"
	v1 "github.com/radiophysiker/link_shortener/internal/controller/http/v1"
	"github.com/radiophysiker/link_shortener/internal/handlers"
	"github.com/radiophysiker/link_shortener/internal/logger"
	"github.com/radiophysiker/link_shortener/internal/repository"
	"github.com/radiophysiker/link_shortener/internal/usecases"
)

func Run() error {
	l, err := logger.Init()
	if err != nil {
		return fmt.Errorf("cannot initialize logger: %w", err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		l.Errorf("cannot load config: %v", err)
		return fmt.Errorf("cannot load config: %w", err)
	}
	urlRepository, err := repository.NewFileURLRepository(cfg)
	if err != nil {
		l.Errorf("cannot create URL repository: %v", err)
		return fmt.Errorf("cannot create URL repository: %w", err)
	}
	useCasesURLShortener := usecases.NewURLShortener(urlRepository, cfg)
	urlHandler := handlers.NewURLHandler(useCasesURLShortener, cfg, l)

	router := v1.NewRouter(urlHandler, l)
	l.Info("starting server on port %s", cfg.ServerPort)
	err = http.ListenAndServe(cfg.ServerPort, router)
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			l.Errorf("HTTP server has encountered an error: %v", err)
			return fmt.Errorf("HTTP server has encountered an error: %w", err)
		}
	}
	return nil
}
