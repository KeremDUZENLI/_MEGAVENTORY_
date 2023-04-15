package service

import (
	"megaventory/dto"
	"megaventory/dto/mapper"
	"megaventory/model"
)

func (h holder) GetProductsByIdService(id int) (model.ProductList, error) {
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
