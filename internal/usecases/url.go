package usecases

import (
	"github.com/radiophysiker/link_shortener/internal/entity"
	"github.com/radiophysiker/link_shortener/internal/utils"
)

type URLRepository interface {
	Save(url entity.URL) error
	GetFullURL(shortURL string) (string, error)
}

type URLUseCase struct {
	urlRepository URLRepository
}

func NewURLShortener(re URLRepository) *URLUseCase {
	return &URLUseCase{
		urlRepository: re,
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
		return "", err
	}
	return shortURL, nil
}

func (us URLUseCase) GetFullURL(shortURL string) (string, error) {
	return us.urlRepository.GetFullURL(shortURL)
}
