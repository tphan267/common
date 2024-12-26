package database

import (
	"fmt"
	"log"

	"github.com/tphan267/common/system"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func InitMySqlDB() {
	DB_PORT := system.Env("DB_PORT")
	if DB_PORT == "" {
		DB_PORT = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", system.Env("DB_USER"), system.Env("DB_PASS"), system.Env("DB_HOST"), DB_PORT, system.Env("DB_NAME"))

	// NOTE:
	// To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
	// To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(system.EnvAsInt("DB_LOG_LEVEL", 3))),
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Printf("Connect to MySQL Database: '%s'\n", system.Env("DB_NAME"))
	DB = db
}

func InitPostgresDB() {
	DB_PORT := system.Env("DB_PORT")
	if DB_PORT == "" {
		DB_PORT = "5432"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", system.Env("DB_HOST"), system.Env("DB_USER"), system.Env("DB_PASS"), system.Env("DB_NAME"), DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(system.EnvAsInt("DB_LOG_LEVEL", 3))),
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Printf("Connect to Postgres Database: '%s'\n", system.Env("DB_NAME"))
	DB = db
}
