package env

import (
	"os"

	"github.com/joho/godotenv"

	"megaventory/dto"
)

var SavedParseStruct dto.DtoSupplierClientHttp

var (
	URL          string
	GET          string
	POS          string
	GETINVENTORY string
	PARAM        string
	APIKEY       string
	HOST         string
	HOSTGIN      string
)

func Load() {
	godotenv.Load(".env")

	URL = os.Getenv("URL")
	GET = os.Getenv("GET")
	POS = os.Getenv("POS")
	GETINVENTORY = os.Getenv("GETINVENTORY")
	PARAM = os.Getenv("PARAM")
	APIKEY = os.Getenv("APIKEY")
	HOST = os.Getenv("HOST")
	HOSTGIN = os.Getenv("HOSTGIN")
}
