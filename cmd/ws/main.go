package ws

import (
	"log"
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/config"
	"github.com/crisnet-dev/fastfood/internal/models"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var connections = make(map[*websocket.Conn]bool)
var pendentOrders = []models.Order{}

func updateAdminStatus() {
	var order = struct {
		Type         string `json:"type"`
		AdminCounter int    `json:"admin_counter"`
	}{
		AdminCounter: len(connections) - 1,
		Type:         "UpdateAdminCounter",
	}

	for conn := range connections {
		if err := conn.WriteJSON(order); err != nil {
			log.Println("Error to send message")
			delete(connections, conn)
		}
	}
}

func SendPendentOrdersToAdmin(ws *websocket.Conn) {
	orders := struct {
		Orders []models.Order `json:"orders"`
	}{Orders: pendentOrders}
	ws.WriteJSON(orders)
}

func isAdmin(r *http.Request) bool {
	env := config.GetEnv()
	passwordAdmin := r.URL.Query().Get("password")

	if env.AdminPassword == "" {
		log.Fatal("Admin password is empty, verify your .env file!")
		return false
	}

	if passwordAdmin == env.AdminPassword {
		return true
	}
	return false
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	if !isAdmin(r) {
		ws.Close()
		log.Println("Invalid Admin connection!")
		return
	}

	connections[ws] = true

	updateAdminStatus()

	if len(pendentOrders) != 0 {
		SendPendentOrdersToAdmin(ws)
		pendentOrders = []models.Order{}
	}

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			delete(connections, ws)
			updateAdminStatus()
			break
		}
	}
}

func NotifyAdmin(order models.Order) {
	if len(connections) == 0 {
		pendentOrders = append(pendentOrders, order)
		return
	}
	for conn := range connections {
		if err := conn.WriteJSON(order); err != nil {
			log.Println("Error to send message")
			delete(connections, conn)
			updateAdminStatus()
		}
	}
}
