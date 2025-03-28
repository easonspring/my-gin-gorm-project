package models

import (
	"gorm.io/gorm"
) // User 用户模型定义
type User struct {
	gorm.Model
	Name  string `gorm:"size:255"`
	Email string `gorm:"uniqueIndex"`
	Age   int
}
