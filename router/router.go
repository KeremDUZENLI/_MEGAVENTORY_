package router

import (
	"megaventory/controller"
	"net/http"
)

type router struct {
	routeList controller.Sender
}

type Router interface {
	Run(host string)
}

func NewRouter(s controller.Sender) Router {
	router := &router{routeList: s}
	router.setup()

	return router
}

func (r *router) Run(host string) {
	http.ListenAndServe(host, nil)
}

func (r *router) setup() {
	http.HandleFunc("/get", r.routeList.GetProductsController)
	http.HandleFunc("/post", r.routeList.PostProductsController)
	http.HandleFunc("/getinventory", r.routeList.GetInventoryController)
}
