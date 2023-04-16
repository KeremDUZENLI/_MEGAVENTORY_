package repository

import (
	"encoding/json"
	"megaventory/common/env"
	"megaventory/model"
)

func (database) GetProducts() (model.ProductList, error) {
	if err := getData(env.URL+env.GET+env.APIKEY, &productList); err != nil {
		return model.ProductList{}, err
	}

	return productList, nil
}

func (database) PostProducts() (model.SupplierClientList, error) {
	if err := postData(env.URL+env.POS, env.RequestBody, &supplierClientList); err != nil {
		return model.SupplierClientList{}, err
	}

	return supplierClientList, nil
}

func (database) GetInventory() (model.InventoryList, error) {
	if err := getData(env.URL+env.GETINVENTORY+env.APIKEY+env.PARAM, &inventoryList); err != nil {
		return model.InventoryList{}, err
	}

	return inventoryList, nil
}

func requestBodyMapper() []byte {
	eRequestBody, err := json.Marshal(eRequestBody)
	if err != nil {
		panic(err)
	}

	return eRequestBody
}
