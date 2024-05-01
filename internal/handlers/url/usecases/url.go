package usecases

type (
	URL interface {
		CreateShortURL(fullURL string) (string, error)
		GetFullURL(shortURL string) (string, error)
	}
)
