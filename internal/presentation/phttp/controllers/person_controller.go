package controllers

import (
	"net/http"
	"strconv"

	"github.com/megalypse/golang-clean-arch/internal/domain/models"
	"github.com/megalypse/golang-clean-arch/internal/domain/usecases"
	"github.com/megalypse/golang-clean-arch/internal/presentation/phttp"
)

type personController struct {
	createPersonUseCase  usecases.CreatePerson
	getPersonByIdUseCase usecases.GetPersonById
	filterPeopleUsecase  usecases.FilterPeople
}

func NewPersonController(personService usecases.PersonService) personController {
	return personController{
		createPersonUseCase:  personService,
		getPersonByIdUseCase: personService,
		filterPeopleUsecase:  personService,
	}
}

func (pc personController) GetHandlers() map[string]phttp.RouteDefinition {
	return map[string]phttp.RouteDefinition{
		"/person": {
			Method:       phttp.POST,
			HandlingFunc: pc.createPerson,
		},
		"/person/{personId}": {
			Method:       phttp.GET,
			HandlingFunc: pc.getPersonById,
		},
		"/person/filter": {
			Method:       phttp.POST,
			HandlingFunc: pc.filter,
		},
	}
}

func (pc personController) createPerson(w http.ResponseWriter, r *http.Request) {
	person, err := phttp.ParseBody[models.Person](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdPerson := pc.createPersonUseCase.CreatePerson(*person)
	phttp.WriteJsonResponse(w, createdPerson)
}

func (pc personController) filter(w http.ResponseWriter, r *http.Request) {
	person, err := phttp.ParseBody[struct {
		models.Person
		models.BaseFilter
	}](r.Body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	result := pc.filterPeopleUsecase.Filter(person.Person, person.BaseFilter)
	phttp.WriteJsonResponse(w, result)
}

func (pc personController) getPersonById(w http.ResponseWriter, r *http.Request) {
	personId, err := strconv.Atoi(phttp.GetUrlParam(r, "personId"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	person := pc.getPersonByIdUseCase.GetPersonById(personId)
	phttp.WriteJsonResponse(w, person)
}
