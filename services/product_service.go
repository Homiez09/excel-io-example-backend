package services

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/Homiez09/excel-io-example-backend/models"
	"github.com/Homiez09/excel-io-example-backend/repositories"
	"github.com/xuri/excelize/v2"
)

type ProductService interface {
	ExportProducts() (*excelize.File, error)
	ImportProducts(file io.Reader) (int, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) ExportProducts() (*excelize.File, error) {
	f := excelize.NewFile()
	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		return nil, err
	}

	if err := sw.SetRow("A1", []interface{}{"ID", "Code", "Name", "Price", "Stock", "Updated At"}); err != nil {
		return nil, err
	}

	rows, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rowIndex := 2
	for rows.Next() {
		var p models.Product
		if err := s.repo.ScanRow(rows, &p); err != nil {
			return nil, err
		}

		cell := fmt.Sprintf("A%d", rowIndex)
		values := []interface{}{p.ID, p.Code, p.Name, p.Price, p.Stock, p.CreatedAt.Format("2006-01-02")}

		if err := sw.SetRow(cell, values); err != nil {
			return nil, err
		}
		rowIndex++
	}

	if err := sw.Flush(); err != nil {
		return nil, err
	}
	return f, nil
}

func (s *productService) ImportProducts(file io.Reader) (int, error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return 0, err
	}

	rows, err := f.Rows("Sheet1")
	if err != nil {
		return 0, err
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
			if err := s.repo.CreateBatch(products); err != nil {
				return count, err
			}
			count += len(products)
			products = nil
		}
	}

	if len(products) > 0 {
		if err := s.repo.CreateBatch(products); err != nil {
			return count, err
		}
		count += len(products)
	}

	return count, nil
}
