package src

import (
	"math/rand"
)

type Result struct {
	OriginalURL string `json:"originalURL"`
	ShortURL    string `json:"shortURL"`
}

func GenerateShortURL() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789%$#@!*^+")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
