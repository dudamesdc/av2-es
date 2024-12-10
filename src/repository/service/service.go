package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

// Cria um novo agendamento com ID incremental
func CreateService(Service model.Service) (model.ServiceResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Verificar se o e-mail já existe
	// for _, existingService := range db.Services {
	// 	if existingService.Email == Service.Email {
	// 		return model.ServiceResponse{}, errors.New("email already registered")
	// 	}
	// }

	// Incrementar o ID
	var newID int
	if len(db.Services) > 0 {
		newID = db.Services[len(db.Services)-1].ID + 1
	} else {
		newID = 1
	}

	// Criar o ServiceResponse com o ID
	ServiceResponse := model.ServiceResponse{
		ID:   newID,
		Name: Service.Name,
	}

	// Adicionar o ServiceResponse ao banco de dados
	db.Services = append(db.Services, ServiceResponse) // Armazenando Service]Response

	return ServiceResponse, nil
}

// Remove um usuário do banco de dados
func DeleteService(id int) error {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	for i, Service := range db.Services {
		if Service.ID == id {
			// Remover usuário pelo índice
			db.Services = append(db.Services[:i], db.Services[i+1:]...)
			return nil
		}
	}
	return errors.New("Service not found")
}

// Retorna todos os usuários
func GetAllServices() []model.ServiceResponse {
	db := config.GetDatabase()
	return db.Services
}

func GetServiceByID(id int) (model.ServiceResponse, error) {
	db := config.GetDatabase()
	for _, Service := range db.Services {
		if Service.ID == id {
			return Service, nil
		}
	}
	return model.ServiceResponse{}, errors.New("Service not found")
}
