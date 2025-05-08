package main

import (
	"github.com/ThakdanaiDL/goex1.git/controller"
	"github.com/ThakdanaiDL/goex1.git/repository"
	"github.com/ThakdanaiDL/goex1.git/service"
)

func main() {
	// กำหนด path ของไฟล์
	filePath := "users.txt"

	// สร้าง instance ของ Repository Layer
	userRepo := repository.NewUserRepository(filePath)

	// สร้าง instance ของ Service Layer โดย Inject UserRepository
	userService := service.NewUserService(userRepo)

	// สร้าง instance ของ Controller Layer โดย Inject UserService
	userController := controller.NewUserController(userService)

	// ตัวอย่างการใช้งาน Controller
	userController.ListUsers()

	userController.AddUser("Alice", "alice@example.com")
	userController.AddUser("Bob", "bob@example.com")

	userController.ListUsers()
}
