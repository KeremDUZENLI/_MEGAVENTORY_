package dto

type DtoSupplierClientHttp struct {
	APIKey string `json:"APIKEY" binding:"required"`

	MvSupplierClient struct {
		SupplierClientID     int    `json:"SupplierClientID" binding:"required"`
		SupplierClientType   string `json:"SupplierClientType" binding:"required"`
		SupplierClientName   string `json:"SupplierClientName" binding:"required"`
		SupplierClientEmail  string `json:"SupplierClientEmail" binding:"required"`
		SupplierClientPhone1 string `json:"SupplierClientPhone1" binding:"required"`

		SupplierClientAddresses struct {
			SupplierClientAddressLine1 string `json:"SupplierClientAddressLine1" binding:"required"`
			SupplierClientCity         string `json:"SupplierClientAddress" binding:"required"`
		} `json:"SupplierClientAddresses" binding:"required"`
	} `json:"mvSupplierClient" binding:"required"`

	MvRecordAction string `json:"mvRecordAction" binding:"required"`
}
