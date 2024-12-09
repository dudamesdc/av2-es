package config

import (
	"sync"

	models "github.com/dudamesdc/av2-es/src/model"
)

var (
	db   *Database
	once sync.Once
)

type Database struct {
	Users        []models.UserResponse
	Pets         []models.PetResponse
	Appointments []models.AppointmentResponse
}

// Retorna uma instância singleton do banco
func GetDatabase() *Database {
	once.Do(func() {
		db = &Database{
			Users:        []models.UserResponse{},
			Pets:         []models.PetResponse{},
			Appointments: []models.AppointmentResponse{},
		}
	})
	return db
}
