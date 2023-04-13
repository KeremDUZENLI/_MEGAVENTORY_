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

func (s SupplierClientList) ConvertStringPost() string {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("%v\n", s.MvSupplierClient))

	return str.String()
}

func (s InventoryList) ConvertStringInventory() string {
	var str strings.Builder

	for _, v := range s.MvProductStockList {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}

	return str.String()
}
