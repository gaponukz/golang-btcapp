package storage

import (
	"errors"

	"btcapp/src/entities"
)

type IStorage[Entity any] interface {
	GetAll() ([]Entity, error)
	Create(Entity) error
	Delete(Entity) error
}

type UserMemoryStorage struct {
	Users []entities.User
}

func (strg UserMemoryStorage) GetAll() ([]entities.User, error) {
	return strg.Users, nil
}

func (strg *UserMemoryStorage) Create(user entities.User) error {
	strg.Users = append(strg.Users, user)
	return nil
}

func (strg *UserMemoryStorage) Delete(userToRemove entities.User) error {
	index := -1

	for idx, user := range strg.Users {
		if user.Gmail == userToRemove.Gmail {
			index = idx
			break
		}
	}

	if index == -1 {
		return errors.New("user not found")
	}

	strg.Users = append(strg.Users[:index], strg.Users[index+1])

	return nil
}
