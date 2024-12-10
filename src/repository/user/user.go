package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

func CreateUser(user model.User) (model.UserResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	for _, existingUser := range db.Users {
		if existingUser.Email == user.Email {
			return model.UserResponse{}, errors.New("email already registered")
		}
	}

	var newID int
	if len(db.Users) > 0 {
		newID = db.Users[len(db.Users)-1].ID + 1
	} else {
		newID = 1
	}

	userResponse := model.UserResponse{
		ID:       newID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Admin:    user.Admin,
	}

	db.Users = append(db.Users, userResponse)

	return userResponse, nil
}

func UpdateUser(id int, updatedUser model.User) (model.UserResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	for i, user := range db.Users {
		if user.ID == id {

			db.Users[i].ID = updatedUser.ID
			db.Users[i].Name = updatedUser.Name
			db.Users[i].Email = updatedUser.Email
			db.Users[i].Password = updatedUser.Password
			db.Users[i].Admin = updatedUser.Admin
			return db.Users[i], nil
		}
	}
	return model.UserResponse{}, errors.New("user not found")
}

func DeleteUser(id int) error {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	for i, user := range db.Users {
		if user.ID == id {

			db.Users = append(db.Users[:i], db.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func GetAllUsers() []model.UserResponse {
	db := config.GetDatabase()
	return db.Users
}

func GetUserByID(id int) (model.UserResponse, error) {
	db := config.GetDatabase()
	for _, user := range db.Users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.UserResponse{}, errors.New("user not found")
}
