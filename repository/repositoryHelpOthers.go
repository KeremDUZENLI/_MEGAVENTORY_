package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"megaventory/common/env"
	"net/http"
)

func verifyConnection(r any) {
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

func decode(r any, parseStruct any) error {
	switch v := r.(type) {
	case *http.Response:
		if err := json.NewDecoder(v.Body).Decode(parseStruct); err != nil {
			return errors.New("decode error")
		}
	case *http.Request:
		if err := json.NewDecoder(v.Body).Decode(parseStruct); err != nil {
			return errors.New("decode error")
		}
	default:
		return errors.New("unexpected type")
	}

	return nil
}

func marshal() []byte {
	eRequestBody, err := json.Marshal(env.SavedParseStruct)
	if err != nil {
		fmt.Println("Marshal error: ", err)
	}

	return eRequestBody
}
