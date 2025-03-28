package controllers

import (
	"net/http"

	"github.com/easonspring/my-gin-gorm-project/models"
	"github.com/easonspring/my-gin-gorm-project/services"

	"github.com/gin-gonic/gin"
)

// GetUsers 获取所有用户
func GetUsers(c *gin.Context) {
	users := services.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// CreateUser 创建新用户
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
