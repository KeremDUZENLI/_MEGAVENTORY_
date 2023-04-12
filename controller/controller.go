package controller

import (
	"fmt"
	"megaventory/service"
	"net/http"
)

type sender struct {
	sendList service.Holder
}

type Sender interface {
	GetProductsController(w http.ResponseWriter, r *http.Request)
	PostProductsController(w http.ResponseWriter, r *http.Request)
}

func NewController(h service.Holder) Sender {
	return &sender{sendList: h}
}

func (s sender) GetProductsController(w http.ResponseWriter, r *http.Request) {
	results := s.sendList.GetProductsService()
	fmt.Fprintf(w, results.ConvertStringGet())
}

func (s sender) PostProductsController(w http.ResponseWriter, r *http.Request) {
	results := s.sendList.PostProductsService()
	fmt.Fprintf(w, results.ConvertStringPost())
}
