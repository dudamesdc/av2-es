package routes

import (
	auth "github.com/dudamesdc/av2-es/src/auth"
	controllerU "github.com/dudamesdc/av2-es/src/controller/User"
	controllerP "github.com/dudamesdc/av2-es/src/controller/pet" 
	controllerA "github.com/dudamesdc/av2-es/src/controller/appointment" 
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Rotas públicas
	r.POST("/login", auth.Login)

	// Grupo protegido com autenticação JWT
	api := r.Group("/api")
	api.Use(auth.JWTAuthMiddleware()) // Middleware JWT aplicado ao grupo

	// Rotas para usuários
	api.GET("/users/:id", controllerU.GetUserByID)
	api.POST("/users", controllerU.CreateUser)
	api.PUT("/users/:id", controllerU.UpdateUser)
	api.DELETE("/users/:id", controllerU.DeleteUser)

	// Rota que lista todos os usuários (apenas admin pode acessar)
	api.GET("/users", func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(403, gin.H{"error": "access denied"})
			return
		}
		controllerU.GetAllUsers(c)
	})

	// Rotas para pets
	api.GET("/pets", func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(403, gin.H{"error": "access denied"})
			return
		}
		controllerP.FindAllPets(c)
	})
	api.GET("/pets/:id", controllerP.FindPetByID)
	api.POST("/pets", controllerP.CreatePet)
	api.PUT("/pets/:id", controllerP.UpdatePet)
	api.DELETE("/pets/:id", controllerP.DeletePet)

	// Rotas para agendamentos
	api.GET("/appointments", func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(403, gin.H{"error": "access denied"})
			return
		}
		controllerA.GetAllAppointments(c)
	})
	api.GET("/appointments/:id", controllerA.GetAppointment)
	api.POST("/appointments", controllerA.CreateAppointment)
	api.PUT("/appointments/:id", controllerA.UpdateAppointment)
	// api.DELETE("/appointments/:id", controllerA.DeleteAppointment)
}
