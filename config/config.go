package config

import (
	"log"
	"os"

	"github.com/easonspring/my-gin-gorm-project/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Config struct {
	ServerPort string
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

// LoadConfig 解析配置文件
func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	Config.ServerPort = os.Getenv("SERVER_PORT")
	Config.DBUser = os.Getenv("DB_USER")
	Config.DBPassword = os.Getenv("DB_PASSWORD")
	Config.DBName = os.Getenv("DB_NAME")
	Config.DBHost = os.Getenv("DB_HOST")
	Config.DBPort = os.Getenv("DB_PORT")
}

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	dsn := Config.DBUser + ":" + Config.DBPassword + "@tcp(" + Config.DBHost + ":" + Config.DBPort + ")/" + Config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// 初始化数据库脚本
	DB = db
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		print("err:{}", err)
	}
	return db
}
