package repository

import (
	"megaventory/common/env"
	"megaventory/dto/mapper"
	"megaventory/model"
)

func (database) GetProducts() (model.ProductList, error) {
	if err := getData(env.URL+env.GET+env.APIKEY, &productList); err != nil {
		return model.ProductList{}, err
	}

	return productList, nil
}

func (database) PostProducts() (model.SupplierClient, error) {
	if err := postData(env.URL+env.POS, marshal(), &dtoSupplierClientHttp); err != nil {
		return model.SupplierClient{}, err
	}

	aMap := mapper.MapperSupplierClientHttp(&dtoSupplierClientHttp)
	return aMap, nil
}

func (database) GetInventory() (model.InventoryList, error) {
	if err := getData(env.URL+env.GETINVENTORY+env.APIKEY+env.PARAM, &inventoryList); err != nil {
		return model.InventoryList{}, err
	}

	return inventoryList, nil
}
