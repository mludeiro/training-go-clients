package routes

import (
	"github.com/gorilla/mux"
	"github.com/pluralsight/webservice/pkg/controllers"
)

var RegisterClientsRoutes = func(router *mux.Router) {
	router.HandleFunc("/clients/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/clients/{clientId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/clients/{clientId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/clients/{clientId}", controllers.DeleteBook).Methods("DELETE")
}
