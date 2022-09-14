package main

import (
	"fmt"
	"net/http"

	"github.com/viveksahu26/url_shortner/src"
)

func handleShortUrl(writer http.ResponseWriter, req *http.Request) {
	// get original URL from GET method by quering
	originalURL := req.URL.Query().Get("longURL")
	fmt.Println("originalURL: ", originalURL)

	// generate random shortURL
	shortURL := src.GenerateShortURL()
	fmt.Println("shortURL: ", shortURL)

	// save short and long URL to file
	src.SaveToFile(shortURL, originalURL)
}

func main() {
	fmt.Println("URL Shorten Service")

	// handleShortUrl function mapped to /short-url
	http.HandleFunc("/short-url", handleShortUrl)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
