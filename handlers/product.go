package handlers

import (
	"github.com/Homiez09/excel-io-example-backend/services"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// ExportProduct - Streaming Logic
func (h *ProductHandler) ExportProduct(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename=products_export.xlsx")

	f, err := h.service.ExportProducts()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return f.Write(c.Response().BodyWriter())
}

// ImportProduct - Batch Insert Logic
func (h *ProductHandler) ImportProduct(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString("Upload failed")
	}

	src, err := fileHeader.Open()
	if err != nil {
		return c.Status(500).SendString("Open failed")
	}
	defer src.Close()

	count, err := h.service.ImportProducts(src)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"status": "success", "rows": count})
}
