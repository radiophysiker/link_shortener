package app

import (
	"net/http"

	v1 "github.com/radiophysiker/link_shortener/internal/controller/http/v1"
	"github.com/radiophysiker/link_shortener/internal/repository"
	"github.com/radiophysiker/link_shortener/internal/usecases"
)

func Run() {
	urlRepository := repository.NewURLStorage()
	useCasesURLShortener := usecases.NewURLShortener(urlRepository)

	router := v1.NewRouter(useCasesURLShortener)
	http.ListenAndServe("localhost:8080", router)
}
