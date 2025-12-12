package services

import "linkfi-coding/models"

type PersonService interface {
	ProcessPeople(people []models.Person) map[string]models.PersonInfo
}

type personService struct{}

func NewPersonService() PersonService {
	return &personService{}
}

func (s *personService) ProcessPeople(people []models.Person) map[string]models.PersonInfo {
	result := make(map[string]models.PersonInfo)

	for _, person := range people {
		result[person.Name] = models.PersonInfo{
			Age:      person.Age,
			Birthday: person.Birthday,
		}
	}

	return result
}
