package service

import (
	"time"

	"github.com/megalypse/golang-clean-arch/internal/data/repository"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
)

type PersonService struct {
	personRepository repository.PersonRepository
}

func NewPersonService(personRepository repository.PersonRepository) PersonService {
	return PersonService{
		personRepository: personRepository,
	}
}

func (ps PersonService) Delete(id int64) {
	now := time.Now()
	targetPerson := models.Person{
		BaseEntity: models.BaseEntity{
			Id:        &id,
			DeletedAt: &now,
		},
	}

	ps.personRepository.Update(&targetPerson)
}

func (ps PersonService) GetAll() []models.Person {
	return ps.personRepository.GetAll()
}

func (ps PersonService) GetPersonById(id int64) *models.Person {
	return ps.personRepository.GetPersonById(id)
}

func (ps PersonService) CreatePerson(person models.Person) *models.Person {
	personId := ps.personRepository.CreatePerson(person)

	return ps.GetPersonById(personId)
}

func (ps PersonService) Filter(baseFilter *models.BaseFilter[models.Person]) models.Paginated[models.Person] {
	return ps.personRepository.Filter(baseFilter)
}

func (ps PersonService) Update(updatedPerson *models.Person) *models.Person {
	id := ps.personRepository.Update(updatedPerson)

	return ps.GetPersonById(id)
}

func (ps PersonService) Exists(id int64) bool {
	return ps.personRepository.Exists(id)
}
