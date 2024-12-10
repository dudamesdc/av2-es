package controller

import (
	"net/http"
	"strconv"

	"github.com/dudamesdc/av2-es/src/auth"
	"github.com/dudamesdc/av2-es/src/model"
	repository "github.com/dudamesdc/av2-es/src/repository/service"
	"github.com/gin-gonic/gin"
)

// Cria um novo usuário
func CreateService(c *gin.Context) {
	var Service model.Service
	// Bind JSON para o objeto Service
	if err := c.ShouldBindJSON(&Service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Chama o repositório para criar o usuário e obter o ServiceResponse (com ID)
	ServiceResponse, err := repository.CreateService(Service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o ServiceResponse, que inclui o ID
	c.JSON(http.StatusCreated, ServiceResponse)
	return
}

// Remove um usuário pelo ID
func DeleteService(c *gin.Context) {
	idStr := c.Param("id")

	// Convertendo o parâmetro de string para inteiro
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Service ID"})
		return
	}
	_, err = auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Chama o repositório para deletar o usuário
	err = repository.DeleteService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna uma mensagem de sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
	return
}

// Retorna todos os usuários
func GetAllServices(c *gin.Context) {
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Services := repository.GetAllServices()
	c.JSON(http.StatusOK, Services)
	return
}

// Retorna um usuário específico pelo ID
func GetServiceByID(c *gin.Context) {
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Service_id := c.Param("id")
	id, err := strconv.Atoi(Service_id)
	Service, err := repository.GetServiceByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o usuário encontrado
	c.JSON(http.StatusOK, Service)
	return
}
