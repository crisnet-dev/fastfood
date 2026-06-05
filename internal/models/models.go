package models

type Product struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"image_url"`
	Quantity    int    `json:"quantity"`
}

type Order struct {
	Name     string    `json:"name"`
	Products []Product `json:"products"`
	Location string    `json:"location"`
	Time     string    `json:"time"`
}

type MessageWs struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type PendingOrder struct {
	Orders []Order `json:"orders"`
}
