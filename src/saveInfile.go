package src

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Saving LongURL corresponding to ShortURL
func SaveInFile(shortURL string, longURL string) {
	fileName := "url.properties"
	prop := shortURL + "=" + longURL
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		os.Create(fileName)
		os.WriteFile(fileName, []byte(prop+"\n"), 0o644)
	} else {

		input, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalln(err)
		}

		lines := strings.Split(string(input), "\n")

		var flag bool
		for i, line := range lines {
			if strings.Contains(line, shortURL) {
				lines[i] = prop
				flag = true
			}
		}
		if flag {
			output := strings.Join(lines, "\n")
			err = ioutil.WriteFile(fileName, []byte(output), 0o644)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
			if err != nil {
				panic(err)
			}
			_, _ = file.WriteString(prop + "\n")
		}
	}
}
