package short

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/viveksahu26/url_shortner/cmd/url_shortner/cli/save"
)

type Result struct {
	OriginalURL string `json:"originalURL"`
	ShortURL    string `json:"shortURL"`
}

// Generate Random URL of 10 characters
// First check LongURL is present in the file or not
// If present, then get that existing shortURL corresponditing to that LongURL
// Otherwise generate new ShortURL
func GenerateShortURL(ctx context.Context, longURL string, args []string) error {
	fmt.Println("Inside GenerateShortURL function ")
	fmt.Println("longURL: ", longURL)
	fmt.Println("args: ", args)
	fileName := "url.properties"
	isFilePresent, _ := save.IsFileExist(fileName)

	if isFilePresent {
		shortAndLongURLKeyValuePair, _, fileContainLongURL := save.IsLongURLPresentInFile(fileName, longURL)
		if fileContainLongURL {
			// then retrieve ShortURL from there.
			if shorturl, ok := shortAndLongURLKeyValuePair[longURL]; ok {
				fmt.Println("Short URL is:", shorturl)
			}
		}
	}
	allCharacters := []rune("%$#@!*^+0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomChar := make([]rune, 10)
	for i := range randomChar {
		randomChar[i] = allCharacters[rand.Intn(len(allCharacters))]
	}
	fmt.Println("Short URL is:", string(randomChar))
	if string(randomChar) == "" {
		return fmt.Errorf("Could Not convert long URL into Short URL")
	}
	return nil
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
	isFilePresent, _ := save.IsFileExist(fileName)

	if isFilePresent {
		shortAndLongURLKeyValuePair, _, fileContainShortURL := save.IsLongURLPresentInFile(fileName, shorturl)
		if fileContainShortURL {
			// then retrieve ShortURL from there.
			if longurl, ok := shortAndLongURLKeyValuePair[shorturl]; ok {
				return longurl
			}
		}
	}

	return ""
}
