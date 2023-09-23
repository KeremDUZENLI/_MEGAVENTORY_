package service

import (
	"megaventory/dto"
	"megaventory/model"
	"megaventory/repository"
)

type holder struct {
	holdList repository.Database
}

type Holder interface {
	GetProductsService() (model.ProductList, error)
	PostProductsService() (model.SupplierClient, error)
	GetInventoryService() (model.InventoryList, error)

	GetProductsByIdService(id int) (model.ProductList, error)
	PostOneProductService(dP dto.DtoProduct) model.Product
	PostSupplierClientService(dSC dto.DtoSupplierClient) model.SupplierClient
	PostInventoryLocationService(dIL dto.DtoInvenoryLocation) model.Inventory
	PostTaxService(dT dto.DtoTax) model.TaxAndDiscount
	PostDiscountService(dD dto.DtoDiscount) model.TaxAndDiscount
}

func NewService(d repository.Database) Holder {
	return holder{holdList: d}
}
