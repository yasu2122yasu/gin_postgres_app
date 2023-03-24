package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string
	Body  string
}

func GetAll() (books []Book) {
	result := Db.Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}
