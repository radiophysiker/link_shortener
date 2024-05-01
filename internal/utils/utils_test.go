package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetShortRandomString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "simple",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetShortRandomString()
			assert.Equal(t, 8, len(got))
		})
	}
}
