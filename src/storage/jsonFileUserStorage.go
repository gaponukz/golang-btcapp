package storage

import (
	"btcapp/src/entities"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type JsonFileUserStorage struct {
	Filename string
}

func (strg *JsonFileUserStorage) GetAll() ([]entities.User, error) {
	jsonFile, err := os.Open(strg.Filename)

	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	users := []entities.User{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)

	return users, nil
}

func (strg *JsonFileUserStorage) Create(user entities.User) error {
	users, err := strg.GetAll()

	if err != nil {
		return err
	}

	users = append(users, user)
	err = strg.writeUsers(users)

	return err
}

func (strg *JsonFileUserStorage) writeUsers(users []entities.User) error {
	usersJSON, err := json.MarshalIndent(users, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(strg.Filename, usersJSON, 0644)

	return err
}

func (strg *JsonFileUserStorage) Delete(userToRemove entities.User) error {
	users, err := strg.GetAll()
	index := -1

	if err != nil {
		return err
	}

	for idx, user := range users {
		if user.Gmail == userToRemove.Gmail {
			index = idx
			break
		}
	}

	if index == -1 {
		return errors.New("user not found")
	}

	users = append(users[:index], users[index+1:]...)
	err = strg.writeUsers(users)

	return err
}
