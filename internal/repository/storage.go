package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/radiophysiker/link_shortener/internal/config"
	"github.com/radiophysiker/link_shortener/internal/entity"
	"github.com/radiophysiker/link_shortener/internal/usecases"
)

const filePermission = 0o600

type (
	ShortURL = string
	FullURL  = string
)
type urlRecord struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type URLFileRepository struct {
	urls   map[ShortURL]FullURL
	config *config.Config
}

func NewFileURLRepository(cfg *config.Config) (*URLFileRepository, error) {
	repo := &URLFileRepository{
		urls:   make(map[ShortURL]FullURL),
		config: cfg,
	}

	if err := repo.loadFromFile(); err != nil {
		return nil, err
	}

	return repo, nil
}

// IsShortURLExists checks if the short URL exists in memory.
func (s URLFileRepository) IsShortURLExists(url entity.URL) bool {
	for shortURL := range s.urls {
		if shortURL == url.ShortURL {
			return true
		}
	}
	return false
}

// Save saves the URL in memory.
func (s URLFileRepository) Save(url entity.URL) error {
	fullURL := url.FullURL
	if fullURL == "" {
		return usecases.ErrEmptyFullURL
	}
	if s.IsShortURLExists(url) {
		return fmt.Errorf("%w for: %s", usecases.ErrURLExists, url.ShortURL)
	}
	s.urls[url.ShortURL] = fullURL
	return s.writeRecordToFile(url)
}

// GetFullURL returns the full URL by the short URL.
func (s URLFileRepository) GetFullURL(shortURL ShortURL) (FullURL, error) {
	if shortURL == "" {
		return "", usecases.ErrEmptyShortURL
	}
	fullURL, exists := s.urls[shortURL]
	if !exists {
		return "", fmt.Errorf("%w for: %s", usecases.ErrURLNotFound, shortURL)
	}
	return fullURL, nil
}

func (s URLFileRepository) writeRecordToFile(url entity.URL) error {
	filePath := s.config.FileStoragePath
	if filePath == "" {
		return nil
	}

	record := urlRecord{
		UUID:        uuid.New().String(),
		ShortURL:    url.ShortURL,
		OriginalURL: url.FullURL,
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePermission)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func(f *os.File) {
		errClose := f.Close()
		if errClose != nil {
			errors.Join(err, fmt.Errorf("failed to close file: %w", err))
		}
	}(f)

	if err := json.NewEncoder(f).Encode(&record); err != nil {
		return fmt.Errorf("failed to encode record: %w", err)
	}
	return err
}

func (s URLFileRepository) loadFromFile() error {
	filePath := s.config.FileStoragePath
	if filePath == "" {
		return nil
	}
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, filePermission)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func(f *os.File) {
		errClose := f.Close()
		if errClose != nil {
			errors.Join(err, fmt.Errorf("failed to close file: %w", err))
		}
	}(f)
	dec := json.NewDecoder(f)
	for dec.More() {
		var record urlRecord
		if err := dec.Decode(&record); err != nil {
			return fmt.Errorf("failed to decode record: %w", err)
		}
		s.urls[record.ShortURL] = record.OriginalURL
	}

	return err
}
