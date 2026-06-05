package ws

import (
	"log"
	"net/http"

	"github.com/crisnet-dev/fastfood/internal/config"
	"github.com/crisnet-dev/fastfood/internal/models"
	"github.com/crisnet-dev/fastfood/internal/repository"
	"github.com/crisnet-dev/fastfood/internal/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var connections = make(map[*websocket.Conn]bool)

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
	orders, err := repository.GetAllPendingOrderRepository()
	if err != nil {
		log.Println(err)
		msg := models.MessageWs{Type: "Error", Message: "Error to proccess pendents orders!"}
		ws.WriteJSON(msg)
		return
	}
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
	if !isAdmin(r) {
		log.Println("Invalid Admin connection!")
		utils.HttpError(w, "Invalid Admin!", 401)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	connections[ws] = true

	updateAdminStatus()

	SendPendentOrdersToAdmin(ws)

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
		repository.AddNewPendingOrderRepository(&order)
		return
	}
	for conn := range connections {
		if err := conn.WriteJSON(order); err != nil {
			log.Println("Error to send message")
			delete(connections, conn)
			conn.Close()
			updateAdminStatus()
		}
	}
}
