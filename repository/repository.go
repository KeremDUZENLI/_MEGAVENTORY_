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
)

type database struct{}

type Database interface {
	GetProducts() *http.Response
	PostProducts() *http.Request
	GetInventory() (model.InventoryList, error)
	GetProductsById() (model.ProductList, error)
}

func NewRepository() Database {
	return &database{}
}

func (database) GetProducts() *http.Response {
	response, err := http.Get(env.URL + env.GET + env.KEY)
	if err != nil {
		fmt.Println("Get error: ", err)
	}

	verifyConncetion(response)
	return response
}

func (database) PostProducts() *http.Request {
	response, err := http.NewRequest("POST", env.URL+env.POS, bytes.NewBuffer(env.RequestBody))
	if err != nil {
		fmt.Println("Post error: ", err)
	}

	return response
}

func (database) GetInventory() (model.InventoryList, error) {
	if err := getData(env.URL+env.GETINVENTORY+env.KEY+env.PARAM, &inventoryList); err != nil {
		return model.InventoryList{}, err
	}

	return inventoryList, nil
}

// GIN ----------------------------------------------------------------
func (database) GetProductsById() (model.ProductList, error) {
	if err := getData(env.URL+env.GET+env.KEY, &productList); err != nil {
		return model.ProductList{}, err
	}

	return productList, nil
}

// HELP ----------------------------------------------------------------
func getData(url string, parseStruct any) error {
	response := connectUrl(url)

	if err := json.NewDecoder(response.Body).Decode(&parseStruct); err != nil {
		return errors.New("decode error")
	}

	return nil
}

func connectUrl(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Get error: ", err)
	}

	return response
}

func verifyConncetion(response *http.Response) {
	fmt.Println(response.Status)
	fmt.Println(response.Header.Get("Content-Type"))
}
