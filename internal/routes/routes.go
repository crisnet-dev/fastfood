package routes

import (
	"net/http"

	"github.com/crisnet-dev/fastfood/cmd/ws"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ws", ws.HandleConnections)

	RegisterProductRoutes(mux)
	RegisterTemplateRoutes(mux)
	RegisterOrderRoutes(mux)

	return mux
}
