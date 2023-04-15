package dto

type DtoProduct struct {
	SKU           int     `json:"SKU" binding:"required"`
	Description   string  `json:"Description" binding:"required"`
	SalesPrice    float64 `json:"SalesPrice" binding:"required"`
	PurchasePrice float64 `json:"PurchasePrice" binding:"required"`
}

type DtoSupplierClient struct {
	Name            string `json:"Name" binding:"required"`
	Email           string `json:"Email" binding:"required"`
	ShippingAddress string `json:"ShippingAddress" binding:"required"`
	ShippingCity    string `json:"ShippingCity" binding:"required"`
	Phone           string `json:"Phone" binding:"required"`
}

type DtoInvenoryLocation struct {
	Abbreviation int    `json:"Abbreviation" binding:"required"`
	Name         int    `json:"Name" binding:"required"`
	Address      string `json:"Address" binding:"required"`
}

type DtoTax struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description" binding:"required"`
	Value       int    `json:"Value" binding:"required"`
}

type DtoDiscount struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description" binding:"required"`
	Value       int    `json:"Value" binding:"required"`
}
