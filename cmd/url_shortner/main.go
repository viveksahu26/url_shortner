package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/viveksahu26/url_shortner/cmd/url_shortner/cli/short"
)

func main() {
	// CLI part
	shortCmd := flag.NewFlagSet("short", flag.ExitOnError)
	longURL := shortCmd.String("url", "", "URL to shorten (CLI mode)")

	// Web server part
	serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	port := serverCmd.String("port", "8080", "Port to run the web server on")

	if len(os.Args) < 2 {
		fmt.Println("expected 'short' or 'server' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "short":
		shortCmd.Parse(os.Args[2:])
		if *longURL == "" {
			fmt.Println("Please provide a URL to shorten using --url=\"http://example.com\"")
			os.Exit(1)
		}
		shortURL := short.GenerateShortURL(*longURL)
		fmt.Printf("Your short url: %s\n", shortURL)
	case "server":
		serverCmd.Parse(os.Args[2:])
		fmt.Printf("Starting server on port %s...\n", *port)
		// setupRoutes() // Function to set up web routes
		setupRoutes()
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	default:
		fmt.Println("expected 'short' or 'server' subcommands")
		os.Exit(1)
	}
}

func setupRoutes() {
	http.HandleFunc("/health", short.HealthCheckUp)
	http.HandleFunc("/short-url", short.HandleShortURL)
	http.HandleFunc("/long-url", short.HandleLongURL)
}
