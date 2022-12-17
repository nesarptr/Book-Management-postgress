package models

import (
	"github.com/nesarptr/Book-Management-postgress/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (book *Book) CreateBook() *Book {
	db.Create(book)
	return book
}

func GetAllBooks() *[]Book {
	var Books []Book
	db.Find(&Books)
	return &Books
}

func GetBookById(id uint) (*Book, *gorm.DB) {
	var Book Book
	db := db.First(&Book, id)
	return &Book, db
}

func DeleteById(id uint) *Book {
	var Book Book
	db.Delete(&Book, id)
	Book.ID = id
	return &Book
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
