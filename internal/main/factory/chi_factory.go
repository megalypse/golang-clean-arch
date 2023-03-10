package factory

import (
	"net/http"

	"github.com/go-chi/chi"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
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
