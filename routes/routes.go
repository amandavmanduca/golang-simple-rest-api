package routes

import (
	controllers "api-rest/controllers"
	"api-rest/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)
	r.HandleFunc("/api/personalities", controllers.Personalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetById).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.DeleteOne).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdateOne).Methods("Post")
	r.HandleFunc("/api/personalities", controllers.Create).Methods("Post")
	log.Fatal(http.ListenAndServe(
		":5000",
		handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
