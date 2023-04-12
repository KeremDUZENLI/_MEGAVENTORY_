package env

import (
	"bytes"
	"os"

	"github.com/joho/godotenv"
)

var eRequestBody = bytes.NewBufferString(`{
    "APIKEY": "8ccd0b3378ef30a5@m140829",
    "mvSupplierClient": {
        "SupplierClientID": 0,
        "SupplierClientType": "Client33",
        "SupplierClientName": "My dummy client33"
    },
    "mvRecordAction": "Insert"
}`)

var (
	URL         string
	GET         string
	POS         string
	KEY         string
	HOST        string
	RequestBody *bytes.Buffer
)

func Load() {
	godotenv.Load(".env")

	URL = os.Getenv("URL")
	GET = os.Getenv("GET")
	POS = os.Getenv("POS")
	KEY = os.Getenv("KEY")
	HOST = os.Getenv("HOST")
	RequestBody = eRequestBody
}
