package main

import (
	"encoding/json"
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

	host := req.Host

	// build Response
	resp := src.BuildURLResponse(host, shortURL, originalURL)
	fmt.Println("response: ", resp)

	// Converting response  JSON form
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte("Failed to generate response"))
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(jsonBytes)
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
