package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pluralsight/webservice/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

/*
	0- the user interacts with de routes
	1- the control is first in the routes
	2- the routes give the control to the controllers
	3- the controllers will give the control to the models
	4- the models hold the operations of the database
*/

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)

	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)

	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)

	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)

	return book
}
