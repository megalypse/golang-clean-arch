package factory

import (
	"github.com/megalypse/golang-clean-arch/internal/data/repository"
	"github.com/megalypse/golang-clean-arch/internal/data/service"
	"github.com/megalypse/golang-clean-arch/internal/infra/repositoryimpl"
	"github.com/megalypse/golang-clean-arch/internal/presentation/phttp/controllers"
)

var personRepository repository.PersonRepository
var personService service.PersonService

func init() {
	personRepository = repositoryimpl.PgPersonRepository{}
	personService = service.NewPersonService(personRepository)
}

func GetControllers() []controllers.Controller {
	return []controllers.Controller{
		controllers.NewPersonController(personService),
	}
}
