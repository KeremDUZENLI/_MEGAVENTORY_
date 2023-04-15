package controller

import (
	"fmt"
	"net/http"
)

func (s sender) GetProductsController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.GetProductsService()
	httpErrorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringGet())
}

func (s sender) PostProductsController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.PostProductsService()
	httpErrorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringPost())
}

func (s sender) GetInventoryController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.GetInventoryService()
	httpErrorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringInventory())
}
