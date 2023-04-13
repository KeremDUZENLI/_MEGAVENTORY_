package model

type SupplierClient struct {
	SupplierClientID   int    `json:"SupplierClientID"`
	SupplierClientType string `json:"SupplierClientType"`
	SupplierClientName string `json:"SupplierClientName"`
}

type SupplierClientList struct {
	MvSupplierClient SupplierClient `json:"mvSupplierClient"`
}
