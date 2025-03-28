package services

import (
	"github.com/easonspring/my-gin-gorm-project/models"
	"github.com/easonspring/my-gin-gorm-project/repositories"
)

// GetAllUsers 获取所有用户
func GetAllUsers() []models.User {
	users, _ := repositories.GetAllUsers()
	return users
}

// CreateUser 创建新用户
func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}
