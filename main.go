package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/viveksahu26/url_shortner/src"
)

func handleShortURL(writer http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	} else {
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
			writer.Write([]byte("<h1>Failed to respond with JSON</h1>"))
		}
	}
}

func healthCheckUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/health" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method other than GET not Supported...", http.StatusNotFound)
		return
	}

	w.Write([]byte("<h1>Health of Server is UP & Running... !!</h1>"))
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
	serveMux.HandleFunc("/health", healthCheckUp)

	// Server Listing on "localhost:8080"
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Could not serve.")
		os.Exit(1)
	}
}
