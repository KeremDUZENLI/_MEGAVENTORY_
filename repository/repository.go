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
	productList model.ProductList
)

type database struct{}

type Database interface {
	GetProducts() *http.Response
	PostProducts() *http.Request
	GetInventory() *http.Response
	GetProductsById(id int) (model.ProductList, error)
}

func NewRepository() Database {
	return &database{}
}

func verifyConncetion(response *http.Response) {
	fmt.Println(response.Status)
	fmt.Println(response.Header.Get("Content-Type"))
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

func (database) GetInventory() *http.Response {
	param := "&ShowOnlyProductsWithPositiveQty=1"
	response, err := http.Get(env.URL + env.GETINVENTORY + env.KEY + param)
	if err != nil {
		fmt.Println("GetInventory error: ", err)
	}

	verifyConncetion(response)
	return response
}

// GIN ----------------------------------------------------------------
func (database) GetProductsById(id int) (model.ProductList, error) {
	response, err := http.Get(env.URL + env.GET + env.KEY)
	if err != nil {
		fmt.Println("Get error: ", err)
	}

	json.NewDecoder(response.Body).Decode(&productList)
	if err != nil {
		return model.ProductList{}, errors.New("responseGet is not parsed")
	}

	return productList, nil
}
