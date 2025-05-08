package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// User struct สำหรับเก็บข้อมูลผู้ใช้งาน
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserRepositoryInterface กำหนด Interface สำหรับ Repository Layer
type UserRepositoryInterface interface {
	GetAll() ([]User, error)
	Save(user User) error
}

// userRepository struct สำหรับ implementation UserRepositoryInterface
type userRepository struct {
	filePath string
}

// NewUserRepository สร้าง instance ของ userRepository
func NewUserRepository(filePath string) UserRepositoryInterface {
	return &userRepository{filePath: filePath}
}

// GetAll อ่านข้อมูลผู้ใช้งานทั้งหมดจากไฟล์
func (r *userRepository) GetAll() ([]User, error) {
	var users []User
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var user User
		line := scanner.Text()
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			fmt.Printf("warning: failed to unmarshal line '%s': %v\n", line, err)
			continue // ข้ามบรรทัดที่อ่านไม่ได้
		}
		users = append(users, user)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return users, nil
}

// Save บันทึกข้อมูลผู้ใช้งานลงในไฟล์
func (r *userRepository) Save(user User) error {
	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	userData, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %w", err)
	}

	_, err = file.WriteString(string(userData) + "\n")
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
