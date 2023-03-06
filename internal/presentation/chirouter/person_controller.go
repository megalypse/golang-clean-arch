package chirouter

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/megalypse/golang-clean-arch/internal/domain/models"
	"github.com/megalypse/golang-clean-arch/internal/domain/usecases"
)

type PersonController struct {
	router               *chi.Mux
	createPersonUsecase  usecases.CreatePerson
	getPersonByIdUsecase usecases.GetPersonById
}

func NewPersonController(
	router *chi.Mux,
	createPersonUsecase usecases.CreatePerson,
	getPersonByIdUsecase usecases.GetPersonById,
) PersonController {
	return PersonController{
		router:               router,
		createPersonUsecase:  createPersonUsecase,
		getPersonByIdUsecase: getPersonByIdUsecase,
	}
}

func (pc PersonController) BootController() {
	pc.router.Route("/person", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			log.Println("Creating new person...")

			person := models.Person{}
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			createdPerson := pc.createPersonUsecase.CreatePerson(person)
			writeJsonResponse(w, createdPerson)
		})

		r.Get("/{personId}", func(w http.ResponseWriter, r *http.Request) {
			personId, err := strconv.Atoi(chi.URLParam(r, "personId"))

			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
				return
			}

			person := pc.getPersonByIdUsecase.GetPersonById(personId)
			writeJsonResponse(w, person)
		})
	})
}
