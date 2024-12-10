package controller

import (
	"net/http"
	"strconv"

	"github.com/dudamesdc/av2-es/src/auth"
	"github.com/dudamesdc/av2-es/src/model"
	repository "github.com/dudamesdc/av2-es/src/repository/service"
	"github.com/gin-gonic/gin"
)

func CreateService(c *gin.Context) {
	var Service model.Service

	if err := c.ShouldBindJSON(&Service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
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

	ServiceResponse, err := repository.CreateService(Service)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ServiceResponse)
	return
}

func DeleteService(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Service ID"})
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

	err = repository.DeleteService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
	return
}

func GetAllServices(c *gin.Context) {
	Claims, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	if Claims.Admin != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	Services := repository.GetAllServices()
	c.JSON(http.StatusOK, Services)
	return
}

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

	c.JSON(http.StatusOK, Service)
	return
}
