package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/viveksahu26/url_shortner/src"
)

func handleShortURL(writer http.ResponseWriter, req *http.Request) {
	// get original URL from GET method by quering
	originalURL := req.URL.Query().Get("longURL")
	fmt.Println("originalURL: ", originalURL)

	// generate random shortURL
	shortURL := src.GenerateShortURL()
	fmt.Println("shortURL: ", shortURL)

	// save short and long URL to file
	src.SaveInFile(shortURL, originalURL)

	host := "www.simplifyurl.com"

	// build Response
	resp := src.BuildURLWithResponse(host, shortURL, originalURL)
	fmt.Println("response: ", resp)

	err := src.RespondWithJSON(writer, 200, resp)
	if err != nil {
		writer.Write([]byte("Failed to respond with JSON"))
	}
}

const addr = "localhost:8080"

func main() {
	fmt.Println("URL Shorten Service Starts ...")

	// multiplexer: It provides seperate server interface for each request.
	serveMux := http.NewServeMux()
	srv := http.Server{
		Handler:      serveMux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// handleShortUrl function mapped to /short-url
	serveMux.HandleFunc("/short-url", handleShortURL)

	// Server Listing on "localhost:8080"
	srv.ListenAndServe()
}
