package usecases

import (
	"fmt"

	"github.com/radiophysiker/link_shortener/internal/config"
	"github.com/radiophysiker/link_shortener/internal/entity"
	"github.com/radiophysiker/link_shortener/internal/utils"
)

//go:generate mockery --name=URLRepository --output=./mocks --filename=fs.go
type URLRepository interface {
	Save(url entity.URL) error
	GetFullURL(shortURL string) (string, error)
}

type URLUseCase struct {
	urlRepository URLRepository
	config        *config.Config
}

func NewURLShortener(re URLRepository, config *config.Config) *URLUseCase {
	return &URLUseCase{
		urlRepository: re,
		config:        config,
	}
}

func (us URLUseCase) CreateShortURL(fullURL string) (string, error) {
	shortURL := utils.GetShortRandomString()
	url := entity.URL{
		ShortURL: shortURL,
		FullURL:  fullURL,
	}
	err := us.urlRepository.Save(url)
	if err != nil {
		return "", fmt.Errorf("failed to save URL: %w", err)
	}
	return shortURL, nil
}

func (us URLUseCase) GetFullURL(shortURL string) (string, error) {
	fullURL, err := us.urlRepository.GetFullURL(shortURL)
	if err != nil {
		return "", fmt.Errorf("failed to get full URL: %w", err)
	}
	return fullURL, nil
}
