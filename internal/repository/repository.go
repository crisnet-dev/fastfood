package repository

import (
	"encoding/json"

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

func DeleAllProductsRepository() error {
	_, err := database.DB.Exec("DELETE * FROM products")
	if err != nil {
		return err
	}
	return nil
}

func AddNewPendingOrderRepository(order *models.Order) error {
	jsonData, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(`INSERT INTO pending_orders (orders) VALUES ($1)`, jsonData)
	if err != nil {
		return err
	}

	return nil
}

func GetAllPendingOrderRepository() (models.PendingOrder, error) {
	rows, err := database.DB.Query(`SELECT orders FROM pending_orders`)
	if err != nil {
		return models.PendingOrder{}, err
	}

	var orders []models.Order = []models.Order{}

	for rows.Next() {
		var order models.Order = models.Order{}
		var productsBytes []byte

		err := rows.Scan(&productsBytes)
		if err != nil {
			return models.PendingOrder{}, err
		}

		err = json.Unmarshal(productsBytes, &order)
		if err != nil {
			return models.PendingOrder{}, err
		}

		orders = append(orders, order)
	}

	data := models.PendingOrder{
		Orders: orders,
	}

	return data, nil
}

func DeleteAllPendingOrders() error {
	_, err := database.DB.Exec("TRUNCATE TABLE pending_orders;")
	if err != nil {
		return err
	}
	return nil
}
