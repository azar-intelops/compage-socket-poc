package daos

import (
	"errors"
	"poc/pkg/websocket/server/models"
	"sync"
)

type UserDAO struct {
	users map[int64]models.User
	mutex sync.RWMutex
}

func NewUserDaos() *UserDAO {
	return &UserDAO{
		users: make(map[int64]models.User),
	}
}

func (dao *UserDAO) Create(user models.User) (models.User, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	dao.users[user.Id] = user
	return user, nil
}

func (dao *UserDAO) GetById(id int64) (models.User, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	user, exists := dao.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (dao *UserDAO) Update(id int64, user models.User) (models.User, error) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	_, exists := dao.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	user.Id = id
	dao.users[id] = user

	return user, nil
}

func (dao *UserDAO) Delete(id int64) error {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()

	_, exists := dao.users[id]
	if !exists {
		return errors.New("user not found")
	}

	delete(dao.users, id)
	return nil
}

func (dao *UserDAO) List() ([]models.User, error) {
	var v []models.User
	for _, user := range dao.users {
		v = append(v, user)
	}
	return v, nil
}
