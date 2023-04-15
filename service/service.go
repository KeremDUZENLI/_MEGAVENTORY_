package service

import (
	"megaventory/dto"
	"megaventory/dto/mapper"
	"megaventory/model"
	"megaventory/repository"
)

type holder struct {
	holdList repository.Database
}

type Holder interface {
	GetProductsService() (model.ProductList, error)
	PostProductsService() (model.SupplierClientList, error)
	GetInventoryService() (model.InventoryList, error)

	GetProductsServiceById(id int) (model.ProductList, error)
	PostOneProductService(dP dto.DtoProduct) model.Product
	PostSupplierClientService(dSC dto.DtoSupplierClient) model.SupplierClient
	PostInventoryLocationService(dIL dto.DtoInvenoryLocation) model.Inventory
	PostTaxService(dT dto.DtoTax) model.TaxAndDiscount
	PostDiscountService(dD dto.DtoDiscount) model.TaxAndDiscount
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
func (h holder) GetProductsServiceById(id int) (model.ProductList, error) {
	return h.holdList.GetProducts()
}

func (h holder) PostOneProductService(dP dto.DtoProduct) model.Product {
	aMap := mapper.MapperProduct(&dP)
	return aMap
}

func (h holder) PostSupplierClientService(dSC dto.DtoSupplierClient) model.SupplierClient {
	aMap := mapper.MapperSupplierClient(&dSC)
	return aMap
}

func (h holder) PostInventoryLocationService(dIL dto.DtoInvenoryLocation) model.Inventory {
	aMap := mapper.MapperInventoryLocation(&dIL)
	return aMap
}

func (h holder) PostTaxService(dT dto.DtoTax) model.TaxAndDiscount {
	aMap := mapper.MapperTax(&dT)
	return aMap
}

func (h holder) PostDiscountService(dD dto.DtoDiscount) model.TaxAndDiscount {
	aMap := mapper.MapperDiscount(&dD)
	return aMap
}
