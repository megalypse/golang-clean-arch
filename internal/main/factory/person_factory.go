package factory

import (
	"github.com/megalypse/golang-clean-arch/internal/data/repository"
	"github.com/megalypse/golang-clean-arch/internal/data/service"
	"github.com/megalypse/golang-clean-arch/internal/infra/repositoryimpl"
	"github.com/megalypse/golang-clean-arch/internal/presentation/chirouter"
)

var personRepository repository.PersonRepository
var personService service.PersonService

func init() {
	personRepository = repositoryimpl.PgPersonRepository{}
	personService = service.PersonService{
		PersonRepository: personRepository,
	}
}

func makePersonController() *chirouter.PersonController {
	controller := chirouter.NewPersonController(GetRouter(), personService, personService)

	return &controller
}

func BootControllers() {
	makePersonController().BootController()
}
