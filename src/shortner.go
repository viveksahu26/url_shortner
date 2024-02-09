// // http://localhost:8080/long-url?sortURL=xtNFxaBwCG
package src

// This was earlier main.go program.

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"github.com/viveksahu26/url_shortner/src"
// )

// // http://localhost:8080/long-url?sortURL=xtNFxaBwCG
// func handleLongURL(writer http.ResponseWriter, req *http.Request) {
// 	// Procedure:
// 	// Finally got Complete url
// 	// Retrieve shortURL from it by Querying it.
// 	// Once retrieve, check  check in file whether it contains short url or not.
// 	// If file exist, then read that file:
// 	// 1. loop over it's content.
// 	// 2. Check line by line,
// 	// 3. short URL is present in that line or not.
// 	// 4. If present retun map and true boolean
// 	// 5. Else  retuen false and empty value

// 	if req.Method != "GET" {
// 		writer.WriteHeader(http.StatusMethodNotAllowed)
// 	} else {
// 		// get original--> shortURL
// 		originalURL := req.URL.Query().Get("sortURL")
// 		fmt.Println("originalURL: ", originalURL)

// 		longURL := src.CheckWhetherShortURLisPresentORNot(originalURL)
// 		fmt.Println("longURL: ", longURL)

// 		if longURL == "" {
// 			writer.Write([]byte("<h1>HSorry, No corresponding longURL is present to the shortURL. !!</h1>"))
// 		}
// 		writer.Write([]byte(longURL))
// 	}
// }

// func handleShortURL(writer http.ResponseWriter, req *http.Request) {
// 	if req.Method != "GET" {
// 		writer.WriteHeader(http.StatusMethodNotAllowed)
// 	} else {
// 		// get original URL from GET method by quering
// 		originalURL := req.URL.Query().Get("longURL")
// 		fmt.Println("originalURL: ", originalURL)

// 		// generate random shortURL
// 		shortURL := src.GenerateShortURL(originalURL)
// 		fmt.Println("shortURL: ", shortURL)

// 		// save short and long URL to file
// 		src.SaveInFile(shortURL, originalURL)

// 		host := req.Host

// 		// build Response
// 		resp := src.BuildURLWithResponse(host, shortURL, originalURL)

// 		err := src.RespondWithJSON(writer, 200, resp)
// 		if err != nil {
// 			writer.Write([]byte("<h1>Failed to respond with JSON</h1>"))
// 		}
// 	}
// }

// func healthCheckUp(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/health" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method other than GET not Supported...", http.StatusNotFound)
// 		return
// 	}

// 	w.Write([]byte("<h1>Health of Server is UP & Running... !!</h1>"))
// }

// // const addr = "localhost:8080"

// func main() {
// 	fmt.Println("URL Shorten Service Starts ...")

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}
// 	fmt.Println("PORT is: ", port)

// 	// /health endpoint is mapped to healthCheckUp
// 	http.HandleFunc("/health", healthCheckUp)

// 	// /short-url endpoint is mapped to handleShortURL
// 	http.HandleFunc("/short-url", handleShortURL)

// 	// /long-url endpoint is mapped to handleLongURL
// 	http.HandleFunc("/long-url", handleLongURL)

// 	// Server Listening on localhost:8080
// 	err := http.ListenAndServe(":"+port, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// }
