package main

import (
	"./routes"
	"log"
	"net/http"
)

func main() {
	// initialize router
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":10001", router))
}