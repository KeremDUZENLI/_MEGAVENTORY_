package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	ProductWeight   int           `json:"ProductWeight"`
}

type MegaventoryList struct {
	MvProducts []Product `json:"mvProducts"`
}

func main() {
	url := "https://api.megaventory.com/v2017a/json/reply/ProductGet?APIKEY=8ccd0b3378ef30a5@m140829"

	// repository
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting values: ", err)
		return
	}
	defer response.Body.Close()

	// service
	var megaventoryList MegaventoryList
	err = json.NewDecoder(response.Body).Decode(&megaventoryList)
	if err != nil {
		fmt.Println("Result is not parsed: ", err)
		return
	}

	// controller
	for _, product := range megaventoryList.MvProducts {
		for _, cost := range product.ProductUnitCost {
			fmt.Printf("ProductID: %v   ProductType: %v   ProductUnitCost: %v   ProductWeight: %v\n",
				product.ProductID, product.ProductType, cost.CompanyName, product.ProductWeight)
		}
	}
}
