package main

import (
	"githb.com/linho1150/repository"
	"githb.com/linho1150/scrapper"
)

func main() {
	filename, content := scrapper.RequestApi()
	repository.SaveJsonToS3(filename, content)
}
