package main

import (
	"fmt"
	"net/http"

	"github.com/viveksahu26/url_shortner/src"
)

func handleShortURL(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
	// get original URL from GET method by quering
	originalURL := req.URL.Query().Get("longURL")
	fmt.Println("originalURL: ", originalURL)

	// generate random shortURL
	shortURL := src.GenerateShortURL(originalURL)
	fmt.Println("shortURL: ", shortURL)

	// save short and long URL to file
	src.SaveInFile(shortURL, originalURL)

	host := req.Host

	// build Response
	resp := src.BuildURLWithResponse(host, shortURL, originalURL)

	err := src.RespondWithJSON(writer, 200, resp)
	if err != nil {
		writer.Write([]byte("Failed to respond with JSON"))
	}
}

func healthCheckUp(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Health of Server is UP & Running..."))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

const addr = "localhost:8080"

func main() {
	fmt.Println("URL Shorten Service Starts ...")

	// multiplexer: It provides seperate server interface for each request.
	serveMux := http.NewServeMux()
	srv := http.Server{
		Handler: serveMux,
		Addr:    addr,
	}

	// handleShortUrl function mapped to /enterLongURL
	serveMux.HandleFunc("/sort-url", handleShortURL)

	// healthCheckUp function mapped to /health
	http.HandleFunc("/health", healthCheckUp)

	// Server Listing on "localhost:8080"
	srv.ListenAndServe()
}
