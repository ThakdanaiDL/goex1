package service

import (
	"github.com/ThakdanaiDL/goex1.git/repository" // เปลี่ยนเป็น import path ที่ถูกต้อง
)

// HelloServiceInterface กำหนด Interface สำหรับ Service Layer
type HelloServiceInterface interface {
	GetGreeting() string
}

// helloService struct สำหรับ implementation HelloServiceInterface
type helloService struct {
	helloRepo repository.HelloRepositoryInterface
}

// NewHelloService สร้าง instance ของ helloService โดยรับ HelloRepositoryInterface
func NewHelloService(repo repository.HelloRepositoryInterface) HelloServiceInterface {
	return &helloService{helloRepo: repo}
}

// GetGreeting เรียกใช้งาน Repository Layer เพื่อดึงข้อความมาสร้างคำทักทาย
func (s *helloService) GetGreeting() string {
	message := s.helloRepo.GetMessage()
	greeting := "Greeting from Service: " + message
	return greeting
}
