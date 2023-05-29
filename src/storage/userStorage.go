package storage

import (
	"errors"

	"../entities"
)

type IStorage[Entity any] interface {
	GetAll() ([]Entity, error)
	Create(Entity) error
	Delete(Entity) error
}

type UserMemoryStorage struct {
	Users []entities.User
}

func (this UserMemoryStorage) GetAll() ([]entities.User, error) {
	return this.Users, nil
}

func (this *UserMemoryStorage) Create(user entities.User) error {
	this.Users = append(this.Users, user)
	return nil
}

func (this *UserMemoryStorage) Delete(userToRemove entities.User) error {
	index := -1

	for idx, user := range this.Users {
		if user.Gmail == userToRemove.Gmail {
			index = idx
			break
		}
	}

	if index == -1 {
		return errors.New("User not found")
	}

	this.Users = append(this.Users[:index], this.Users[index+1])

	return nil
}
