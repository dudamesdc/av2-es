package config

import (
	"fmt"
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
	Appointments []models.AppointmentsResponse
	Services     []models.ServiceResponse
}

// Retorna uma instância singleton do banco
func GetDatabase() *Database {
	once.Do(func() {
		db = &Database{
			Users:        []models.UserResponse{},
			Pets:         []models.PetResponse{},
			Appointments: []models.AppointmentsResponse{},
			Services:     []models.ServiceResponse{},
		}
	})
	return db
}

func (db *Database) GetUserIDByEmail(email string) (int, error) {
	for _, user := range db.Users {
		if user.Email == email { // Supondo que o modelo UserResponse tenha um campo Email
			return user.ID, nil // Supondo que UserResponse tenha um campo ID
		}
	}
	return 0, fmt.Errorf("usuário com o e-mail '%s' não encontrado", email)
}
