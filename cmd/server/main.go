package main

import (
	"log"
	"net/http"

	"github.com/megalypse/golang-clean-arch/internal/main/factory"
)

func main() {
	router := factory.GetRouter()
	port := ":3000"

	factory.BootControllers()

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(port, router)
	log.Println(err.Error())
}
