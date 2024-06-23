package models

import (
	"github.com/Hey-Mynk/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

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

func GetBookById(bookId int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", bookId).Find(&getBook)
	return &getBook, db
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}

func DeleteBook(bookId int) Book {
	var book Book
	db.Where("ID = ?", bookId).Delete(book)
	return book
}
