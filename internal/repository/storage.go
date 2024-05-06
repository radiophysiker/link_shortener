package repository

import (
	"fmt"

	"github.com/radiophysiker/link_shortener/internal/entity"
	"github.com/radiophysiker/link_shortener/internal/usecases"
)

type (
	ShortURL = string
	FullURL  = string
)

type URLStorage struct {
	urls map[ShortURL]FullURL
}

func NewURLRepository() *URLStorage {
	return &URLStorage{
		urls: make(map[ShortURL]FullURL),
	}
}

func (s URLStorage) IsFullURLExists(fullURL FullURL) bool {
	for _, url := range s.urls {
		if url == fullURL {
			return true
		}
	}
	return false
}

func (s URLStorage) Save(url entity.URL) error {
	fullURL := url.FullURL
	if fullURL == "" {
		return usecases.ErrEmptyFullURL
	}
	s.urls[url.ShortURL] = fullURL
	return nil
}

func (s URLStorage) GetFullURL(shortURL ShortURL) (FullURL, error) {
	if shortURL == "" {
		return "", usecases.ErrEmptyShortURL
	}
	fullURL, exists := s.urls[shortURL]
	if !exists {
		return "", fmt.Errorf("%w for: %s", usecases.ErrURLNotFound, shortURL)
	}
	return fullURL, nil
}
