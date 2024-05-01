package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/radiophysiker/link_shortener/internal/config"
	"github.com/radiophysiker/link_shortener/internal/usecases/mocks"
)

func TestURLUseCase_CreateShortURL(t *testing.T) {
	type args struct {
		fullURL string
	}

	mocksRepoURL := mocks.NewURLRepository(t)
	mocksRepoURL.
		On("Save", mock.AnythingOfType("entity.URL")).Return(nil)
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				fullURL: "https://yandex.ru",
			},
			want:    "short_url",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := URLUseCase{
				urlRepository: mocksRepoURL,
				config: &config.Config{
					BaseURL:    "http://localhost:8080",
					ServerPort: "localhost:8080",
				},
			}
			got, err := us.CreateShortURL(tt.args.fullURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEmpty(t, got)
		})
	}
}

func TestURLUseCase_GetFullURL(t *testing.T) {
	type args struct {
		shortURL string
	}

	mocksRepoURL := mocks.NewURLRepository(t)
	mocksRepoURL.
		On("GetFullURL", mock.AnythingOfType("string")).Return("https://yandex.ru", nil)
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				shortURL: "sdfsd356",
			},
			want:    "https://yandex.ru",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := URLUseCase{
				urlRepository: mocksRepoURL,
				config: &config.Config{
					BaseURL:    "http://localhost:8080",
					ServerPort: "localhost:8080",
				},
			}
			got, err := us.GetFullURL(tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
