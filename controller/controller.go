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
	GetInventoryController(w http.ResponseWriter, r *http.Request)
}

func NewController(h service.Holder) Sender {
	return &sender{sendList: h}
}

func (s sender) GetProductsController(w http.ResponseWriter, r *http.Request) {
	results := s.sendList.GetProductsService()
	fmt.Fprint(w, results.ConvertStringGet())
}

func (s sender) PostProductsController(w http.ResponseWriter, r *http.Request) {
	results := s.sendList.PostProductsService()
	fmt.Fprint(w, results.ConvertStringPost())
}

func (s sender) GetInventoryController(w http.ResponseWriter, r *http.Request) {
	results := s.sendList.GetInventoryService()
	fmt.Fprint(w, results.ConvertStringInventory())
}
