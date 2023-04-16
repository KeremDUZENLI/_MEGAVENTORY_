package mapper

import (
	"megaventory/dto"
	"megaventory/model"
)

func MapperProduct(d *dto.DtoProduct) model.Product {
	return model.Product{
		ProductID:            d.SKU,
		ProductDescription:   d.Description,
		ProductSellingPrice:  d.SalesPrice,
		ProductPurchasePrice: d.PurchasePrice,
	}
}

func MapperSupplierClientHttp(d *dto.DtoSupplierClientHttp) model.SupplierClient {
	return model.SupplierClient{
		SupplierClientName:  d.MvSupplierClient.SupplierClientName,
		SupplierClientEmail: d.MvSupplierClient.SupplierClientEmail,
		SupplierClientAddresses: []model.SupplierClientAddress{
			{
				AddressLine1: d.MvSupplierClient.SupplierClientAddresses.SupplierClientAddressLine1,
				City:         d.MvSupplierClient.SupplierClientAddresses.SupplierClientCity,
			},
		},
		SupplierClientPhone1: d.MvSupplierClient.SupplierClientPhone1,
	}
}

func MapperSupplierClient(d *dto.DtoSupplierClient) model.SupplierClient {
	return model.SupplierClient{
		SupplierClientName:  d.Name,
		SupplierClientEmail: d.Email,
		SupplierClientAddresses: []model.SupplierClientAddress{
			{
				AddressLine1: d.ShippingAddress,
				City:         d.ShippingCity,
			},
		},
		SupplierClientPhone1: d.Phone,
	}
}

func MapperInventoryLocation(d *dto.DtoInvenoryLocation) model.Inventory {
	return model.Inventory{
		ProductID: d.Abbreviation,
		MvStock: []model.Stock{
			{
				InventoryLocationID: d.Name,
				SubLocation:         d.Address,
			},
		},
	}
}

func MapperTax(d *dto.DtoTax) model.TaxAndDiscount {
	return model.TaxAndDiscount{
		Name:        d.Name,
		Description: d.Description,
		Value:       d.Value,
	}
}

func MapperDiscount(d *dto.DtoDiscount) model.TaxAndDiscount {
	return model.TaxAndDiscount{
		NameDiscount:        d.Name,
		DescriptionDiscount: d.Description,
		ValueDiscount:       d.Value,
	}
}
