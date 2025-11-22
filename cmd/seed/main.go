package main

import (
	"log"

	"github.com/Homiez09/excel-io-example-backend/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables")
	}

	// Connect to database
	database.Connect()

	// Seed the database
	database.SeedProducts()
}
