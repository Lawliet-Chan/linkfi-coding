package handlers

import (
	"net/http"

	"linkfi-coding/models"
	"linkfi-coding/services"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	personService services.PersonService
}

func NewPersonHandler(personService services.PersonService) *PersonHandler {
	return &PersonHandler{
		personService: personService,
	}
}

func (h *PersonHandler) ProcessPeople(c *gin.Context) {
	var people []models.Person

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format: " + err.Error(),
		})
		return
	}

	result := h.personService.ProcessPeople(people)
	c.JSON(http.StatusOK, result)
}
