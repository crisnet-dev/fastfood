package routes

import (
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/handlers"
)

func RegisterProductRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /products", handlers.GetProductHandler)
	mux.HandleFunc("POST /product/upload", handlers.UploadProduc)
}
