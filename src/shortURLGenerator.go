package src

import (
	"math/rand"
)

type Result struct {
	OriginalURL string `json:"originalURL"`
	ShortURL    string `json:"shortURL"`
}

// Generate Random URL of 10 characters
func GenerateShortURL() string {
	allCharacters := []rune("%$#@!*^+0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomChar := make([]rune, 10)
	for i := range randomChar {
		randomChar[i] = allCharacters[rand.Intn(len(allCharacters))]
	}
	return string(randomChar)
}
