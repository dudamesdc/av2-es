package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role    string `json:"role"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Validação fictícia (substitua pela lógica real)
	if loginData.Role == "admin"  {
		token, err := GenerateToken(1, "admin")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}else{
		token, err := GenerateToken(2, "user") // Usuário com ID 2
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return

	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
}
