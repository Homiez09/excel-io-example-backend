package database

import (
	"fmt"
	"log"

	"github.com/Homiez09/excel-io-example-backend/models"
)

// SeedProducts inserts sample product data into the database
func SeedProducts() {
	products := []models.Product{
		{Code: "P001", Name: "MacBook Pro 16\"", Price: 89900, Stock: 15},
		{Code: "P002", Name: "iPhone 15 Pro Max", Price: 45900, Stock: 50},
		{Code: "P003", Name: "iPad Air M2", Price: 24900, Stock: 30},
		{Code: "P004", Name: "AirPods Pro (2nd gen)", Price: 8900, Stock: 100},
		{Code: "P005", Name: "Apple Watch Series 9", Price: 13900, Stock: 45},
		{Code: "P006", Name: "Magic Keyboard", Price: 3900, Stock: 75},
		{Code: "P007", Name: "Magic Mouse", Price: 2900, Stock: 80},
		{Code: "P008", Name: "Apple Pencil (2nd gen)", Price: 4500, Stock: 60},
		{Code: "P009", Name: "HomePod mini", Price: 3500, Stock: 40},
		{Code: "P010", Name: "AirTag (4 pack)", Price: 3900, Stock: 120},
		{Code: "P011", Name: "USB-C to Lightning Cable", Price: 790, Stock: 200},
		{Code: "P012", Name: "20W USB-C Power Adapter", Price: 790, Stock: 150},
		{Code: "P013", Name: "MagSafe Charger", Price: 1390, Stock: 90},
		{Code: "P014", Name: "Studio Display", Price: 49900, Stock: 10},
		{Code: "P015", Name: "Mac Studio", Price: 69900, Stock: 8},
		{Code: "P016", Name: "Mac mini M2", Price: 21900, Stock: 25},
		{Code: "P017", Name: "Apple TV 4K", Price: 5490, Stock: 55},
		{Code: "P018", Name: "Beats Studio Pro", Price: 11900, Stock: 35},
		{Code: "P019", Name: "iPhone 15", Price: 29900, Stock: 70},
		{Code: "P020", Name: "MacBook Air M2", Price: 39900, Stock: 40},
	}

	for _, product := range products {
		// Check if product already exists
		var existing models.Product
		if err := DB.Where("code = ?", product.Code).First(&existing).Error; err == nil {
			log.Printf("⏭️  Product %s already exists, skipping...", product.Code)
			continue
		}

		// Create new product
		if err := DB.Create(&product).Error; err != nil {
			log.Printf("❌ Failed to seed product %s: %v", product.Code, err)
		} else {
			fmt.Printf("✅ Seeded: %s - %s\n", product.Code, product.Name)
		}
	}

	fmt.Println("🌱 Seeding completed!")
}
