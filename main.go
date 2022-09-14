package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("URL Shorten Service")

	// handleShortUrl function mapped to /short-url
	http.HandleFunc("/short-url", handleShortUrl)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
