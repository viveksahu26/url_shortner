package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/viveksahu26/url_shortner/cmd/url_shortner/cli"
	urlShortnerError "github.com/viveksahu26/url_shortner/cmd/url_shortner/errors"
)

func main() {
	ctx := context.Background()
	for i, arg := range os.Args {
		if (strings.HasPrefix(arg, "-") && len(arg) == 2) || (strings.HasPrefix(arg, "--") && len(arg) >= 4) {
			continue
		}
		if strings.HasPrefix(arg, "--") && len(arg) == 3 {
			// Handle --o, convert to -o
			newArg := fmt.Sprintf("-%c", arg[2])
			msg := fmt.Sprintf("%s", ctx, "the flag %s is deprecated and will be removed in a future release. Please use the flag %s.", arg, newArg)
			fmt.Fprintf(os.Stderr, "WARNING: %s\n", msg)
			os.Args[i] = newArg
		} else if strings.HasPrefix(arg, "-") && len(arg) > 1 {
			// Handle -output, convert to --output
			newArg := fmt.Sprintf("-%s", arg)
			newArgType := "flag"
			if newArg == "--version" {
				newArg = "version"
				newArgType = "subcommand"
			}
			msg := fmt.Sprintf("%s", ctx, "the %s flag is deprecated and will be removed in a future release. "+
				"Please use the %s %s instead.",
				arg, newArg, newArgType,
			)
			fmt.Fprintf(os.Stderr, "WARNING: %s\n", msg)
			os.Args[i] = newArg
		}
	}
	if err := cli.New().Execute(); err != nil {
		// if the error is a `CosignError` then we want to use the exit code that
		// is related to the type of error that has occurred.
		var urlShortnerError *urlShortnerError.UrlShortnerError
		if errors.As(err, &urlShortnerError) {
			log.Printf("error during command execution: %v", err)
			os.Exit(urlShortnerError.ExitCode())
		}

		// we don't call os.Exit as Fatalf does both PrintF and os.Exit(1)
		log.Fatalf("error during command execution: %v", err)
	}
	// // CLI part
	// shortCmd := flag.NewFlagSet("short", flag.ExitOnError)
	// longURL := shortCmd.String("url", "", "URL to shorten (CLI mode)")

	// // Web server part
	// serverCmd := flag.NewFlagSet("server", flag.ExitOnError)
	// port := serverCmd.String("port", "8080", "Port to run the web server on")

	// if len(os.Args) < 2 {
	// 	fmt.Println("expected 'short' or 'server' subcommands")
	// 	os.Exit(1)
	// }

	// switch os.Args[1] {
	// case "short":
	// 	shortCmd.Parse(os.Args[2:])
	// 	if *longURL == "" {
	// 		fmt.Println("Please provide a URL to shorten using --url=\"http://example.com\"")
	// 		os.Exit(1)
	// 	}
	// 	shortURL := short.GenerateShortURL(*longURL)
	// 	fmt.Printf("Your short url: %s\n", shortURL)
	// case "server":
	// 	serverCmd.Parse(os.Args[2:])
	// 	fmt.Printf("Starting server on port %s...\n", *port)
	// 	// setupRoutes() // Function to set up web routes
	// 	setupRoutes()
	// 	log.Fatal(http.ListenAndServe(":"+*port, nil))
	// default:
	// 	fmt.Println("expected 'short' or 'server' subcommands")
	// 	os.Exit(1)
	// }
}

// func setupRoutes() {
// 	http.HandleFunc("/health", short.HealthCheckUp)
// 	http.HandleFunc("/short-url", short.HandleShortURL)
// 	http.HandleFunc("/long-url", short.HandleLongURL)
// }
