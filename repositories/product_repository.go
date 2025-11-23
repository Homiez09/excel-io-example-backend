package repositories

import (
	"database/sql"

	"github.com/Homiez09/excel-io-example-backend/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll() (*sql.Rows, error)
	ScanRow(rows *sql.Rows, dest interface{}) error
	CreateBatch(products []models.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll() (*sql.Rows, error) {
	return r.db.Model(&models.Product{}).Rows()
}

func (r *productRepository) ScanRow(rows *sql.Rows, dest interface{}) error {
	return r.db.ScanRows(rows, dest)
}

func (r *productRepository) CreateBatch(products []models.Product) error {
	return r.db.CreateInBatches(products, len(products)).Error
}
