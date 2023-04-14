package controller

import (
	"fmt"
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
	GetProductsByIdController(c *gin.Context)
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
func (s sender) GetProductsByIdController(c *gin.Context) {
	sku := c.Param("sku")
	id, _ := strconv.Atoi(sku)

	results, err := s.sendList.GetProductsByIdService()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results.ConvertGinGet(id))
}

// HELP ----------------------------------------------------------------
func errorCheck(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
