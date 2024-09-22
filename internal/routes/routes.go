package routes

import (
	"net/http"

	"github.com/TalesPalma/api-go-rest/controllers"
	"github.com/TalesPalma/api-go-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	gorila := mux.NewRouter()
	gorila.Use(middleware.ContentTypeMiddlware)
	gorila.HandleFunc("/", controllers.Home)
	gorila.HandleFunc("/personalidade", controllers.Personalidades).Methods("Get")
	gorila.HandleFunc("/personalidade/{id}", controllers.ReturnById).Methods("Get")
	gorila.HandleFunc("/personalidade", controllers.PersonalidadePost).Methods("Post")
	gorila.HandleFunc("/personalidade/{id}", controllers.PersonalidadePut).Methods("Put")
	gorila.HandleFunc("/personalidade/{id}", controllers.PersonalidadeDelete).Methods("Delete")
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(gorila))
}
