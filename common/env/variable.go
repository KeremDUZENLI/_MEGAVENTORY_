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
	KEY          string
	HOST         string
	RequestBody  []byte
)

func Load() {
	godotenv.Load(".env")

	URL = os.Getenv("URL")
	GET = os.Getenv("GET")
	POS = os.Getenv("POS")
	GETINVENTORY = os.Getenv("GETINVENTORY")
	KEY = os.Getenv("KEY")
	HOST = os.Getenv("HOST")
	RequestBody = requestBodyMapper()
}
