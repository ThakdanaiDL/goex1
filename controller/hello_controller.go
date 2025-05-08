package controller

import (
	"fmt"

	"github.com/ThakdanaiDL/goex1.git/service"
)

type HelloController struct {
	helloService service.HelloServiceInterface
}

func NewHelloContrller(svc service.HelloServiceInterface) *HelloController {
	return &HelloController{helloService: svc}
}

func (c *HelloController) HandleGreeting() {
	greeting := c.helloService.GetGreeting()
	fmt.Println("Message from controller:", greeting)
}
