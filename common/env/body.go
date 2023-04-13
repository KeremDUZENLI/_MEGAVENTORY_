package env

import (
	"encoding/json"
	"megaventory/model"
)

type Address model.SupplierClientAddress

var (
	SupplierClientName      string    = "xxx"
	SupplierClientEmail     string    = "xxx@gmail.com"
	SupplierClientPhone1    string    = "1234567890"
	mvRecordAction          string    = "Insert"
	SupplierClientAddresses []Address = []Address{{
		AddressLine1: "xxx street",
		City:         "athens",
	}}
)

var eRequestBody = map[string]interface{}{
	"APIKEY": KEY,
	"mvSupplierClient": map[string]interface{}{
		"SupplierClientName":      SupplierClientName,
		"SupplierClientEmail":     SupplierClientEmail,
		"SupplierClientAddresses": SupplierClientAddresses,
		"SupplierClientPhone1":    SupplierClientPhone1,
	},
	"mvRecordAction": mvRecordAction,
}

func requestBodyMapper() []byte {
	eRequestBody, err := json.Marshal(eRequestBody)
	if err != nil {
		panic(err)
	}

	return eRequestBody
}
