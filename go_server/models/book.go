package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Book
type Book struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Year     string    `json:"year"`
	Tags     []Tag     `json:"tags" gorm:"ForeignKey:UserID"`
	Read     bool      `json:"read"`
	ReadDate time.Time `json:"readDate"`
}

//Tags
type Tag struct {
	Title  string
	BookID uint `gorm:"column:book_id"`
}

// Tester

type Tester struct {
	gorm.Model
	Name   string
	BookID uint
}
