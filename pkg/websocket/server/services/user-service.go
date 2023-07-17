package services

import (
	"poc/pkg/websocket/server/daos"
	"poc/pkg/websocket/server/models"
)

type UserService struct {
	userDao *daos.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		userDao: daos.NewUserDaos(),
	}
}

func (s *UserService) CreateItem(user models.User) (models.User, error) {
	return s.userDao.Create(user)
}

func (s *UserService) GetItem(id int64) (models.User, error) {
	return s.userDao.GetById(id)
}

func (s *UserService) UpdateItem(id int64, updatedItem models.User) (models.User, error) {
	return s.userDao.Update(id, updatedItem)
}

func (s *UserService) DeleteItem(id int64) error {
	return s.userDao.Delete(id)
}

func (s *UserService) ListItem() ([]models.User, error) {
	return s.userDao.List()
}
