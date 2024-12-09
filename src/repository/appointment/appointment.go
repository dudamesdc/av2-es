package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

// Cria um novo agendamento de consulta
func CreateAppointment(appointment model.Appointment, userID int) (model.AppointmentResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Verificar se o usuário (dono do pet) existe e se o pet existe
	userExists := false
	for _, user := range db.Users {
		if user.ID == userID {
			userExists = true
			break
		}
	}

	if !userExists {
		return model.AppointmentResponse{}, errors.New("user does not exist")
	}

	// Verificar se o pet existe e pertence ao usuário
	petExists := false
	for _, pet := range db.Pets {
		if pet.ID == appointment.Pet_id && pet.OwnerID == userID {
			petExists = true
			break
		}
	}

	if !petExists {
		return model.AppointmentResponse{}, errors.New("pet not found or does not belong to the user")
	}

	// Incrementar o ID do agendamento
	var newID int
	if len(db.Appointments) > 0 {
		newID = db.Appointments[len(db.Appointments)-1].ID + 1
	} else {
		newID = 1
	}

	// Criar o AppointmentResponse
	appointmentResponse := model.AppointmentResponse{
		ID:       newID,
		Date:     appointment.Date,
		Admin_id: appointment.Admin_id,
		Pet_id:   appointment.Pet_id,
		OwnerID:  userID,
	}

	// Adicionar o AppointmentResponse ao banco de dados
	db.Appointments = append(db.Appointments, appointmentResponse)

	return appointmentResponse, nil
}

// Atualiza os dados de um agendamento
func UpdateAppointment(id int, updatedAppointment model.Appointment, userID int) (model.AppointmentResponse, error) {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Verificar se o agendamento existe
	for i, appointment := range db.Appointments {
		if appointment.ID == id {
			// Verificar se o pet pertence ao usuário
			if appointment.OwnerID != userID {
				return model.AppointmentResponse{}, errors.New("you can only update your own appointment")
			}

			// Atualizar os dados do agendamento
			db.Appointments[i].Date = updatedAppointment.Date
			db.Appointments[i].Admin_id = updatedAppointment.Admin_id
			return db.Appointments[i], nil
		}
	}
	return model.AppointmentResponse{}, errors.New("appointment not found")
}

// Retorna um agendamento específico
func GetAppointmentByID(id int, userID int) (model.AppointmentResponse, error) {
	db := config.GetDatabase()

	for _, appointment := range db.Appointments {
		if appointment.ID == id {
			// Verifica se o agendamento pertence ao usuário
			if appointment.OwnerID == userID {
				return appointment, nil
			}
			return model.AppointmentResponse{}, errors.New("you can only access your own appointment")
		}
	}
	return model.AppointmentResponse{}, errors.New("appointment not found")
}

// Retorna todos os agendamentos (somente para admin)
func GetAllAppointments() ([]model.AppointmentResponse, error) {
	db := config.GetDatabase()
	return db.Appointments, nil
}

func DeleteAppointment(id int, userID int) error {
	db := config.GetDatabase()
	mu.Lock()
	defer mu.Unlock()

	// Encontrar o agendamento com o ID especificado
	for i, appointment := range db.Appointments {
		if appointment.ID == id {
			// Verifica se o agendamento pertence ao usuário
			if appointment.OwnerID != userID {
				return errors.New("you can only delete your own appointment")
			}

			// Deletar o agendamento removendo-o da lista
			db.Appointments = append(db.Appointments[:i], db.Appointments[i+1:]...)
			return nil
		}
	}

	return errors.New("appointment not found")
}
