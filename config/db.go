package config

import (
	model "go_book_api/internal/infrastructure/gorm"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	dsn := os.Getenv("DATABASE_URL")

	// Retry logic for database connection
	maxRetries := 10
	retryDelay := 3 * time.Second

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to database")
			break
		}

		if i < maxRetries-1 {
			log.Printf("Failed to connect to database (attempt %d/%d): %v. Retrying in %v...", i+1, maxRetries, err, retryDelay)
			time.Sleep(retryDelay)
		} else {
			log.Fatal("Failed to connect to the database after all retries: ", err)
		}
	}

	// migrate the schema
	if err := DB.AutoMigrate(&model.Book{}, &model.User{}); err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

}
