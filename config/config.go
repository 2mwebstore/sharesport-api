package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Load .env only in development
	if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, continuing...")
		}
	}

	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT") // optional, default 3306
	dbName := os.Getenv("DB_NAME")

	if dbPort == "" {
		dbPort = "3306"
	}

	if dbUser == "" || dbHost == "" || dbName == "" {
		log.Fatal("❌ Missing one or more required database environment variables (DB_USER, DB_HOST, DB_NAME)")
	}

	// MySQL DSN string (password optional)
	var dsn string
	if dbPass != "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPass, dbHost, dbPort, dbName)
	} else {
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbHost, dbPort, dbName)
	}

	// Connect using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	log.Println("✅ Database connected successfully")
}
