package daos

import (
	"sync"

	"github.com/azar-intelops/websockets/pkg/websockets/server/models"
)

// MessageDAO handles the storage and retrieval of messages.
type MessageDAO struct {
    messages []models.Message
    mutex    sync.RWMutex
}

// NewMessageDAO creates a new instance of MessageDAO.
func NewMessageDAO() *MessageDAO {
    return &MessageDAO{}
}

// SaveMessage saves a message to the DAO.
func (dao *MessageDAO) SaveMessage(message models.Message) {
    dao.mutex.Lock()
    defer dao.mutex.Unlock()

    dao.messages = append(dao.messages, message)
}

// GetMessages returns all the saved messages.
func (dao *MessageDAO) GetMessages() []models.Message {
    dao.mutex.RLock()
    defer dao.mutex.RUnlock()

    return dao.messages
}
