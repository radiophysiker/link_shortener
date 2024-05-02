package app

import (
	"log"
	"net/http"

	"github.com/radiophysiker/link_shortener/internal/config"
	v1 "github.com/radiophysiker/link_shortener/internal/controller/http/v1"
	"github.com/radiophysiker/link_shortener/internal/repository"
	"github.com/radiophysiker/link_shortener/internal/usecases"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config! %s", err)
	}
	urlRepository := repository.NewURLRepository()
	useCasesURLShortener := usecases.NewURLShortener(urlRepository, cfg)

	router := v1.NewRouter(useCasesURLShortener, cfg)
	err = http.ListenAndServe(cfg.GetServerPort(), router)
	if err != nil {
		log.Fatalf("cannot start server! %s", err)
	}
}
