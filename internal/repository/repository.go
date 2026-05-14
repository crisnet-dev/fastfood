package repository

import (
	"log"

	"github.com/crisnet-dev/fastfood/internal/models"
)

func AddNewProductRepository(product *models.Product) {
	log.Println(*product)
}

func GetProduct() []models.Product { return nil }
