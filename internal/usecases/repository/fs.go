package repository

import "github.com/radiophysiker/link_shortener/internal/entity"

//go:generate mockery --name=URLRepository --output=../mocks --filename=fs.go
type URLRepository interface {
	Save(url entity.URL) error
	GetFullURL(shortURL string) (string, error)
}
