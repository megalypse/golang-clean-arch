package http

import (
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
	"github.com/megalypse/golang-clean-arch/internal/domain/usecases"
)

type PersonController struct {
	createPersonUsecase usecases.CreatePerson
}

func (pc PersonController) CreatePerson() models.Person {
	return pc.createPersonUsecase.CreatePerson(models.Person{})
}
