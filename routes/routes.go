package routes

import (
	"net/http"

	"github.com/TalesPalma/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	gorila := mux.NewRouter()
	gorila.HandleFunc("/", controllers.Home)
	gorila.HandleFunc("/personalidade", controllers.Personalidades).Methods("Get")
	gorila.HandleFunc("/personalidade", controllers.PersonalidadePost).Methods("Post")
	gorila.HandleFunc("/personalidade/{id}", controllers.ReturnById).Methods("Get")
	gorila.HandleFunc("/personalidade/{id}", controllers.PersonalidadeDelete).Methods("Delete")
	gorila.HandleFunc("/personalidade/{id}", controllers.PersonalidadePut).Methods("Put")
	http.ListenAndServe(":8080", gorila)
}
