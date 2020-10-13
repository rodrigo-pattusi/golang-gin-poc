package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "localhost:8080/videos"
	method := "POST"

	payload := strings.NewReader("{\n	\"title\": \"title-1\",\n	\"description\": \"a description\",\n	\"url\": \"url-1\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
