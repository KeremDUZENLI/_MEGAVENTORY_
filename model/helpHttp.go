package model

import (
	"fmt"
	"strings"
)

func (l ProductList) ConvertStringGet() string {
	var str strings.Builder

	for _, v := range l.MvProducts {
		str.WriteString(fmt.Sprintf("-SKU: %v\n-Description: %v\n-Sales Price: %v\n-Purchase Price: %v\n\n",
			v.ProductID,
			v.ProductDescription,
			v.ProductSellingPrice,
			v.ProductPurchasePrice))
	}

	return str.String()
}

func (s SupplierClient) ConvertStringPost() string {
	var str strings.Builder

	for _, l := range s.SupplierClientAddresses {
		str.WriteString(fmt.Sprintf("-Name: %v\n-E-mail: %v\n-Shipping Address: %v, %v\n-Phone: %v\n\n",
			s.SupplierClientName,
			s.SupplierClientEmail,
			l.AddressLine1, l.City,
			s.SupplierClientPhone1))
	}

	return str.String()
}

func (s InventoryList) ConvertStringInventory() string {
	var str strings.Builder

	for _, v := range s.MvProductStockList {
		for _, l := range v.MvStock {
			str.WriteString(fmt.Sprintf("-ProductID: %v\n-InventoryLocationID: %v\n-SubLocation: %v\n\n",
				v.ProductID,
				l.InventoryLocationID,
				l.SubLocation))
		}

	}

	return str.String()
}
