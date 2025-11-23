package container

import (
	"github.com/Homiez09/excel-io-example-backend/handlers"
	"github.com/Homiez09/excel-io-example-backend/repositories"
	"github.com/Homiez09/excel-io-example-backend/services"
	"gorm.io/gorm"
)

type Container struct {
	ProductHandler *handlers.ProductHandler
}

func NewContainer(db *gorm.DB) *Container {
	// Repositories
	productRepo := repositories.NewProductRepository(db)

	// Services
	productService := services.NewProductService(productRepo)

	// Handlers
	productHandler := handlers.NewProductHandler(productService)

	return &Container{
		ProductHandler: productHandler,
	}
}
