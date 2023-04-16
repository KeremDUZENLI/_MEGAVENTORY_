package controller

import (
	"encoding/json"
	"fmt"
	"megaventory/common/env"
	"megaventory/dto"
	"net/http"
)

func (s sender) GetProductsController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.GetProductsService()
	httpErrorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringGet())
}

func (s sender) PostProductsController(w http.ResponseWriter, r *http.Request) {
	var parseStruct dto.DtoSupplierClientHttp

	if err := json.NewDecoder(r.Body).Decode(&parseStruct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	env.SavedParseStruct = parseStruct

	results, err := s.sendList.PostProductsService()
	httpErrorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringPost())
}

func (s sender) GetInventoryController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.GetInventoryService()
	httpErrorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringInventory())
}
