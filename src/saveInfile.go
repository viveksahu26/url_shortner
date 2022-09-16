package src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Saving LongURL corresponding to ShortURL
func SaveInFile(shortURL string, longURL string) {
	fileName := "url.properties"
	isFileExist, _ := IsFileExist(fileName)

	combinedURL := shortURL + "=" + longURL

	if !isFileExist {
		fmt.Printf("New file %s is created.", fileName)
		os.Create(fileName)
		os.WriteFile(fileName, []byte(combinedURL+"\n"), 0o644)

	} else {
		fmt.Printf("File %s already exist.", fileName)
		_, _, fileContainLongURL := IsLongURLPresentInFile(fileName, longURL)

		if !fileContainLongURL {
			file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
			if err != nil {
				panic(err)
			}
			_, _ = file.WriteString(combinedURL + "\n")
		}
	}
}

func IsFileExist(fileName string) (bool, error) {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return false, err
	}
	return true, nil
}

// It can check whether LongURL is present in the file or not.
// As well as also it can check whether shortURL is present in the file or not.
func IsLongURLPresentInFile(fileName, url string) (map[string]string, string, bool) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	fileContent := strings.Split(string(fileBytes), "\n")

	incommingLongURLisShortURL := false
	if len(url) == 10 {
		// way  of differentialting bewtween URL: whether the argument url
		// is actual longURL or shortURL
		// ShortURL is of length 10

		// yes, it's a shortURL
		incommingLongURLisShortURL = true
	}
	fmt.Println("incommingLongURLisShortURL: ", incommingLongURLisShortURL)

	fileContainLongURL := false
	var message string
	ShortAndLongURLValue := make(map[string]string)

	for _, line := range fileContent {
		totalURL := strings.Split(line, "=")
		if len(totalURL) == 2 {
			// if Long URL, then store key as LongURL and value as ShortURL in map
			ShortAndLongURLValue[totalURL[1]] = totalURL[0]

			// if short URL, then store key as shortURL and value as LongURL in map
			if incommingLongURLisShortURL {
				fmt.Println("storing map values for shortURL as key and longURL as value")
				ShortAndLongURLValue[totalURL[0]] = totalURL[1]
			}
		}

		if strings.HasSuffix(line, url) && !incommingLongURLisShortURL {
			message = fmt.Sprintln("Yes, URL is already present in the file")
			fileContainLongURL = true
			return ShortAndLongURLValue, message, fileContainLongURL
		} else if strings.HasPrefix(line, url) && incommingLongURLisShortURL {
			message = fmt.Sprintln("Yes, URL is already present in the file")
			fileContainLongURL = true
			return ShortAndLongURLValue, message, fileContainLongURL
		}

		message = fmt.Sprintln("URL is not present in the file.")
	}
	fmt.Println("\n", message)
	return ShortAndLongURLValue, message, fileContainLongURL
}
