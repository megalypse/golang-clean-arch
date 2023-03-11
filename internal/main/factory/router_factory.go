package factory

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/megalypse/golang-clean-arch/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func BootControllers() {
	router := GetRouter()

	controllers := GetControllers()

	for _, controller := range controllers {
		for _, routeDefinition := range controller.GetHandlers() {
			handlingFunc := routeDefinition.HandlingFunc
			route := routeDefinition.Route

			switch routeDefinition.Method {
			case http.MethodGet:
				router.Get(route, handlingFunc)
			case http.MethodPost:
				router.Post(route, handlingFunc)
			case http.MethodPut:
				router.Put(route, handlingFunc)
			case http.MethodPatch:
				router.Patch(route, handlingFunc)
			case http.MethodDelete:
				router.Delete(route, handlingFunc)
			default:
				log.Fatalf("Http method not supported: %q", routeDefinition.Method)
			}
		}
	}

	bootSwagger(router)
}

func bootSwagger(router CustomHttpHandler) {
	rawPort := os.Getenv("SERVER_HOST_PORT")
	swaggerUrl := fmt.Sprintf("http://localhost:%s/swagger/doc.json", rawPort)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerUrl),
	))
}
