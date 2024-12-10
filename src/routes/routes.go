package routes

import (
	auth "github.com/dudamesdc/av2-es/src/auth"
	controllerU "github.com/dudamesdc/av2-es/src/controller/User"
	controllerA "github.com/dudamesdc/av2-es/src/controller/appointment"
	controllerP "github.com/dudamesdc/av2-es/src/controller/pet"
	controllerS "github.com/dudamesdc/av2-es/src/controller/service"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	r.POST("/login", auth.Login)

	r.GET("/users/:id", controllerU.GetUserByID)
	r.POST("/users", controllerU.CreateUser)
	r.PUT("/users/:id", controllerU.UpdateUser)
	r.DELETE("/users/:id", controllerU.DeleteUser)
	r.GET("/users", controllerU.GetAllUsers)

	r.GET("/pets", controllerP.GetAllPets)
	r.GET("/pets/:id", controllerP.GetPetByID)
	r.POST("/pets", controllerP.CreatePet)
	r.PUT("/pets/:id", controllerP.UpdatePet)
	r.DELETE("/pets/:id", controllerP.DeletePet)

	r.GET("/appointments", controllerA.GetAllAppointmentss)
	r.GET("/appointments/:id", controllerA.GetAppointmentsByID)
	r.POST("/appointments", controllerA.CreateAppointments)
	// r.PUT("/appointments/:id", controllerA.UpdateAppointment)

	r.GET("/service", controllerS.GetAllServices)
	r.GET("/service/:id", controllerS.GetServiceByID)
	r.POST("/service", controllerS.CreateService)
	r.DELETE("/service/:id", controllerS.DeleteService)
}
