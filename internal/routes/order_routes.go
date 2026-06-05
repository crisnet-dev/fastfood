package routes

import (
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/handlers"
)

func RegisterOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /order", handlers.OrderHandler)
	mux.HandleFunc("DELETE /order", handlers.DeleteAllOrderHandler)
}
