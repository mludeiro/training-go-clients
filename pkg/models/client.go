package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pluralsight/webservice/pkg/config"
)

var db2 *gorm.DB

type Client struct {
	gorm.Model
	ID      uint   `gorm:"primarykey"`
	Name    string `gorm:"not null"`
	HasDebt bool   `gorm:"-"`
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

func GetClientById(Id int64) (*Client, *gorm.DB) {
	var client Client
	db := db.Where("ID=?", Id).Find(&client)

	return &client, db
}

func DeleteClient(Id int64) Client {
	var client Client
	db.Where("ID=?", Id).Delete(client)
	return client
}
