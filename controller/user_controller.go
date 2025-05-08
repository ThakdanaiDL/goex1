package controller

import (
	"fmt"

	"github.com/ThakdanaiDL/goex1.git/service"
)

// UserController struct
type UserController struct {
	userService service.UserServiceInterface
}

// NewUserController สร้าง instance ของ UserController
func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

// ListUsers เรียกใช้งาน Service Layer เพื่อดึงรายชื่อผู้ใช้งานและแสดงผล
func (c *UserController) ListUsers() {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("--- User List ---")
	for _, user := range users {
		fmt.Printf("Name: %s, Email: %s\n", user.Name, user.Email)
	}
	fmt.Println("------------------")
}

// AddUser เรียกใช้งาน Service Layer เพื่อสร้างผู้ใช้งานใหม่
func (c *UserController) AddUser(name, email string) {
	err := c.userService.CreateUser(name, email)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User '%s' added successfully.\n", name)
}
