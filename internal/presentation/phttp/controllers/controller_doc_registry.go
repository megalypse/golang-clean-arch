package controllers

import (
	"github.com/go-chi/chi"
)

func Main() {
	r := chi.NewRouter()
	personController := PersonController{}

	r.Get("/person", personController.createPerson)
}
