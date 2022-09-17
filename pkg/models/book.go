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
