package main

import (
	"log"

	"github.com/Homiez09/excel-io-example-backend/database"
	"github.com/Homiez09/excel-io-example-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Connect to database
	database.Connect()

	// Setup fiber app
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // 100 MB
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	// 3. Setup Routes
	routes.SetupRoutes(app)

	// 4. Start Server
	log.Fatal(app.Listen(":8080"))
}
