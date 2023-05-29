package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"megaventory/dto"
)

func (s sender) GetProductsById(c *gin.Context) {
	sku := c.Param("sku")
	id, _ := strconv.Atoi(sku)

	results, err := s.sendList.GetProductsByIdService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results.ConvertGinGet(id))
}

func (s sender) PostOneProduct(c *gin.Context) {
	var parseStruct dto.DtoProduct
	if ginBindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostOneProductService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostClientInfos(c *gin.Context) {
	var parseStruct dto.DtoSupplierClient
	if ginBindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostSupplierClientService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostInventoryLocation(c *gin.Context) {
	var parseStruct dto.DtoInvenoryLocation
	if ginBindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostInventoryLocationService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostTax(c *gin.Context) {
	var parseStruct dto.DtoTax
	if ginBindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostTaxService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}

func (s sender) PostDiscount(c *gin.Context) {
	var parseStruct dto.DtoDiscount
	if ginBindJSON(c, &parseStruct) {
		return
	}

	mapped := s.sendList.PostDiscountService(parseStruct)
	c.JSON(http.StatusOK, mapped)
}
