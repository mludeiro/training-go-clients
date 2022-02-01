package routes

import (
	"github.com/gorilla/mux"
	"github.com/pluralsight/webservice/pkg/controllers"
)

var RegisterClientsRoutes = func(router *mux.Router) {
	router.HandleFunc("/client/", controllers.CreateClient).Methods("POST")
	router.HandleFunc("/client/", controllers.GetClient).Methods("GET")
	router.HandleFunc("/client/{clientId}", controllers.GetClientById).Methods("GET")
	router.HandleFunc("/client/{clientId}", controllers.UpdateClient).Methods("PUT")
	router.HandleFunc("/client/{clientId}", controllers.DeleteClient).Methods("DELETE")
}
