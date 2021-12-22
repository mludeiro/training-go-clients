package service

import (
	"training-go-clients/entity"
	"training-go-clients/repository"
)

type IClientService interface {
	Get(uint, []string) (*entity.Client, error)
	GetAll(query entity.Query) (entity.ClientResultSet, error)
	Add(*entity.Client) (*entity.Client, error)
	Delete(uint) (*entity.Client, error)
}

type Client struct {
	repository.IClientRepository
}

// if you want to, you can wrap or redefine the repository method
func (a *Client) GetAll(query entity.Query) (entity.ClientResultSet, error) {
	return a.IClientRepository.GetAll(query)
}
