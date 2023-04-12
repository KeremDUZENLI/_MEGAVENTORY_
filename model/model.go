package model

import (
	"fmt"
	"strings"
)

type ProductCost struct {
	CompanyName         string `json:"CompanyName"`
	CompanyCurrencyCode string `json:"CompanyCurrencyCode"`
	ProductUnitCost     int    `json:"ProductUnitCost"`
}

type Product struct {
	ProductID       int           `json:"ProductID"`
	ProductType     string        `json:"ProductType"`
	ProductUnitCost []ProductCost `json:"ProductUnitCost"`
}

type SupplierClient struct {
	SupplierClientID   int    `json:"SupplierClientID"`
	SupplierClientType string `json:"SupplierClientType"`
	SupplierClientName string `json:"SupplierClientName"`
}

type ProductList struct {
	MvProducts []Product `json:"mvProducts"`
}

type SupplierClientList struct {
	MvSupplierClient SupplierClient `json:"mvSupplierClient"`
}

func (l ProductList) ConvertStringGet() string {
	var str strings.Builder

	for _, v := range l.MvProducts {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}

	return str.String()
}

func (s SupplierClientList) ConvertStringPost() string {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("%v\n", s.MvSupplierClient))

	return str.String()
}
