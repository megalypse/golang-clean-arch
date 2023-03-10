package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/megalypse/golang-clean-arch/internal/main/factory"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	router := factory.GetRouter()
	port := os.Getenv("ARCHER_PORT")

	factory.BootControllers()

	// server := http.Server{
	// 	Addr:    port,
	// 	Handler: router,
	// }
	// log.Fatal(server.ListenAndServeTLS("certfile.crt", "server.key"))
	log.Printf("Listening on port %s", port)
	err = http.ListenAndServe(port, router)
	log.Println(err.Error())
}
