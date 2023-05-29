package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
