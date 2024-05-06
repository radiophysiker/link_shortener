package usecases

import (
	"errors"
	"fmt"

	"github.com/radiophysiker/link_shortener/internal/config"
	"github.com/radiophysiker/link_shortener/internal/entity"
	"github.com/radiophysiker/link_shortener/internal/utils"
)

const lenShortenedURL = 6

var (
	ErrURLExists     = errors.New("URL already exists")
	ErrEmptyFullURL  = errors.New("empty full URL")
	ErrEmptyShortURL = errors.New("empty short URL")
	ErrURLNotFound   = errors.New("URL not found")
)

//go:generate mockery --name=URLRepository --output=./mocks --filename=fs.go
type URLRepository interface {
	Save(url entity.URL) error
	GetFullURL(shortURL string) (string, error)
	IsFullURLExists(fullURL string) bool
}

type URLUseCase struct {
	urlRepository URLRepository
	config        *config.Config
}

func NewURLShortener(re URLRepository, cfg *config.Config) *URLUseCase {
	return &URLUseCase{
		urlRepository: re,
		config:        cfg,
	}
}

func (us URLUseCase) CreateShortURL(fullURL string) (string, error) {
	isExists := us.urlRepository.IsFullURLExists(fullURL)
	if isExists {
		return "", ErrURLExists
	}
	shortURL := utils.GetShortRandomString(lenShortenedURL)
	url := entity.URL{
		ShortURL: shortURL,
		FullURL:  fullURL,
	}
	err := us.urlRepository.Save(url)
	if err != nil {
		if errors.Is(err, ErrEmptyFullURL) {
			return "", ErrEmptyFullURL
		}
		return "", fmt.Errorf("failed to save URL: %w", err)
	}
	return shortURL, nil
}

func (us URLUseCase) GetFullURL(shortURL string) (string, error) {
	fullURL, err := us.urlRepository.GetFullURL(shortURL)
	if err != nil {
		if errors.Is(err, ErrEmptyShortURL) {
			return "", ErrEmptyShortURL
		}
		if errors.Is(err, ErrURLNotFound) {
			return "", fmt.Errorf("%w for: %s", ErrURLNotFound, shortURL)
		}
		return "", fmt.Errorf("failed to get full URL: %w", err)
	}
	return fullURL, nil
}
