package repository

import (
	"megaventory/dto"
	"megaventory/model"
	"net/http"
)

var (
	productList           model.ProductList
	dtoSupplierClientHttp dto.DtoSupplierClientHttp
	inventoryList         model.InventoryList
	client                http.Client
)

type database struct{}

type Database interface {
	GetProducts() (model.ProductList, error)
	PostProducts() (model.SupplierClient, error)
	GetInventory() (model.InventoryList, error)
}

func NewRepository() Database {
	return &database{}
}
