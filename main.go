package main

import (
	"megaventory/common/env"
	"megaventory/controller"
	"megaventory/repository"
	"megaventory/router"
	"megaventory/service"
)

func main() {
	env.Load()
	router := settingValues()

	router.Run(env.HOST)
}

func settingValues() router.Router {
	repo := repository.NewRepository()
	serv := service.NewService(repo)
	cont := controller.NewController(serv)
	rout := router.NewRouter(cont)

	return rout
}
