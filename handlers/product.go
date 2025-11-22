package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Homiez09/excel-io-example-backend/database"
	"github.com/Homiez09/excel-io-example-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

// ExportHandler - Streaming Logic
func ExportProduct(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename=products_export.xlsx")

	f := excelize.NewFile()
	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	sw.SetRow("A1", []interface{}{"Code", "Name", "Price", "Stock", "Updated At"})

	// เรียก DB จาก package database
	rows, err := database.DB.Model(&models.Product{}).Rows()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	rowIndex := 2
	for rows.Next() {
		var p models.Product
		database.DB.ScanRows(rows, &p)

		cell := fmt.Sprintf("A%d", rowIndex)
		values := []interface{}{p.Code, p.Name, p.Price, p.Stock, p.CreatedAt}

		if err := sw.SetRow(cell, values); err != nil {
			return c.Status(500).SendString("Write error")
		}
		rowIndex++
	}

	sw.Flush()
	return f.Write(c.Response().BodyWriter())
}

// ImportHandler - Batch Insert Logic
func ImportProduct(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).SendString("Upload failed")
	}

	src, err := fileHeader.Open()
	if err != nil {
		return c.Status(500).SendString("Open failed")
	}
	defer src.Close()

	f, err := excelize.OpenReader(src)
	if err != nil {
		return c.Status(500).SendString("Excel error")
	}

	rows, err := f.Rows("Sheet1")
	if err != nil {
		return c.Status(500).SendString("Sheet error")
	}

	var products []models.Product
	const BatchSize = 1000
	count := 0

	if rows.Next() {
		rows.Columns()
	} // Skip Header

	for rows.Next() {
		col, _ := rows.Columns()
		if len(col) >= 4 {
			price, _ := strconv.ParseFloat(col[2], 64)
			stock, _ := strconv.Atoi(col[3])

			products = append(products, models.Product{
				Code: col[0], Name: col[1], Price: price, Stock: stock, CreatedAt: time.Now(),
			})
		}

		if len(products) >= BatchSize {
			database.DB.CreateInBatches(products, BatchSize)
			count += len(products)
			products = nil
		}
	}

	if len(products) > 0 {
		database.DB.CreateInBatches(products, len(products))
		count += len(products)
	}

	return c.JSON(fiber.Map{"status": "success", "rows": count})
}
