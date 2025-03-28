package main

import (
	"github.com/easonspring/my-gin-gorm-project/config"
	"github.com/easonspring/my-gin-gorm-project/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// 解析配置文件
	config.LoadConfig()

	// 初始化数据库连接
	db = config.InitDB()

	// 设置Gin引擎
	r := gin.Default()

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务器
	r.Run(config.Config.ServerPort)
}
