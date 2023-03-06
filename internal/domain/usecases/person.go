package usecases

import "github.com/megalypse/golang-clean-arch/internal/domain/models"

type GetPersonById interface {
	GetPersonById(int) *models.Person
}

type CreatePerson interface {
	CreatePerson(models.Person) *models.Person
}
