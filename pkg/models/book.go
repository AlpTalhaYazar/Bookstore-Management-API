package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/talha-yazar/Bookstore-Management-API/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	config.Connect()

	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var Books []Book

	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, error) {
	var book Book
	err := db.Where("ID=?", Id).Find(&book).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (b *Book) UpdateBookById() bool {
	isRecordNotFound := db.Model(&Book{}).Where("ID=?", b.Model.ID).Find(&Book{}).RecordNotFound()
	if isRecordNotFound {
		fmt.Println("Book not found")
		return false
	}

	db.Model(&Book{}).Where("ID=?", b.Model.ID).Update(&b)
	return true
}

func DeleteBookById(Id int64) bool {
	isRecordNotFound := db.Model(&Book{}).Where("ID=?", Id).Find(&Book{}).RecordNotFound()
	if isRecordNotFound {
		fmt.Println("Book not found")
		return false
	}

	db.Where("ID=?", Id).Delete(&Book{})
	return true
}
