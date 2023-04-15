package repository

import (
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
