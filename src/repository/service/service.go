package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

func CreateService(Service model.Service) (model.ServiceResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	var newID int
	if len(db.Services) > 0 {
		newID = db.Services[len(db.Services)-1].ID + 1
	} else {
		newID = 1
	}

	ServiceResponse := model.ServiceResponse{
		ID:   newID,
		Name: Service.Name,
	}

	db.Services = append(db.Services, ServiceResponse)

	return ServiceResponse, nil
}

func DeleteService(id int) error {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	for i, Service := range db.Services {
		if Service.ID == id {

			db.Services = append(db.Services[:i], db.Services[i+1:]...)
			return nil
		}
	}
	return errors.New("Service not found")
}

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
