package routes

import (
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/handlers"
)

func RegisterProductRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /product", handlers.GetProductHandler)
	mux.HandleFunc("POST /product/upload", handlers.UploadProduc)
}
