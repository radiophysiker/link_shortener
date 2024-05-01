package repository

import (
	"errors"
	"fmt"

	"github.com/radiophysiker/link_shortener/internal/entity"
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

func (s URLStorage) Save(url entity.URL) error {
	fullURL := url.FullURL
	if fullURL == "" {
		return errors.New("empty full URL")
	}
	s.urls[url.ShortURL] = fullURL
	fmt.Println(url.ShortURL, fullURL)
	return nil
}

func (s URLStorage) GetFullURL(shortURL ShortURL) (FullURL, error) {
	if shortURL == "" {
		return "", errors.New("empty short URL")
	}
	fullURL, exists := s.urls[shortURL]
	if !exists {
		return "", errors.New("URL not found for " + shortURL)
	}
	return fullURL, nil
}
