package model

import (
	"github.com/gin-gonic/gin"
)

func (l ProductList) ConvertGinGet(id int) gin.H {
	var product gin.H

	for _, v := range l.MvProducts {
		if v.ProductID == id {
			product = gin.H{
				"sku":           v.ProductID,
				"description":   v.ProductDescription,
				"salesPrice":    v.ProductSellingPrice,
				"purchasePrice": v.ProductPurchasePrice,
			}
			break
		}
	}

	return gin.H{"products": product}
}
