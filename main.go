package main

import (
	"bytes"
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

func main() {
	urlGet := "https://api.megaventory.com/v2017a/json/reply/ProductGet?APIKEY=8ccd0b3378ef30a5@m140829"
	urlPost := "https://api.megaventory.com/v2017a/SupplierClient/SupplierClientUpdate"

	requestBody := bytes.NewBufferString(`{
        "APIKEY": "8ccd0b3378ef30a5@m140829",
        "mvSupplierClient": {
          "SupplierClientID": 0,
          "SupplierClientType": "Client25",
          "SupplierClientName": "My dummy client25"
        },
        "mvRecordAction": "Insert"
      }`)

	// repository
	responseGet, errGet := http.Get(urlGet)
	if errGet != nil {
		fmt.Println("Get error: ", errGet)
		return
	}
	defer responseGet.Body.Close()

	responsePost, errPost := http.Post(urlPost, "application/json", requestBody)
	if errPost != nil {
		fmt.Println("Post error: ", errPost)
	}
	defer responsePost.Body.Close()

	// service
	var productList ProductList
	errG := json.NewDecoder(responseGet.Body).Decode(&productList)
	if errG != nil {
		fmt.Println("responseGet is not parsed: ", errG)
		return
	}

	var supplierClientList SupplierClientList
	errP := json.NewDecoder(responsePost.Body).Decode(&supplierClientList)
	if errP != nil {
		fmt.Println("responsePost is not parsed: ", errP)
	}

	// controller
	// for _, product := range productList.MvProducts {
	// 	fmt.Printf("ProductID: %v   ProductType: %v   ProductUnitCost: %v\n",
	// 		product.ProductID, product.ProductType, product.ProductUnitCost)
	// }

	fmt.Println(responseGet.Header.Get("Content-Type"))
	fmt.Println(responsePost.Header.Get("Content-Type"))
}
