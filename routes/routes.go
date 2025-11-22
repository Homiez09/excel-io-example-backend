package routes

import (
	"github.com/Homiez09/excel-io-example-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/import", handlers.ImportProduct)
	api.Get("/export", handlers.ExportProduct)
}
