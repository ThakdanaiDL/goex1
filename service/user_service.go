package service

import (
	"fmt"

	"github.com/ThakdanaiDL/goex1.git/repository" // เปลี่ยนเป็น import path ที่ถูกต้องของ repository
)

// UserServiceInterface กำหนด Interface สำหรับ Service Layer
type UserServiceInterface interface {
	GetAllUsers() ([]repository.User, error)
	CreateUser(name, email string) error
}

// userService struct สำหรับ implementation UserServiceInterface
type userService struct {
	userRepo repository.UserRepositoryInterface
}

// NewUserService สร้าง instance ของ userService
func NewUserService(userRepo repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{userRepo: userRepo}
}

// GetAllUsers เรียกใช้งาน Repository Layer เพื่อดึงข้อมูลผู้ใช้งานทั้งหมด
func (s *userService) GetAllUsers() ([]repository.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return users, nil
}

// CreateUser สร้างผู้ใช้งานใหม่และบันทึกลงไฟล์
func (s *userService) CreateUser(name, email string) error {
	if name == "" || email == "" {
		return fmt.Errorf("name and email cannot be empty")
	}

	newUser := repository.User{Name: name, Email: email}
	err := s.userRepo.Save(newUser)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}
