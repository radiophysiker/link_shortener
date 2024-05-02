package utils

import (
	"math/rand"
	"strings"
)

const (
	countOfCharacters = 8
	alphabet          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
)

func GetShortRandomString() string {
	countOfAlphabet := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(countOfCharacters)
	for range countOfCharacters {
		encodedBuilder.WriteByte(alphabet[rand.Intn(countOfAlphabet)])
	}

	return encodedBuilder.String()
}
