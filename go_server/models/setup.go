package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database

	DB.DropTableIfExists(&Book{}, &Tag{})

	DB.AutoMigrate(&Book{}, &Tag{})

	for _, b := range Books {
		DB.Create(&b)
	}
}
