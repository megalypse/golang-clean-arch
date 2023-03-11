package controllers

import (
	"net/http"
	"strconv"

	"github.com/megalypse/golang-clean-arch/internal/domain/models"
	"github.com/megalypse/golang-clean-arch/internal/domain/usecases"
	"github.com/megalypse/golang-clean-arch/internal/presentation/phttp"
)

type PersonController struct {
	createPersonUseCase  usecases.CreatePerson
	getPersonByIdUseCase usecases.GetPersonById
	filterPeopleUsecase  usecases.FilterPeople
}

func NewPersonController(personService usecases.PersonService) PersonController {
	return PersonController{
		createPersonUseCase:  personService,
		getPersonByIdUseCase: personService,
		filterPeopleUsecase:  personService,
	}
}

func (pc PersonController) GetHandlers() map[string]phttp.RouteDefinition {
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

// @Summary Creates a new person
// @Success 200 {object} models.Person
// @Router /person [post]
func (pc PersonController) createPerson(w http.ResponseWriter, r *http.Request) {
	person, err := phttp.ParseBody[models.Person](r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdPerson := pc.createPersonUseCase.CreatePerson(*person)
	phttp.WriteJsonResponse(w, createdPerson)
}

// @Summary Filter person
// @Success 200 {object} models.Paginated[models.Person]
// @Router /person/filter [post]
func (pc PersonController) filter(w http.ResponseWriter, r *http.Request) {
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

// @Summary Gets a person by id
// @Success 200 {object} models.Person
// @Param id path int true "Person ID"
// @Router /person/{personId} [get]
func (pc PersonController) getPersonById(w http.ResponseWriter, r *http.Request) {
	personId, err := strconv.Atoi(phttp.GetUrlParam(r, "personId"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	person := pc.getPersonByIdUseCase.GetPersonById(personId)
	phttp.WriteJsonResponse(w, person)
}
