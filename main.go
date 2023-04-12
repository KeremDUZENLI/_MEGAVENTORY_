package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// variables
var (
	URL = "https://api.megaventory.com/v2017a/"

	GET = "json/reply/ProductGet?APIKEY="
	POS = "SupplierClient/SupplierClientUpdate"

	KEY = "8ccd0b3378ef30a5@m140829"

	HOST = ":8080"
)

var requestBody = bytes.NewBufferString(`{
	"APIKEY": "8ccd0b3378ef30a5@m140829",
	"mvSupplierClient": {
	  "SupplierClientID": 0,
	  "SupplierClientType": "Client33",
	  "SupplierClientName": "My dummy client33"
	},
	"mvRecordAction": "Insert"
}`)

var productList ProductList
var supplierClientList SupplierClientList

// model
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

// repository
func GetProducts() *http.Response {
	response, err := http.Get(URL + GET + KEY)
	if err != nil {
		fmt.Println("Get error: ", err)
	}

	return response
}

func PostProducts() *http.Request {
	response, err := http.NewRequest("POST", URL+POS, requestBody)
	if err != nil {
		fmt.Println("Post error: ", err)
	}

	return response
}

// service
func GetProductsService() ProductList {
	results := GetProducts()

	err := json.NewDecoder(results.Body).Decode(&productList)
	if err != nil {
		fmt.Println("responseGet is not parsed: ", err)
	}

	return productList
}

func PostProductsService() SupplierClientList {
	results := PostProducts()

	err := json.NewDecoder(results.Body).Decode(&supplierClientList)
	if err != nil {
		fmt.Println("responsePost is not parsed: ", err)
	}

	return supplierClientList
}

// controller
func GetProductsController(w http.ResponseWriter, r *http.Request) {
	results := GetProductsService()
	fmt.Fprintf(w, convertStringGet(results))
}

func PostProductsController(w http.ResponseWriter, r *http.Request) {
	results := PostProductsService()
	fmt.Fprintf(w, convertStringPost(results))
}

func convertStringGet(listeGet ProductList) string {
	var str strings.Builder

	for _, v := range listeGet.MvProducts {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}

	return str.String()
}

func convertStringPost(listePost SupplierClientList) string {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("%v\n", listePost.MvSupplierClient))

	return str.String()
}

// router
func Run(host string) {
	http.HandleFunc("/get", GetProductsController)
	http.HandleFunc("/post", PostProductsController)
	http.ListenAndServe(host, nil)
}

func main() {
	Run(HOST)
}
