package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func httpErrorCheck(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ginBindJSON(c *gin.Context, parseStruct any) bool {
	if err := c.ShouldBindJSON(&parseStruct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not bind JSON"})
		return true
	}
	return false
}
