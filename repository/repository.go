package repository

import (
	"bytes"
	"fmt"
	"megaventory/common/env"
	"net/http"
)

type database struct{}

type Database interface {
	GetProducts() *http.Response
	PostProducts() *http.Request
}

func NewRepository() Database {
	return &database{}
}

func (database) GetProducts() *http.Response {
	response, err := http.Get(env.URL + env.GET + env.KEY)
	if err != nil {
		fmt.Println("Get error: ", err)
	}

	return response
}

func (database) PostProducts() *http.Request {
	response, err := http.NewRequest("POST", env.URL+env.POS, bytes.NewBuffer(env.RequestBody))
	if err != nil {
		fmt.Println("Post error: ", err)
	}

	return response
}
