package main

import (
	"context"
	"os"
	"strings"

	"githb.com/linho1150/repository"
	"githb.com/linho1150/scrapper"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) {
	API_KEYS := strings.Split(os.Getenv("API_KEY"), "/")
	for _, api_key := range API_KEYS {
		filename, content := scrapper.RequestApi(api_key)
		if filename != "ALL" {
			repository.SaveJsonToS3(filename, content)
			break
		}
	}
}

func main() {
	lambda.Start(HandleRequest)
}
