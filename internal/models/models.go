package models

type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
}

type Order struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
	Location string    `json:"location"`
}
