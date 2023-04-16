package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *router) setupHttp() {
	http.HandleFunc("/get", r.routeList.GetProductsController)
	http.HandleFunc("/post", r.routeList.PostProductsController)
	http.HandleFunc("/getinventory", r.routeList.GetInventoryController)
}

func (r *router) setupGin() {
	r.engine.Use(gin.Logger())

	r.engine.GET("/get/:sku", r.routeList.GetProductsById)
	r.engine.POST("/product", r.routeList.PostOneProduct)
	r.engine.POST("/client", r.routeList.PostClientInfos)
	r.engine.POST("/inventorylocation", r.routeList.PostInventoryLocation)
	r.engine.POST("/tax", r.routeList.PostTax)
	r.engine.POST("/discount", r.routeList.PostDiscount)
}
