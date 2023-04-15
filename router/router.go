package router

import (
	"megaventory/controller"

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
	router.setupHttp()

	return router
}
