package main

import (
	"log"
	"net/http"
	"os"

	"github.com/megalypse/golang-clean-arch/internal/main/factory"
)

func main() {
	router := factory.GetRouter()
	rawPort := os.Getenv("SERVER_CONTAINER_PORT")
	port := ":" + rawPort

	factory.BootControllers()

	// server := http.Server{
	// 	Addr:    port,
	// 	Handler: router,
	// }
	// log.Fatal(server.ListenAndServeTLS("certfile.crt", "server.key"))
	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(port, router)
	log.Println(err.Error())
}
