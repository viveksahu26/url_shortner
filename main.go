package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/viveksahu26/url_shortner/src" // Adjust the import path according to your project structure
)

func main() {
	// Define command line flags
	shortCmd := flag.NewFlagSet("short", flag.ExitOnError)
	longURL := shortCmd.String("url", "", "URL to shorten")

	// Check the command provided
	if len(os.Args) < 2 {
		fmt.Println("expected 'short' subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "short":
		shortCmd.Parse(os.Args[2:])
		if *longURL == "" {
			fmt.Println("Please provide a URL to shorten using --url=\"http://example.com\"")
			os.Exit(1)
		}
		shortURL := src.GenerateShortURL(*longURL)
		fmt.Printf("Your short url: %s\n", shortURL)
	default:
		fmt.Println("expected 'short' subcommand")
		os.Exit(1)
	}
}
