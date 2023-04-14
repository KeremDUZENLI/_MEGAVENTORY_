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

	go router.Run(env.HOST)
	go router.RunGin(env.HOSTGIN)

	// channel to keep the program running
	done := make(chan bool)
	<-done
}

func settingValues() router.Router {
	repo := repository.NewRepository()
	serv := service.NewService(repo)
	cont := controller.NewController(serv)
	rout := router.NewRouter(cont)

	return rout
}
