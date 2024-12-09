package controller

import (
	"net/http"
	"strconv"

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

	// Chama o repositório para criar o usuário e obter o UserResponse (com ID)
	userResponse, err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o UserResponse, que inclui o ID
	c.JSON(http.StatusCreated, userResponse)
}

// Atualiza um usuário existente
func UpdateUser(c *gin.Context) {
	var updatedUser model.User
	// Bind JSON para o objeto updatedUser
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Recupera o userID e role do token JWT
	userID := c.GetInt("userID") // Recupera o userID do token JWT
	role := c.GetString("role")

	// Se o usuário não for admin, ele só pode atualizar suas próprias informações
	if role != "admin" && updatedUser.ID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You can only update your own information"})
		return
	}

	// Chama o repositório para atualizar o usuário
	userResponse, err := repository.UpdateUser(userID, updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o UserResponse
	c.JSON(http.StatusOK, userResponse)
}

// Remove um usuário pelo ID
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	// Convertendo o parâmetro de string para inteiro
	id, err := strconv.Atoi(idStr)
	// Recupera o userID e role do token JWT
	userID := c.GetInt("userID") // Recupera o userID do token JWT
	role := c.GetString("role")

	// Se o usuário não for admin, ele só pode excluir sua própria conta
	if role != "admin" && userID != id {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You can only delete your own account"})
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
}

// Retorna todos os usuários
func GetAllUsers(c *gin.Context) {
	users := repository.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// Retorna um usuário específico pelo ID
func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	// Convertendo o parâmetro de string para inteiro
	id, err := strconv.Atoi(idStr)
	// Chama o repositório para buscar o usuário
	user, err := repository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retorna o usuário encontrado
	c.JSON(http.StatusOK, user)
}
