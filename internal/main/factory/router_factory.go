package factory

import (
	"fmt"
	"log"
	"os"

	"github.com/megalypse/golang-clean-arch/internal/presentation/phttp"
	httpSwagger "github.com/swaggo/http-swagger"
)

func BootControllers() {
	router := GetRouter()

	controllers := GetControllers()

	for _, controller := range controllers {
		for route, routeData := range controller.GetHandlers() {
			handlingFunc := routeData.HandlingFunc

			switch routeData.Method {
			case phttp.GET:
				router.Get(route, handlingFunc)
			case phttp.POST:
				router.Post(route, handlingFunc)
			case phttp.PUT:
				router.Put(route, handlingFunc)
			case phttp.PATCH:
				router.Patch(route, handlingFunc)
			case phttp.DELETE:
				router.Delete(route, handlingFunc)
			default:
				log.Fatalf("Http method not supported: %q", routeData.Method)
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
