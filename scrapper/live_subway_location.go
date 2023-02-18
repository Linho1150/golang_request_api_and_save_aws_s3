package scrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func JsonPrettyPrint(in string) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return []byte(in)
	}
	return out.Bytes()
}

func RequestApi(api_key string) []byte {
	fmt.Println("Start process ...")
	fmt.Println("Request API")
	url := "http://swopenapi.seoul.go.kr/api/subway/" + api_key + "/json/realtimeStationArrival/ALL"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("End process ...")
	return JsonPrettyPrint(string(body))
}
