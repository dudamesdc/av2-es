package controller

import (
	"net/http"
	"strconv"

	"github.com/dudamesdc/av2-es/src/auth"
	"github.com/dudamesdc/av2-es/src/model"
	repository "github.com/dudamesdc/av2-es/src/repository/pet"
	"github.com/gin-gonic/gin"
)

// Cria um novo usuário
func CreatePet(c *gin.Context) {
	var Pet model.Pet
	// Bind JSON para o objeto Pet
	if err := c.ShouldBindJSON(&Pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Chama o repositório para criar o usuário e obter o PetResponse (com ID)
	PetResponse, err := repository.CreatePet(Pet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o PetResponse, que inclui o ID
	c.JSON(http.StatusCreated, PetResponse)
	return
}

// Atualiza um usuário existente
func UpdatePet(c *gin.Context) {
	var updatedPet model.Pet
	// Bind JSON para o objeto updatedPet
	if err := c.ShouldBindJSON(&updatedPet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	PetID := c.Param("id")
	id_Pet, err := strconv.Atoi(PetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pet ID"})
		return
	}
	// Chama o repositório para atualizar o usuário
	PetResponse, err := repository.UpdatePet(id_Pet, updatedPet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o PetResponse
	c.JSON(http.StatusOK, PetResponse)
	return
}

// Remove um usuário pelo ID
func DeletePet(c *gin.Context) {
	idStr := c.Param("id")

	// Convertendo o parâmetro de string para inteiro
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pet ID"})
		return
	}
	_, err = auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Chama o repositório para deletar o usuário
	err = repository.DeletePet(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna uma mensagem de sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully"})
	return
}

// Retorna todos os usuários
func GetAllPets(c *gin.Context) {
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Pets := repository.GetAllPets()
	c.JSON(http.StatusOK, Pets)
	return
}

// Retorna um usuário específico pelo ID
func GetPetByID(c *gin.Context) {
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Pet_id := c.Param("id")
	id, err := strconv.Atoi(Pet_id)
	Pet, err := repository.GetPetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o usuário encontrado
	c.JSON(http.StatusOK, Pet)
	return
}
