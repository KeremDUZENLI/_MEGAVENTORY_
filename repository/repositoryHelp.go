package repository

import (
	"bytes"
	"fmt"
	"net/http"
)

func getData(url string, parseStruct any) error {
	verifyConnection(getFromUrl(url))

	response := getFromUrl(url)
	decode(response, parseStruct)
	return nil
}

func postData(url string, requestBody []byte, parseStruct any) error {
	verifyConnection(postToUrl(url, requestBody))

	request := postToUrl(url, requestBody)
	decode(request, parseStruct)
	return nil
}

func getFromUrl(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Get error: ", err)
	}

	return response
}

func postToUrl(url string, requestBody []byte) *http.Request {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Post error: ", err)
	}

	return request
}
