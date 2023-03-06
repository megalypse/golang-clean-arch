package service

import (
	"github.com/megalypse/golang-clean-arch/internal/data/repository"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
)

type PersonService struct {
	personRepository repository.PersonRepository
}

func (ps PersonService) GetPersonById(id int) models.Person {
	return ps.personRepository.GetPersonById(id)
}

func (ps PersonService) CreatePerson(person models.Person) models.Person {
	return ps.personRepository.CreatePerson(person)
}
