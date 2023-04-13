package model

type SupplierClientAddresses struct {
	AddressLine1 string `json:"AddressLine1"`
	City         string `json:"City"`
}

type SupplierClients struct {
	SupplierClientName      string                    `json:"supplier_client_name"`
	SupplierClientEmail     string                    `json:"SupplierClientEmail"`
	SupplierClientAddresses []SupplierClientAddresses `json:"SupplierClientAddresses"`
	SupplierClientPhone1    string                    `json:"SupplierClientPhone1"`
}

type SupplierClientsTotal struct {
	MvSupplierClients []SupplierClients `json:"mvSupplierClients"`
}
