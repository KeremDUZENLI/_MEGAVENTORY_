package controller

import (
	"fmt"
	"megaventory/dto"
	"megaventory/service"
	"net/http"
	"strconv"

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

func (s sender) GetProductsController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.GetProductsService()
	errorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringGet())
}

func (s sender) PostProductsController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.PostProductsService()
	errorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringPost())
}

func (s sender) GetInventoryController(w http.ResponseWriter, r *http.Request) {
	results, err := s.sendList.GetInventoryService()
	errorCheck(w, err)
	fmt.Fprint(w, results.ConvertStringInventory())
}

// GIN ----------------------------------------------------------------
func (s sender) GetProductsById(c *gin.Context) {
	sku := c.Param("sku")
	id, _ := strconv.Atoi(sku)

	results, err := s.sendList.GetProductsServiceById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results.ConvertGinGet(id))
}

func (s sender) PostOneProduct(c *gin.Context) {
	var parseStruct dto.DtoProduct
	if bindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostOneProductService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostClientInfos(c *gin.Context) {
	var parseStruct dto.DtoSupplierClient
	if bindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostSupplierClientService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostInventoryLocation(c *gin.Context) {
	var parseStruct dto.DtoInvenoryLocation
	if bindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostInventoryLocationService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostTax(c *gin.Context) {
	var parseStruct dto.DtoTax
	if bindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostTaxService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostDiscount(c *gin.Context) {
	var parseStruct dto.DtoDiscount
	if bindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostDiscountService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

// HELP ----------------------------------------------------------------
func errorCheck(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func bindJSON(c *gin.Context, parseStruct any) bool {
	if err := c.ShouldBindJSON(&parseStruct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return true
	}
	return false
}
