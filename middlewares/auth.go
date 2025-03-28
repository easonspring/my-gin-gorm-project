package middlewares

import (
	"net/http"
	"strings"

	"github.com/easonspring/my-gin-gorm-project/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求头中缺少Token"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的Token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
