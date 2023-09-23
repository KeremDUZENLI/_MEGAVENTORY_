package model

type SupplierClientAddress struct {
	AddressLine1 string `json:"AddressLine1"`
	City         string `json:"City"`
}

type SupplierClient struct {
	SupplierClientName      string                  `json:"SupplierClientName"`
	SupplierClientEmail     string                  `json:"SupplierClientEmail"`
	SupplierClientAddresses []SupplierClientAddress `json:"SupplierClientAddresses"`
	SupplierClientPhone1    string                  `json:"SupplierClientPhone1"`
}
