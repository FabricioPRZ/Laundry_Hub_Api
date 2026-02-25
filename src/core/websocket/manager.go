package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type NotificationPayload struct {
	ID            int    `json:"id"`
	Message       string `json:"message"`
	Type          string `json:"type"`
	ReservationID *int   `json:"reservationId,omitempty"`
}

var GlobalHub *Hub

func InitWebSocket() {
	GlobalHub = NewHub()
	go GlobalHub.Run()
	log.Println("WebSocket Hub iniciado")
}

func HandleConnection(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No autenticado"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error al hacer upgrade WebSocket: %v", err)
		return
	}

	client := NewClient(userID.(int), GlobalHub, conn)
	GlobalHub.register <- client

	go client.WritePump()
	go client.ReadPump()
}

func SendNotificationToUser(userID int, payload NotificationPayload) {
	if GlobalHub == nil {
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error al serializar notificación: %v", err)
		return
	}

	GlobalHub.broadcast <- &Message{
		UserID:  userID,
		Payload: data,
	}
}
