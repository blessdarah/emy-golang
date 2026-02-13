package config

import (
	"go_book_api/internal/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// migrate the schema
	if err := DB.AutoMigrate(&model.Book{}, &model.User{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

}
