package container

import (
	"training-go-clients/database"
	"training-go-clients/presentation"
	"training-go-clients/repository"
	"training-go-clients/service"
)

type Container struct {
	WebServer presentation.WebServer
	DataBase  *database.Database
}

func NewContainer() Container {
	database := database.Database{}

	return Container{
		DataBase: &database,

		WebServer: presentation.WebServer{
			Router: presentation.WebRouter{
				ClientController: presentation.ClientController{
					Service: &service.Client{
						IClientRepository: &repository.Client{
							DataBase: &database},
					},
				},
			},
		},
	}
}
