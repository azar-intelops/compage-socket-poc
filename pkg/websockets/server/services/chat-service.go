package services

import (
	"sync"

	"github.com/azar-intelops/websockets/pkg/websockets/server/daos"
	"github.com/azar-intelops/websockets/pkg/websockets/server/models"
)

type MessageService struct {
	messageDAO *daos.MessageDAO
	mutex      sync.Mutex
}

func NewMessageService() *MessageService {
	return &MessageService{
		messageDAO: daos.NewMessageDAO(),
	}
}

func (s *MessageService) SaveMessage(message models.Message) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.messageDAO.SaveMessage(message)
}

func (s *MessageService) GetMessages() []models.Message {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.messageDAO.GetMessages()
}
