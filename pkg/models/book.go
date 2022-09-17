package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/talha-yazar/Bookstore-Management-API/pkg/config"
)

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}
