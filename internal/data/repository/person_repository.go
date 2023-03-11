package repository

import "github.com/megalypse/golang-clean-arch/internal/domain/models"

type PersonRepository interface {
	GetPersonById(int64) *models.Person
	CreatePerson(models.Person) int64
	Filter(*models.BaseFilter[models.Person]) models.Paginated[models.Person]
	Update(*models.Person) int64
}
