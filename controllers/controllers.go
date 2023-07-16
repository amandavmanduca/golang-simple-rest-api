package controllers

import (
	"api-rest/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Personalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.GetAll())
}

func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	json.NewEncoder(w).Encode(models.GetById(id))
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	json.NewEncoder(w).Encode(models.DeleteOne(id))
}

func Create(w http.ResponseWriter, r *http.Request) {
	var p models.Personality
	json.NewDecoder(r.Body).Decode(&p)
	json.NewEncoder(w).Encode(models.Create(p))
}

func UpdateOne(w http.ResponseWriter, r *http.Request) {
	var p models.Personality
	vars := mux.Vars(r)
	id := vars["id"]
	json.NewDecoder(r.Body).Decode(&p)
	json.NewEncoder(w).Encode(models.UpdateOne(id, p))
}
