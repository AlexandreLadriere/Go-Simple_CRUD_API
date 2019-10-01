package main

import (
	"github.com/AlexandreLadriere/Go-Simple_CRUD_API/pkg/routes"
	"log"
	"net/http"
)

func main() {
	// initialize router
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":10001", router))
}