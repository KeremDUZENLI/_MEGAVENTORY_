package model

type Product struct {
	ProductID            int     `json:"ProductID"`
	ProductDescription   string  `json:"ProductDescription"`
	ProductSellingPrice  float64 `json:"ProductSellingPrice"`
	ProductPurchasePrice float64 `json:"ProductPurchasePrice"`
}

type ProductList struct {
	MvProducts []Product `json:"mvProducts"`
}
