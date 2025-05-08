package controller

import (
	"fmt"

	"github.com/ThakdanaiDL/goex1.git/service"
)

type helloController struct {
	helloService service.HelloServiceInterface
}

func NewHelloContrller(helloService service.HelloServiceInterface) HelloControllerInterface {
	return &helloController{helloService}
}

func (c *helloController) HandleGreeting() {
	greeting := c.helloService.GetGreeting()
	fmt.Println("Message from controller:", greeting)
}
