package model

type Stock struct {
	InventoryLocationID int    `json:"InventoryLocationID"`
	SubLocation         string `json:"SubLocation"`
}

type Inventory struct {
	ProductID int     `json:"productID"`
	MvStock   []Stock `json:"mvStock"`
}

type InventoryList struct {
	MvProductStockList []Inventory `json:"mvProductStockList"`
}
