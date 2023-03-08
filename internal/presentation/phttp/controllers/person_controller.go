package controllers

import (
	"net/http"
	"strconv"

	"github.com/megalypse/golang-clean-arch/internal/data/service"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
	"github.com/megalypse/golang-clean-arch/internal/presentation/phttp"
)

type personController struct {
	personService service.PersonService
}

func NewPersonController(personService service.PersonService) personController {
	return personController{
		personService: personService,
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

	createdPerson := pc.personService.CreatePerson(*person)
	phttp.WriteJsonResponse(w, createdPerson)
}

func (pc personController) filter(w http.ResponseWriter, r *http.Request) {
	person, err := phttp.ParseBody[models.Person](r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	result := pc.personService.Filter(*person)
	phttp.WriteJsonResponse(w, result)
}

func (pc personController) getPersonById(w http.ResponseWriter, r *http.Request) {
	personId, err := strconv.Atoi(phttp.GetUrlParam(r, "personId"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	person := pc.personService.GetPersonById(personId)
	phttp.WriteJsonResponse(w, person)
}
