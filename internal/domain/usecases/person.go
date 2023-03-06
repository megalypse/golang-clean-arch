package usecases

import "github.com/megalypse/golang-clean-arch/internal/domain/models"

type GetPersonById interface {
	GetPerson(int) models.Person
}

type GetAllPeople interface {
	GetAll() []models.Person
}

type GetByFilter interface {
	GetByFilter(models.Person) []models.Person
}

type CreatePerson interface {
	CreatePerson(models.Person) models.Person
}
