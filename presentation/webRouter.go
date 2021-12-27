package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebRouter struct {
	router           *mux.Router
	ClientController ClientController
}

func (this *WebRouter) GetRouter() *mux.Router {
	if this.router != nil {
		return this.router
	}

	sm := mux.NewRouter()

	sm.Methods(http.MethodGet).Path("/api/v1/clients/{id:[0-9]+}").HandlerFunc(this.ClientController.GetClient)
	sm.Methods(http.MethodGet).Path("/api/v1/clients").HandlerFunc(this.ClientController.GetClients)

	sm.Methods(http.MethodDelete).Path("/api/v1/clients/{id:[0-9]+}").HandlerFunc(this.ClientController.DeleteClient)
	sm.Methods(http.MethodPost).Path("/api/v1/clients").HandlerFunc(this.ClientController.PostClient)
	sm.Methods(http.MethodPut).Path("/api/v1/clients").HandlerFunc(this.ClientController.PostClient)

	// Just to log the calls, response code and time spent
	sm.Use(LogMiddleWare)

	this.router = sm
	return this.router
}
