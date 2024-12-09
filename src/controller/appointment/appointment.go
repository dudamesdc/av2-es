package controller

import (
	"net/http"
	"strconv"

	models "github.com/dudamesdc/av2-es/src/model"
	repositories "github.com/dudamesdc/av2-es/src/repository/appointment"
	"github.com/gin-gonic/gin"
)

// Cria um novo agendamento de consulta
func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	// Bind JSON to the appointment object
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Recupera o userID do contexto (deve vir do token ou de algum middleware)
	userID := c.GetInt("userID")

	// Chama o repositório para criar o agendamento
	appointmentResponse, err := repositories.CreateAppointment(appointment, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o AppointmentResponse (com ID)
	c.JSON(http.StatusCreated, appointmentResponse)
}

// Atualiza um agendamento de consulta
func UpdateAppointment(c *gin.Context) {
	var updatedAppointment models.Appointment
	// Bind JSON to the updatedAppointment object
	if err := c.ShouldBindJSON(&updatedAppointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Recupera o ID do agendamento e o userID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	userID := c.GetInt("userID")

	// Chama o repositório para atualizar o agendamento
	appointmentResponse, err := repositories.UpdateAppointment(id, updatedAppointment, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o AppointmentResponse
	c.JSON(http.StatusOK, appointmentResponse)
}

// Retorna os detalhes de um agendamento específico
func GetAppointment(c *gin.Context) {
	// Recupera o ID do agendamento e o userID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	userID := c.GetInt("userID")

	// Chama o repositório para obter o agendamento
	appointmentResponse, err := repositories.GetAppointmentByID(id, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o AppointmentResponse
	c.JSON(http.StatusOK, appointmentResponse)
}

// Retorna todos os agendamentos (somente para admin)
func GetAllAppointments(c *gin.Context) {
	role := c.GetString("role")

	// Verifica se o usuário tem permissão de admin
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can view all appointments"})
		return
	}

	// Chama o repositório para obter todos os agendamentos
	appointments, err := repositories.GetAllAppointments()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna todos os agendamentos
	c.JSON(http.StatusOK, appointments)
}
