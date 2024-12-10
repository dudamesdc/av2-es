package controller

import (
	"net/http"
	"strconv"

	"github.com/dudamesdc/av2-es/src/auth"
	"github.com/dudamesdc/av2-es/src/model"
	repository "github.com/dudamesdc/av2-es/src/repository/pet"
	"github.com/gin-gonic/gin"
)

func CreatePet(c *gin.Context) {
	var Pet model.Pet

	if err := c.ShouldBindJSON(&Pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	PetResponse, err := repository.CreatePet(Pet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, PetResponse)
	return
}

func UpdatePet(c *gin.Context) {
	var updatedPet model.Pet

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

	PetResponse, err := repository.UpdatePet(id_Pet, updatedPet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, PetResponse)
	return
}

func DeletePet(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pet ID"})
		return
	}
	Claims, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	if Claims.Admin != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	err = repository.DeletePet(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet deleted successfully"})
	return
}

func GetAllPets(c *gin.Context) {
	Claims, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	if Claims.Admin != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Pets := repository.GetAllPets()
	c.JSON(http.StatusOK, Pets)
	return
}

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

	c.JSON(http.StatusOK, Pet)
	return
}
