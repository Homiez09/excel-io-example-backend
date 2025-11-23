package database

import (
	"fmt"
	"log"

	"github.com/Homiez09/excel-io-example-backend/config"
	"github.com/Homiez09/excel-io-example-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode, cfg.DBTimeZone)

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
