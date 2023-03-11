package usecases

import "github.com/megalypse/golang-clean-arch/internal/domain/models"

type GetPersonById interface {
	GetPersonById(int) *models.Person
}

type CreatePerson interface {
	CreatePerson(models.Person) *models.Person
}

type FilterPeople interface {
	Filter(baseFilter *models.BaseFilter[models.Person]) models.Paginated[models.Person]
}

type PersonService interface {
	GetPersonById
	CreatePerson
	FilterPeople
	FilterPeople
}
