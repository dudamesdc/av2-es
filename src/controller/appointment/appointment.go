package controller

import (
	"net/http"
	"strconv"

	"github.com/dudamesdc/av2-es/src/auth"
	"github.com/dudamesdc/av2-es/src/model"
	repository "github.com/dudamesdc/av2-es/src/repository/appointment"
	"github.com/gin-gonic/gin"
)

// Cria um novo agendamento
func CreateAppointments(c *gin.Context) {
	var Appointments model.Appointments
	// Bind JSON para o objeto Appointments
	if err := c.ShouldBindJSON(&Appointments); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Chama o repositório para criar o agendamento e obter o AppointmentsResponse (com ID)
	AppointmentsResponse, err := repository.CreateAppointments(Appointments)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, AppointmentsResponse)
	return
}

func DeleteAppointments(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Appointments ID"})
		return
	}
	_, err = auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	err = repository.DeleteAppointments(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointments deleted successfully"})
	return
}


func GetAllAppointmentss(c *gin.Context) {
	Claims, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	if Claims.Admin != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Appointmentss := repository.GetAllAppointments()
	c.JSON(http.StatusOK, Appointmentss)
	return
}


func GetAppointmentsByID(c *gin.Context) {
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Appointments_id := c.Param("id")
	id, err := strconv.Atoi(Appointments_id)
	Appointments, err := repository.GetAppointmentsByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o agendamento encontrado
	c.JSON(http.StatusOK, Appointments)
	return
}
