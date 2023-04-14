package service

import (
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
	GetProductsService() (model.ProductList, error)
	PostProductsService() (model.SupplierClientList, error)
	GetInventoryService() (model.InventoryList, error)
	GetProductsByIdService() (model.ProductList, error)
}

func NewService(d repository.Database) Holder {
	return holder{holdList: d}
}

func (h holder) GetProductsService() (model.ProductList, error) {
	return h.holdList.GetProducts()
}

func (h holder) PostProductsService() (model.SupplierClientList, error) {
	return h.holdList.PostProducts()
}

func (h holder) GetInventoryService() (model.InventoryList, error) {
	return h.holdList.GetInventory()
}

// GIN ----------------------------------------------------------------
func (h holder) GetProductsByIdService() (model.ProductList, error) {
	return h.holdList.GetProductsById()
}
