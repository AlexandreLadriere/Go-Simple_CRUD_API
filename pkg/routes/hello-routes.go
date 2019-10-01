package routes

import (
	"github.com/AlexandreLadriere/Go-Simple_CRUD_API/pkg/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	addHelloHandler(r)
	return r
}

func addHelloHandler(router *mux.Router) {
	router.HandleFunc("/", controllers.GetHelloDefault).Methods("GET")
	router.HandleFunc("/hellos", controllers.GetHellos).Methods("GET")
	router.HandleFunc("/hello", controllers.GetHello).Methods("GET")
	router.HandleFunc("/hello", controllers.CreateHello).Methods("POST")
	router.HandleFunc("/hello", controllers.DeleteHello).Methods("DELETE")
}