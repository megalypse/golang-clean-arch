package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
