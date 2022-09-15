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

func IsLongURLPresentInFile(fileName, longURL string) (map[string]string, string, bool) {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	fileContent := strings.Split(string(fileBytes), "\n")

	fileContainLongURL := false
	var message string
	ShortAndLongURLValue := make(map[string]string)
	for _, line := range fileContent {
		totalURL := strings.Split(line, "=")
		if len(totalURL) == 2 {
			ShortAndLongURLValue[totalURL[1]] = totalURL[0]
		}

		if strings.HasSuffix(line, longURL) {
			message = fmt.Sprintln("Yes, URL is already present in the file")
			fileContainLongURL = true
			return ShortAndLongURLValue, message, fileContainLongURL
		}
		message = fmt.Sprintln("URL is not present in the file.")
	}
	fmt.Println("\n", message)
	return ShortAndLongURLValue, message, fileContainLongURL
}
