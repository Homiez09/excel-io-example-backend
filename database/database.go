package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Homiez09/excel-io-example-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Read from environment variables with fallback defaults
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "123456789")
	dbname := getEnv("DB_NAME", "io_example_db")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")
	timezone := getEnv("DB_TIMEZONE", "Asia/Bangkok")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto Migrate
	DB.AutoMigrate(&models.Product{})
	fmt.Println("✅ Database Connected & Migrated!")
}

// getEnv gets environment variable with fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
