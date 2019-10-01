package controllers

import (
	"encoding/json"
	"github.com/AlexandreLadriere/Go-Simple_CRUD_API/pkg/models"
	"github.com/AlexandreLadriere/Go-Simple_CRUD_API/pkg/utils"
	"net/http"
)

// Index
func GetHellos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := models.GetHellos()
	_ = json.NewEncoder(w).Encode(res)
}

// display default Hello
func GetHelloDefault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Hello"
	_ = json.NewEncoder(w).Encode(res)
}

// get hello according to the language in the query
func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := models.GetHello(utils.GetQueryParam(r, "lang"))
	_ = json.NewEncoder(w).Encode(res)
}

// create a new hello
func CreateHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var newHello models.Hello
	_ = json.NewDecoder(r.Body).Decode(&newHello)
	res := models.CreateNewHello(newHello)
	_ = json.NewEncoder(w).Encode(res)
}

// delete a Hello
func DeleteHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := models.DeleteHello(utils.GetQueryParam(r, "lang"))
	_ = json.NewEncoder(w).Encode(res)
}
