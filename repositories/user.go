package repositories

import (
	"github.com/easonspring/my-gin-gorm-project/config"
	"github.com/easonspring/my-gin-gorm-project/models"
)

// var DB *gorm.DB

// // SetDB 设置数据库实例
// func SetDB(db *gorm.DB) {
// 	DB = config.DB
// }

// GetAllUsers 获取所有用户
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser 创建新用户
func CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
