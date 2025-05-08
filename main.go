package main

import (
	"github.com/ThakdanaiDL/goex1.git/controller"
	"github.com/ThakdanaiDL/goex1.git/repository"
	"github.com/ThakdanaiDL/goex1.git/service"
)

func main() {
	varRepo := repository.NewHelloRepository()
	varService := service.NewHelloService(varRepo)
	varController := controller.NewHelloContrller(varService)

	varController.HandleGreeting()
}
