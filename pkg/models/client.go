package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pluralsight/webservice/pkg/config"
)

var db2 *gorm.DB

type Client struct {
	gorm.Model
	Name    string `gorm:""json:"name"`
	HasDebt bool   `json:"hasDebt"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Client{})
}

func (c *Client) CreateClient() *Client {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllClients() []Client {
	var clients []Client
	db.Find(&clients)

	return clients
}

func GetClientById(Id int64) *Client {
	var client Client
	db.Where("ID=?", Id).Find(&client)

	return &client
}

func (c *Client) UpdateClient(clientDetails *Client) Client {
	if c.Name != "" {
		clientDetails.Name = c.Name
	}

	clientDetails.HasDebt = c.HasDebt

	db.Save(&clientDetails)
	return *clientDetails
}

func DeleteClient(Id int64) Client {
	var client Client
	db.Where("ID=?", Id).Delete(client)
	return client
}
