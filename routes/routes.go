package routes

import (
	"github.com/Homiez09/excel-io-example-backend/container"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, c *container.Container) {
	api := app.Group("/api")

	api.Post("/import", c.ProductHandler.ImportProduct)
	api.Get("/export", c.ProductHandler.ExportProduct)
}
