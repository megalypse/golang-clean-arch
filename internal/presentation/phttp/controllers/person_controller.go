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
// @Tags Person
// @Success 200 {object} models.Person
// @Param request body models.Person true "Create person request"
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
// @Tags Person
// @Success 200 {object} models.Paginated[models.Person]
// @Param request body models.BaseFilter[models.Person] true "Filter people request"
// @Router /person/filter [post]
func (pc PersonController) filter(w http.ResponseWriter, r *http.Request) {
	filter, err := phttp.ParseBody[models.BaseFilter[models.Person]](r.Body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	result := pc.filterPeopleUsecase.Filter(filter)
	phttp.WriteJsonResponse(w, result)
}

// @Summary Gets a person by id
// @Tags Person
// @Success 200 {object} models.Person
// @Param personId path int true "Person ID"
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
