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
	eURL = "https://api.megaventory.com/v2017a/"

	eGET = "json/reply/ProductGet?APIKEY="
	ePOS = "SupplierClient/SupplierClientUpdate"

	eKEY = "8ccd0b3378ef30a5@m140829"

	eHOST = ":8080"

	eRequestBody = bytes.NewBufferString(`{
		"APIKEY": "8ccd0b3378ef30a5@m140829",
		"mvSupplierClient": {
		  "SupplierClientID": 0,
		  "SupplierClientType": "Client33",
		  "SupplierClientName": "My dummy client33"
		},
		"mvRecordAction": "Insert"
	}`)
)

var (
	URL string

	GET string
	POS string

	KEY string

	HOST string

	RequestBody *bytes.Buffer
)

var (
	productList        ProductList
	supplierClientList SupplierClientList
)

func Load() {
	URL = eURL

	GET = eGET
	POS = ePOS

	KEY = eKEY

	HOST = eHOST

	RequestBody = eRequestBody
}

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

func (l ProductList) convertStringGet() string {
	var str strings.Builder

	for _, v := range l.MvProducts {
		str.WriteString(fmt.Sprintf("%v\n", v))
	}

	return str.String()
}

func (s SupplierClientList) convertStringPost() string {
	var str strings.Builder

	str.WriteString(fmt.Sprintf("%v\n", s.MvSupplierClient))

	return str.String()
}

// repository
type database struct{}

type Database interface {
	GetProducts() *http.Response
	PostProducts() *http.Request
}

func NewRepository() Database {
	return &database{}
}

func (database) GetProducts() *http.Response {
	response, err := http.Get(URL + GET + KEY)
	if err != nil {
		fmt.Println("Get error: ", err)
	}

	return response
}

func (database) PostProducts() *http.Request {
	response, err := http.NewRequest("POST", URL+POS, RequestBody)
	if err != nil {
		fmt.Println("Post error: ", err)
	}

	return response
}

// service
type holder struct {
	holdList Database
}

type Holder interface {
	GetProductsService() ProductList
	PostProductsService() SupplierClientList
}

func NewService(d Database) Holder {
	return holder{holdList: d}
}

func (h holder) GetProductsService() ProductList {
	results := h.holdList.GetProducts()

	err := json.NewDecoder(results.Body).Decode(&productList)
	if err != nil {
		fmt.Println("responseGet is not parsed: ", err)
	}

	return productList
}

func (h holder) PostProductsService() SupplierClientList {
	results := h.holdList.PostProducts()

	err := json.NewDecoder(results.Body).Decode(&supplierClientList)
	if err != nil {
		fmt.Println("responsePost is not parsed: ", err)
	}

	return supplierClientList
}

// controller
type sender struct {
	sendList Holder
}

type Sender interface {
	GetProductsController(w http.ResponseWriter, r *http.Request)
	PostProductsController(w http.ResponseWriter, r *http.Request)
}

func NewController(h Holder) Sender {
	return &sender{sendList: h}
}

func (s sender) GetProductsController(w http.ResponseWriter, r *http.Request) {
	results := s.sendList.GetProductsService()
	fmt.Fprintf(w, results.convertStringGet())
}

func (s sender) PostProductsController(w http.ResponseWriter, r *http.Request) {
	results := s.sendList.PostProductsService()
	fmt.Fprintf(w, results.convertStringPost())
}

// router
type router struct {
	routeList Sender
}

type Router interface {
	Run(host string)
}

func NewRouter(s Sender) Router {
	router := &router{routeList: s}
	router.setup()

	return router
}

func (r *router) Run(host string) {
	http.ListenAndServe(host, nil)
}

func (r *router) setup() {
	http.HandleFunc("/get", r.routeList.GetProductsController)
	http.HandleFunc("/post", r.routeList.PostProductsController)
}

func main() {
	Load()
	set := settingValues()

	set.Run(HOST)
}

func settingValues() Router {
	repo := NewRepository()
	serv := NewService(repo)
	cont := NewController(serv)
	rout := NewRouter(cont)

	return rout
}
