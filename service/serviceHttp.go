package service

import (
	"megaventory/dto"
	"megaventory/dto/mapper"
	"megaventory/model"
)

func (h holder) GetProductsService() (model.ProductList, error) {
	return h.holdList.GetProducts()
}

func (h holder) PostProductsService() (model.SupplierClientList, error) {
	return h.holdList.PostProducts()
}

func (h holder) PostProductsService2(dSCH dto.DtoSupplierClientHttp) model.SupplierClient {
	aMap := mapper.MapperSupplierClientHttp(&dSCH)
	return aMap
}

func (h holder) GetInventoryService() (model.InventoryList, error) {
	return h.holdList.GetInventory()
}
