package factory

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/megalypse/golang-clean-arch/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()

	rawPort := os.Getenv("SERVER_HOST_PORT")
	swaggerUrl := fmt.Sprintf("http://localhost:%s/swagger/doc.json", rawPort)

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerUrl),
	))
}

func GetRouter() CustomHttpHandler {
	return router
}

type CustomHttpHandler interface {
	http.Handler

	Get(string, http.HandlerFunc)
	Post(string, http.HandlerFunc)
	Put(string, http.HandlerFunc)
	Patch(string, http.HandlerFunc)
	Delete(string, http.HandlerFunc)
}
