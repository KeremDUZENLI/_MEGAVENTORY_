package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"megaventory/common/env"
	"net/http"
)

func getData(url string, parseStruct any) error {
	verifyConnection(getFromUrl(url))

	response := getFromUrl(url)
	if err := json.NewDecoder(response.Body).Decode(&parseStruct); err != nil {
		return errors.New("decode error")
	}

	return nil
}

func postData(url string, requestBody []byte, parseStruct any) error {
	verifyConnection(postToUrl(url, requestBody))

	request := postToUrl(url, requestBody)
	if err := json.NewDecoder(request.Body).Decode(&parseStruct); err != nil {
		return errors.New("decode error")
	}

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

func verifyConnection(r interface{}) {
	switch response := r.(type) {
	case *http.Response:
		fmt.Printf("\nResponse Status: \t%v", response.Status)
		fmt.Printf("\nResponse Header: \t%v\n", response.Header.Get("Content-Type"))
	case *http.Request:
		request, _ := client.Do(response)
		fmt.Printf("\nRequest Status: \t%v", request.Status)
		fmt.Printf("\nRequest Header: \t%v\n", request.Header.Get("Content-Type"))
	default:
		fmt.Println("Invalid type")
	}
}

func requestBodyMapper() []byte {
	eRequestBody, err := json.Marshal(env.SavedParseStruct)
	if err != nil {
		fmt.Println("Marshal error: ", err)
	}

	return eRequestBody
}
