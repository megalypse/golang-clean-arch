package service

import (
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
