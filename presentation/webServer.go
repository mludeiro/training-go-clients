package presentation

import (
	"net/http"
	"time"
	"training-go-clients/tools"
)

type WebServer struct {
	Router WebRouter
	server http.Server
}

func (this *WebServer) CreateServer() {
	this.server = http.Server{
		Addr:         ":5000",                 // configure the bind address
		Handler:      this.Router.GetRouter(), // set the default handler
		ErrorLog:     tools.GetLogger(),       // set the logger for the server
		ReadTimeout:  5 * time.Second,         // max time to read request from the client
		WriteTimeout: 10 * time.Second,        // max time to write response to the client
		IdleTimeout:  120 * time.Second,       // max time for connections using TCP Keep-Alive
	}

	tools.GetLogger().Println("Starting server on port 5000")

	err := this.server.ListenAndServe()
	if err != nil {
		tools.GetLogger().Printf("Error starting server: %s\n", err)
	}
}
