package collectors

import (
	"log"
	"sync"

	"github.com/azar-intelops/websockets/pkg/websockets/server/models"
	"github.com/azar-intelops/websockets/pkg/websockets/server/services"
	"github.com/gorilla/websocket"
)

type ChatController struct {
	clients     map[*websocket.Conn]bool
	clientsLock sync.Mutex
	chatService *services.MessageService
}

func NewChatController() *ChatController {
	return &ChatController{
		clients:     make(map[*websocket.Conn]bool),
		chatService: services.NewMessageService(),
	}
}

func (cs *ChatController) RegisterClient(client *websocket.Conn) {
	cs.clientsLock.Lock()
	defer cs.clientsLock.Unlock()

	cs.clients[client] = true

	messages := cs.chatService.GetMessages()
	for _, message := range messages {
		err := client.WriteJSON(message)
		if err != nil {
			log.Println("Error sending message to client:", err)
		}
	}
}

func (cs *ChatController) UnregisterClient(client *websocket.Conn) {
	cs.clientsLock.Lock()
	defer cs.clientsLock.Unlock()

	delete(cs.clients, client)
}

func (cs *ChatController) BroadcastMessage(message models.Message) {
	cs.clientsLock.Lock()
	defer cs.clientsLock.Unlock()

	for client := range cs.clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Println("Error sending message to client:", err)
		}
	}

	cs.chatService.SaveMessage(message)
}
