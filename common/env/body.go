package env

import (
	"encoding/json"
)

var (
	supplierClientID   int    = 0
	supplierClientType string = "Client30"
	supplierClientName string = "My dummy client30"
	mvRecordAction     string = "Insert"
)

var eRequestBody = map[string]interface{}{
	"APIKEY": KEY,
	"mvSupplierClient": map[string]interface{}{
		"SupplierClientID":   supplierClientID,
		"SupplierClientType": supplierClientType,
		"SupplierClientName": supplierClientName,
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
