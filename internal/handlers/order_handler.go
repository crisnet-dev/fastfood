package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/crisnet-dev/fastfood/cmd/ws"
	"github.com/crisnet-dev/fastfood/internal/models"
	"github.com/crisnet-dev/fastfood/internal/services"
	"github.com/crisnet-dev/fastfood/internal/utils"
)

func OrderHandler(w http.ResponseWriter, r *http.Request) {

	var order models.Order

	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.HttpError(w, "Invalid JSON", 400)
		return
	}

	ws.NotifyAdmin(order)

	utils.HttpResponse(w, map[string]string{
		"message": "Order sent",
	}, 200)
}

func DeleteAllOrderHandler(w http.ResponseWriter, r *http.Request) {
	err := services.DeleteAllPendingOrderService()
	if err != nil {
		utils.HttpError(w, "Error to delete peding orders", 500)
		return
	}

	utils.HttpResponse(w, map[string]string{
		"message": "Order deleted",
	}, 200)
}
