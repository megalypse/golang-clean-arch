package usecases

import "github.com/megalypse/golang-clean-arch/internal/domain/models"

type GetPersonById interface {
	GetPersonById(int64) *models.Person
}

type CreatePerson interface {
	CreatePerson(models.Person) *models.Person
}

type FilterPeople interface {
	Filter(*models.BaseFilter[models.Person]) models.Paginated[models.Person]
}

type UpdatePerson interface {
	Update(*models.Person) *models.Person
}

type PersonExists interface {
	Exists(int64) bool
}

type GetAll interface {
	GetAll() []models.Person
}

type SoftDeletePerson interface {
	SoftDelete(int64)
}

type IsNotDeleted interface {
	IsNotDeleted(int64) bool
}

type PersonService interface {
	IsNotDeleted
	GetPersonById
	CreatePerson
	FilterPeople
	FilterPeople
	UpdatePerson
	PersonExists
	GetAll
	SoftDeletePerson
}
