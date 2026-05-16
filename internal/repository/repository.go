package repository

import (
	"github.com/crisnet-dev/fastfood/internal/database"
	"github.com/crisnet-dev/fastfood/internal/models"
)

func AddNewProductRepository(product *models.Product) error {
	_, err := database.DB.Exec("INSERT INTO products (product_name, price, image_url) VALUES ($1, $2, $3)", product.ProductName, product.Price, product.ImageURL)
	if err != nil {
		return err
	}
	return nil
}

func GetProduct() ([]models.Product, error) {
	rows, err := database.DB.Query(`
		SELECT id, product_name, price, image_url
		FROM products
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`, 10, 0)
	if err != nil {
		return nil, err
	}

	var products []models.Product = []models.Product{}
	var product models.Product

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.ProductName, &product.Price, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
