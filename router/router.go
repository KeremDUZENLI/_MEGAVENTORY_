package router

import (
	"megaventory/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type router struct {
	routeList controller.Sender
	engine    *gin.Engine
}

type Router interface {
	Run(host string)
	RunGin(host string)
}

func NewRouter(s controller.Sender) Router {
	router := &router{routeList: s}
	router.setup()

	return router
}

func (r *router) Run(host string) {
	go func() {
		http.ListenAndServe(host, nil)
	}()
}

func (r *router) RunGin(host string) {
	go func() {
		r.engine = gin.New()
		r.setupGin()
		r.engine.Run(host)
	}()
}

func (r *router) setup() {
	http.HandleFunc("/get", r.routeList.GetProductsController)
	http.HandleFunc("/post", r.routeList.PostProductsController)
	http.HandleFunc("/getinventory", r.routeList.GetInventoryController)
}

func (r *router) setupGin() {
	r.engine.Use(gin.Logger())
	r.engine.GET("/get/:sku", r.routeList.GetProductsByIdController)
}
