package factory

import "github.com/go-chi/chi"

var router *chi.Mux

func init() {
	router = chi.NewRouter()
}

func GetRouter() *chi.Mux {
	return router
}
