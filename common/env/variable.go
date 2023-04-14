package env

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	URL          string
	GET          string
	POS          string
	GETINVENTORY string
	PARAM        string
	KEY          string
	HOST         string
	HOSTGIN      string
	RequestBody  []byte
)

func Load() {
	godotenv.Load(".env")

	URL = os.Getenv("URL")
	GET = os.Getenv("GET")
	POS = os.Getenv("POS")
	GETINVENTORY = os.Getenv("GETINVENTORY")
	PARAM = os.Getenv("PARAM")
	KEY = os.Getenv("KEY")
	HOST = os.Getenv("HOST")
	HOSTGIN = os.Getenv("HOSTGIN")
	RequestBody = requestBodyMapper()
}
