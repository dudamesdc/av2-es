package repositories

import (
	"errors"
	"sync"

	"github.com/dudamesdc/av2-es/src/config"
	model "github.com/dudamesdc/av2-es/src/model"
)

var mu sync.Mutex // Mutex para evitar condições de corrida

func CreateAppointments(Appointments model.Appointments) (model.AppointmentsResponse, error) {
	db := config.GetDatabase()
	mu.Lock()

	var newID int
	if len(db.Appointments) > 0 {
		newID = db.Appointments[len(db.Appointments)-1].ID + 1
	} else {
		newID = 1
	}

	AppointmentsResponse := model.AppointmentsResponse{
		ID:        newID,
		Date:      Appointments.Date,
		ServiceID: Appointments.ServiceID,
	}

	db.Appointments = append(db.Appointments, AppointmentsResponse) //

	return AppointmentsResponse, nil
}

func DeleteAppointments(id int) error {
	db := config.GetDatabase()

	for i, Appointments := range db.Appointments {
		if Appointments.ID == id {

			db.Appointments = append(db.Appointments[:i], db.Appointments[i+1:]...)
			return nil
		}
	}
	return errors.New("Appointments not found")
}

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
