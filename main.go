package main

import (
	"os"
	"strings"

	"githb.com/linho1150/repository"
	"githb.com/linho1150/scrapper"
)

func main() {
	API_KEYS := strings.Split(os.Getenv("API_KEY"), "/")
	for _, api_key := range API_KEYS {
		filename, content := scrapper.RequestApi(api_key)
		if filename != "ALL" {
			repository.SaveJsonToS3(filename, content)
			break
		}
	}
}
