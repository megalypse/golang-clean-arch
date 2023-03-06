package service

import (
	"github.com/megalypse/golang-clean-arch/internal/data/repository"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
)

type PersonService struct {
	PersonRepository repository.PersonRepository
}

func (ps PersonService) GetPersonById(id int) *models.Person {
	return ps.PersonRepository.GetPersonById(id)
}

func (ps PersonService) CreatePerson(person models.Person) *models.Person {
	personId := ps.PersonRepository.CreatePerson(person)

	return ps.GetPersonById(personId)
}
