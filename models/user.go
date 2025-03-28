package models

import (
	"gorm.io/gorm"
) // User 用户模型定义

type User struct {
	gorm.Model
	Name  string `gorm:"size:255;comment:用户名"`
	Email string `gorm:"type:varchar(100);uniqueIndex;comment:邮箱"`
	Age   int    `gorm:"default:18;comment:年龄"`
}
