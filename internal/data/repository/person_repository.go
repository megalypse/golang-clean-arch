package repository

import "github.com/megalypse/golang-clean-arch/internal/domain/models"

type PersonRepository interface {
	GetPersonById(int) *models.Person
	CreatePerson(models.Person) int
	Filter(*models.BaseFilter[models.Person]) models.Paginated[models.Person]
}
