package server

import (
	"log"
	"net/http"
	"time"

	"github.com/azar-intelops/websockets/pkg/websockets/server/collectors"
	"github.com/azar-intelops/websockets/pkg/websockets/server/models"
	"github.com/gorilla/websocket"
)

type ChatServer struct {
	chatController *collectors.ChatController
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		chatController: collectors.NewChatController(),
	}
}

func (cs *ChatServer) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Add your security checks here to validate the request origin
			// For example, you can check the request's Origin or referer header
			// and return true only if it matches your allowed origins.
			// You can also perform authentication or other security checks.

			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	cs.chatController.RegisterClient(conn)

	for {
		var message models.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		message.Timestamp = time.Now()

		cs.chatController.BroadcastMessage(message)
	}

	cs.chatController.UnregisterClient(conn)
}
