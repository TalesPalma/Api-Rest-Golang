package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TalesPalma/api-go-rest/models"
	"github.com/TalesPalma/api-go-rest/services"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(services.Home())
}
func ReturnById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convertId, _ := strconv.Atoi(id)
	personalidade := services.GetById(convertId)
	if personalidade.Id == 0 {
		w.Write([]byte("Personalidade não encontrada"))
		return
	}

	json.NewEncoder(w).Encode(personalidade)
}

func PersonalidadePost(w http.ResponseWriter, r *http.Request) {

	var novaPersonalidade models.Personalidade
	err := json.NewDecoder(r.Body).Decode(&novaPersonalidade) // pega o body da requisição e decodifica em struct (decodificar)

	if err != nil {
		http.Error(w, "Body is required", http.StatusBadRequest)
		return
	}

	services.Insert(novaPersonalidade)           // Inserido no database
	json.NewEncoder(w).Encode(novaPersonalidade) // retorna o body da requisição (codificar)

}

func PersonalidadeDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convertId, _ := strconv.Atoi(id)
	services.DeleteById(convertId)
}

func PersonalidadePut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convertId, _ := strconv.Atoi(id)

	var newPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&newPersonalidade)
	services.PutByid(convertId, newPersonalidade)
	json.NewEncoder(w).Encode(newPersonalidade)
}
