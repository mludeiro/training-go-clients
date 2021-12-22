package repository

import (
	"training-go-clients/database"
	"training-go-clients/entity"
)

type IClientRepository interface {
	Get(id uint, fetchs []string) (*entity.Client, error)
	GetAll(query entity.Query) (entity.ClientResultSet, error)
	Add(a *entity.Client) (*entity.Client, error)
	Delete(id uint) (*entity.Client, error)
}

type Client struct {
	DataBase *database.Database
}

func (this *Client) Get(id uint, fetchs []string) (*entity.Client, error) {
	client := entity.Client{}
	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	query := db.Find(&client, id)
	if query.Error == nil && query.RowsAffected == 1 {
		return &client, nil
	} else {
		return nil, query.Error
	}
}

func (this *Client) GetAll(query entity.Query) (entity.ClientResultSet, error) {
	clients := entity.ClientResultSet{Query: query}
	err := this.DataBase.GetQueryDB(query).GetResult(&clients.ResultSet, &clients.Data)

	return clients, err
}

func (this *Client) Add(a *entity.Client) (*entity.Client, error) {
	query := this.DataBase.GetDB().Create(a)
	if query.Error != nil {
		return nil, query.Error
	}
	return a, nil
}

func (this *Client) Delete(id uint) (*entity.Client, error) {
	client := entity.Client{}
	query := this.DataBase.GetDB().Delete(&client, id)
	if query.Error != nil {
		return nil, query.Error
	} else {
		if query.RowsAffected == 1 {
			return &client, nil
		} else {
			return nil, nil
		}
	}
}
