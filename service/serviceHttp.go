package service

import "megaventory/model"

func (h holder) GetProductsService() (model.ProductList, error) {
	return h.holdList.GetProducts()
}

func (h holder) PostProductsService() (model.SupplierClientList, error) {
	return h.holdList.PostProducts()
}

func (h holder) GetInventoryService() (model.InventoryList, error) {
	return h.holdList.GetInventory()
}
