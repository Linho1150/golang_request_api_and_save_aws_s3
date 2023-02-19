package main

import (
	"os"
	"strings"

	"githb.com/linho1150/repository"
	"githb.com/linho1150/scrapper"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(evnet MyEvent) (string, error) {
	API_KEYS := strings.Split(os.Getenv("API_KEY"), "/")
	for _, api_key := range API_KEYS {
		content := scrapper.RequestApi(api_key)
		if strings.Contains(string(content), "INFO-000") {
			repository.SaveJsonToS3(content)
			return "Successful", nil
		}
	}
	panic("No data")
}

func main() {
	lambda.Start(HandleRequest)
}
