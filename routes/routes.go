package routes

import (
	"linkfi-coding/handlers"
	"linkfi-coding/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	personService := services.NewPersonService()
	personHandler := handlers.NewPersonHandler(personService)

	r.POST("/process-people", personHandler.ProcessPeople)

	return r
}
