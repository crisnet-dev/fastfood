package routes

import (
	"net/http"

	"github.com/crisnet-dev/fastfood/cmd/ws"
	"github.com/crisnet-dev/fastfood/internal/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ws", ws.HandleConnections)

	mux.HandleFunc("POST /order", handlers.OrderHandler)

	RegisterProductRoutes(mux)
	RegisterTemplateRoutes(mux)

	return mux
}
