package models

import (
	"github.com/cunnellp5/pt3_book_store/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

func (b *Book) UpdateBook(id int64) *Book {
	panic("update unimplemented")
}

func (b *Book) DeleteBook(id int64) string {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return "Deleted, fool"
}
