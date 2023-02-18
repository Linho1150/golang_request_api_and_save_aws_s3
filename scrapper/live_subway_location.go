package scrapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func JsonPrettyPrint(in string) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return []byte(in)
	}
	return out.Bytes()
}

func RequestApi(api_key string) (string, []byte) {
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
	url_parse := strings.Split(resp.Request.URL.String(), "/")
	fmt.Println("End process ...")
	return url_parse[len(url_parse)-1], JsonPrettyPrint(string(body))
}
