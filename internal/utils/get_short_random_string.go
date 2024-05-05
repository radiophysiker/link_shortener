package utils

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

func GetShortRandomString(countOfCharacters int) string {
	countOfAlphabet := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(countOfCharacters)
	for range countOfCharacters {
		encodedBuilder.WriteByte(alphabet[rand.Intn(countOfAlphabet)])
	}

	return encodedBuilder.String()
}
