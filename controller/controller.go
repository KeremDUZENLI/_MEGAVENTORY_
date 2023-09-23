package controller

import (
	"megaventory/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sender struct {
	sendList service.Holder
}

type Sender interface {
	GetProductsController(w http.ResponseWriter, r *http.Request)
	PostProductsController(w http.ResponseWriter, r *http.Request)
	GetInventoryController(w http.ResponseWriter, r *http.Request)

	GetProductsById(c *gin.Context)
	PostOneProduct(c *gin.Context)
	PostClientInfos(c *gin.Context)
	PostInventoryLocation(c *gin.Context)
	PostTax(c *gin.Context)
	PostDiscount(c *gin.Context)
}

func NewController(h service.Holder) Sender {
	return &sender{sendList: h}
}
