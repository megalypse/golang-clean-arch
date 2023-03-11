package factory

import (
	"github.com/megalypse/golang-clean-arch/internal/data/repository"
	"github.com/megalypse/golang-clean-arch/internal/data/service"
	"github.com/megalypse/golang-clean-arch/internal/infra/pgrepository"
	"github.com/megalypse/golang-clean-arch/internal/presentation/phttp/controllers"
)

var personRepository repository.PersonRepository
var personService service.PersonService
var personController controllers.PersonController

func init() {
	personRepository = pgrepository.PgPersonRepository{}
	personService = service.NewPersonService(personRepository)
	personController = controllers.NewPersonController(personService)
}

func GetControllers() []controllers.Controller {
	return []controllers.Controller{
		personController,
	}
}

func GetPersonController() controllers.PersonController {
	return personController
}
