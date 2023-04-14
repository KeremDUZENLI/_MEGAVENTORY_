package service

import (
	"encoding/json"
	"fmt"
	"megaventory/model"
	"megaventory/repository"
)

var (
	productList        model.ProductList
	supplierClientList model.SupplierClientList
	inventoryList      model.InventoryList
)

type holder struct {
	holdList repository.Database
}

type Holder interface {
	GetProductsService() model.ProductList
	PostProductsService() model.SupplierClientList
	GetInventoryService() (model.InventoryList, error)
	GetProductsByIdService() (model.ProductList, error)
}

func NewService(d repository.Database) Holder {
	return holder{holdList: d}
}

func (h holder) GetProductsService() model.ProductList {
	results := h.holdList.GetProducts()

	err := json.NewDecoder(results.Body).Decode(&productList)
	if err != nil {
		fmt.Println("responseGet is not parsed: ", err)
	}

	return productList
}

func (h holder) PostProductsService() model.SupplierClientList {
	results := h.holdList.PostProducts()

	err := json.NewDecoder(results.Body).Decode(&supplierClientList)
	if err != nil {
		fmt.Println("responsePost is not parsed: ", err)
	}

	return supplierClientList
}

func (h holder) GetInventoryService() (model.InventoryList, error) {
	return h.holdList.GetInventory()
}

// GIN ----------------------------------------------------------------
func (h holder) GetProductsByIdService() (model.ProductList, error) {
	return h.holdList.GetProductsById()
}
