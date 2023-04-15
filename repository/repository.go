package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"megaventory/common/env"
	"megaventory/model"
	"net/http"
)

var (
	productList        model.ProductList
	supplierClientList model.SupplierClientList
	inventoryList      model.InventoryList
	client             http.Client
)

type database struct{}

type Database interface {
	GetProducts() (model.ProductList, error)
	PostProducts() (model.SupplierClientList, error)
	GetInventory() (model.InventoryList, error)
}

func NewRepository() Database {
	return &database{}
}

func (database) GetProducts() (model.ProductList, error) {
	if err := getData(env.URL+env.GET+env.KEY, &productList); err != nil {
		return model.ProductList{}, err
	}

	return productList, nil
}

func (database) PostProducts() (model.SupplierClientList, error) {
	if err := postData(env.URL+env.POS, env.RequestBody, &supplierClientList); err != nil {
		return model.SupplierClientList{}, err
	}

	return supplierClientList, nil
}

func (database) GetInventory() (model.InventoryList, error) {
	if err := getData(env.URL+env.GETINVENTORY+env.KEY+env.PARAM, &inventoryList); err != nil {
		return model.InventoryList{}, err
	}

	return inventoryList, nil
}

// HELP ----------------------------------------------------------------
func getData(url string, parseStruct any) error {
	response := getFromUrl(url)

	if err := json.NewDecoder(response.Body).Decode(&parseStruct); err != nil {
		return errors.New("decode error")
	}

	verifyConnection(response)
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
