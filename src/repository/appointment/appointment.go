package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

// Cria um novo agendamento com ID incremental
func CreateAppointments(Appointments model.Appointments) (model.AppointmentsResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	// Incrementar o ID
	var newID int
	if len(db.Appointments) > 0 {
		newID = db.Appointments[len(db.Appointments)-1].ID + 1
	} else {
		newID = 1
	}

	// Criar o AppointmentsResponse com o ID
	AppointmentsResponse := model.AppointmentsResponse{
		ID:        newID,
		Date:      Appointments.Date,
		ServiceID: Appointments.ServiceID,
	}

	// Adicionar o AppointmentsResponse ao banco de dados
	db.Appointments = append(db.Appointments, AppointmentsResponse) // Armazenando AppointmentsResponse

	return AppointmentsResponse, nil
}

// Atualiza os dados de um usuário existente
// func UpdateAppointments(id int, updatedAppointments model.Appointments) (model.AppointmentsResponse, error) {
// 	db := config.GetDatabase()
// 	mu.Lock()
// 	defer mu.Unlock()

// 	for i, Appointments := range db.Appointments {
// 		if Appointments.ID == id {
// 			// Atualizar apenas os campos permitidos
// 			db.Appointments[i].ID = updatedAppointments.
// 			db.Appointments[i].Date = updatedAppointments.Date

// 			return db.Appointments[i], nil
// 		}
// 	}
// 	return model.AppointmentsResponse{}, errors.New("Appointments not found")
// }

// Remove um usuário do banco de dados
func DeleteAppointments(id int) error {
	db := config.GetDatabase()

	for i, Appointments := range db.Appointments {
		if Appointments.ID == id {
			// Remover usuário pelo índice
			db.Appointments = append(db.Appointments[:i], db.Appointments[i+1:]...)
			return nil
		}
	}
	return errors.New("Appointments not found")
}

// Retorna todos os usuários
func GetAllAppointments() []model.AppointmentsResponse {
	db := config.GetDatabase()
	return db.Appointments
}

func GetAppointmentsByID(id int) (model.AppointmentsResponse, error) {
	db := config.GetDatabase()
	for _, Appointments := range db.Appointments {
		if Appointments.ID == id {
			return Appointments, nil
		}
	}
	return model.AppointmentsResponse{}, errors.New("Appointments not found")
}
