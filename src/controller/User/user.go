package controller

import (
	"net/http"
	"strconv"

	"github.com/dudamesdc/av2-es/src/auth"
	"github.com/dudamesdc/av2-es/src/model"
	repository "github.com/dudamesdc/av2-es/src/repository/user"
	"github.com/gin-gonic/gin"
)

// Cria um novo usuário
func CreateUser(c *gin.Context) {
	var user model.User
	// Bind JSON para o objeto user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Chama o repositório para criar o usuário e obter o UserResponse (com ID)
	userResponse, err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o UserResponse, que inclui o ID
	c.JSON(http.StatusCreated, userResponse)
	return
}

// Atualiza um usuário existente
func UpdateUser(c *gin.Context) {
	var updatedUser model.User
	// Bind JSON para o objeto updatedUser
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	userID := c.Param("id")
	id_user, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	// Chama o repositório para atualizar o usuário
	userResponse, err := repository.UpdateUser(id_user, updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o UserResponse
	c.JSON(http.StatusOK, userResponse)
	return
}

// Remove um usuário pelo ID
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	// Convertendo o parâmetro de string para inteiro
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	_, err = auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}

	// Chama o repositório para deletar o usuário
	err = repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna uma mensagem de sucesso
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	return
}

// Retorna todos os usuários
func GetAllUsers(c *gin.Context) {
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	users := repository.GetAllUsers()
	c.JSON(http.StatusOK, users)
	return
}

// Retorna um usuário específico pelo ID
func GetUserByID(c *gin.Context) {
	_, err := auth.ValidateToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
		return
	}
	user_id := c.Param("id")
	id, err := strconv.Atoi(user_id)
	user, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o usuário encontrado
	c.JSON(http.StatusOK, user)
	return
}
