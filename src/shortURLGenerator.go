package src

import (
	"fmt"
	"math/rand"
)

type Result struct {
	OriginalURL string `json:"originalURL"`
	ShortURL    string `json:"shortURL"`
}

// Generate Random URL of 10 characters
// First check LongURL is present in the file or not
// If present, then get that existing shortURL corresponditing to that LongURL
// Otherwise generate new ShortURL
func GenerateShortURL(longURL string) string {
	fileName := "url.properties"
	isFilePresent, _ := IsFileExist(fileName)

	if isFilePresent {
		shortAndLongURLKeyValuePair, _, fileContainLongURL := IsLongURLPresentInFile(fileName, longURL)
		if fileContainLongURL {
			// then retrieve ShortURL from there.
			if shorturl, ok := shortAndLongURLKeyValuePair[longURL]; ok {
				return shorturl
			}
		}
	}
	allCharacters := []rune("%$#@!*^+0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomChar := make([]rune, 10)
	for i := range randomChar {
		randomChar[i] = allCharacters[rand.Intn(len(allCharacters))]
	}
	return string(randomChar)
}

// If file exist, then read that file, otherwise return empty:
// 1. loop over it's content.
// 2. Check line by line,
// 3. that short URL is present in that line or not.
// 4. retun.

// Check whether Short URL is present or not.
// If present returns shorturl, else empty string
func CheckWhetherShortURLisPresentORNot(shorturl string) string {
	fmt.Println("Yes, you are inside CheckWhetherShortURLisPresentORNot")

	fileName := "url.properties"
	isFilePresent, _ := IsFileExist(fileName)

	if isFilePresent {
		shortAndLongURLKeyValuePair, _, fileContainShortURL := IsLongURLPresentInFile(fileName, shorturl)
		if fileContainShortURL {
			// then retrieve ShortURL from there.
			if longurl, ok := shortAndLongURLKeyValuePair[shorturl]; ok {
				return longurl
			}
		}
	}

	return ""
}
