package controllers

import (
	"net/http"
	"strconv"

	"github.com/megalypse/golang-clean-arch/internal/domain/models"
	"github.com/megalypse/golang-clean-arch/internal/domain/usecases"
	"github.com/megalypse/golang-clean-arch/internal/presentation/phttp"
)

type PersonController struct {
	createPersonUsecase  usecases.CreatePerson
	getPersonByIdUsecase usecases.GetPersonById
	filterPeopleUsecase  usecases.FilterPeople
	updatePersonUsecase  usecases.UpdatePerson
	personExistsUsecase  usecases.PersonExists
	getAllUsecase        usecases.GetAll
	deletePersonUsecase  usecases.DeletePerson
}

func NewPersonController(personService usecases.PersonService) PersonController {
	return PersonController{
		createPersonUsecase:  personService,
		getPersonByIdUsecase: personService,
		filterPeopleUsecase:  personService,
		updatePersonUsecase:  personService,
		personExistsUsecase:  personService,
		getAllUsecase:        personService,
		deletePersonUsecase:  personService,
	}
}

func (pc PersonController) GetHandlers() map[string]phttp.RouteDefinition {
	return map[string]phttp.RouteDefinition{
		"Create person": {
			Method:       http.MethodPost,
			Route:        "/person",
			HandlingFunc: pc.createPerson,
		},
		"Get person by id": {
			Method:       http.MethodGet,
			Route:        "/person/{personId}",
			HandlingFunc: pc.getPersonById,
		},
		"Filter people": {
			Method:       http.MethodPost,
			Route:        "/person/filter",
			HandlingFunc: pc.filter,
		},
		"Update person": {
			Method:       http.MethodPut,
			Route:        "/person",
			HandlingFunc: pc.updatePerson,
		},
		"Get all people": {
			Method:       http.MethodGet,
			Route:        "/person",
			HandlingFunc: pc.getAllPeople,
		},
		"Delete person": {
			Method:       http.MethodDelete,
			Route:        "/person/{personId}",
			HandlingFunc: func(w http.ResponseWriter, r *http.Request) {},
		},
	}
}

func (pc PersonController) deletePerson(w http.ResponseWriter, r *http.Request) {
	personId, err := strconv.ParseInt(phttp.GetUrlParam(r, "personId"), 10, 64)
	if err != nil {
		phttp.WriteError(w, http.StatusUnprocessableEntity)
		return
	}

	exists := pc.personExistsUsecase.Exists(personId)
	if !exists {
		phttp.WriteError(w, http.StatusNotFound)
		return
	}

	pc.deletePersonUsecase.Delete(personId)
	phttp.WriteJsonResponse(w, "Success", http.StatusNoContent)
}

// @Summary Gets all people from database
// @Tags Person
// @Success 200 {array} models.Person
// @Failure 500 {object} phttp.RequestFailed "Internal Server Error"
// @Router /person [get]
func (pc PersonController) getAllPeople(w http.ResponseWriter, r *http.Request) {
	people := pc.getAllUsecase.GetAll()
	phttp.WriteJsonResponse(w, people, http.StatusOK)
}

// @Summary Creates a new person
// @Tags Person
// @Success 201 {object} models.Person
// @Failure 422 {object} phttp.RequestFailed "Unprocessable Entity"
// @Failure 500 {object} phttp.RequestFailed "Internal Server Error"
// @Param request body models.Person true "Create person request"
// @Router /person [post]
func (pc PersonController) createPerson(w http.ResponseWriter, r *http.Request) {
	person, err := phttp.ParseBody[models.Person](r.Body)
	if err != nil {
		phttp.WriteError(w, http.StatusUnprocessableEntity)
		return
	}

	createdPerson := pc.createPersonUsecase.CreatePerson(*person)
	phttp.WriteJsonResponse(w, createdPerson, http.StatusCreated)
}

// @Summary Filter person
// @Tags Person
// @Success 200 {object} models.Paginated[models.Person]
// @Param request body models.BaseFilter[models.Person] true "Filter people request"
// @Router /person/filter [post]
func (pc PersonController) filter(w http.ResponseWriter, r *http.Request) {
	filter, err := phttp.ParseBody[models.BaseFilter[models.Person]](r.Body)

	if err != nil {
		phttp.WriteError(w, http.StatusUnprocessableEntity)
		return
	}

	result := pc.filterPeopleUsecase.Filter(filter)
	phttp.WriteJsonResponse(w, result, http.StatusOK)
}

// @Summary Gets a person by id
// @Tags Person
// @Success 200 {object} models.Person
// @Failure 404 {object} phttp.RequestFailed "Person not found"
// @Failure 500 {object} phttp.RequestFailed "Internal Server Error"
// @Param personId path int true "Person ID"
// @Router /person/{personId} [get]
func (pc PersonController) getPersonById(w http.ResponseWriter, r *http.Request) {
	personId, err := strconv.ParseInt(phttp.GetUrlParam(r, "personId"), 10, 64)

	if err != nil {
		phttp.WriteError(w, http.StatusUnprocessableEntity)
		return
	}

	person := pc.getPersonByIdUsecase.GetPersonById(personId)
	phttp.WriteJsonResponse(w, person, http.StatusOK)
}

// @Summary Updates a person
// @Tags Person
// @Success 200 {object} models.Person
// @Failure 404 {object} phttp.RequestFailed "Person not found"
// @Failure 422 {object} phttp.RequestFailed "Unprocessable Entity"
// @Failure 500 {object} phttp.RequestFailed "Internal Server Error"
// @Param request body models.Person true "Person model containing updated data"
// @Router /person [put]
func (pc PersonController) updatePerson(w http.ResponseWriter, r *http.Request) {
	updatedPerson, err := phttp.ParseBody[models.Person](r.Body)

	if err != nil {
		phttp.WriteError(w, http.StatusUnprocessableEntity)
		return
	}

	exists := pc.personExistsUsecase.Exists(*updatedPerson.Id)
	if !exists {
		phttp.WriteError(w, http.StatusNotFound)
		return
	}

	updatedPerson.DeletedAt = nil
	person := pc.updatePersonUsecase.Update(updatedPerson)
	phttp.WriteJsonResponse(w, person, http.StatusOK)
}
