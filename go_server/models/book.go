package models

import (
	"fmt"
	"time"
)

//Book
type Book struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Year     string    `json:"year"`
	Tags     []string  `json:"tags"`
	Read     bool      `json:"read"`
	ReadDate time.Time `json:"readDate"`
}

// Update the book
func (b *Book) UpdateBook(updatedBook *Book) Book {
	fmt.Println("value of b: ", b)
	fmt.Println("value of updatedBook: ", updatedBook)
	if updatedBook.Author != "" {
		b.Author = updatedBook.Author
	}

	b.Read = updatedBook.Read

	if updatedBook.Tags != nil {
		b.Tags = updatedBook.Tags
	}
	if updatedBook.Title != "" {
		b.Title = updatedBook.Title
	}
	if updatedBook.Year != "" {
		b.Year = updatedBook.Year
	}
	return *b
}
