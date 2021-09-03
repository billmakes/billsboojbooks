package controllers

import (
	"fmt"
	"go-books/db"
	"go-books/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all books
func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"books": db.Books})
}

// Get book by ID
func GetBook(c *gin.Context) {
	id := c.Param("id")

	for _, b := range db.Books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
}

// Create a new Book
func AddBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	db.Books = append(db.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// Edit an existing Book
func UpdateBook(c *gin.Context) {
	var updatedBook models.Book
	id := c.Param("id")

	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}
	for i, b := range db.Books {
		if b.ID == id {
			fmt.Println("FOR LOOP B: ", b)
			db.Books[i] = b.UpdateBook(&updatedBook)
			fmt.Println("FOR LOOP B after UPDATE: ", b)
			fmt.Println("the books: ", db.Books)
		}
	}
}

// Delete a Book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, b := range db.Books {
		if b.ID == id {
			db.Books = append(db.Books[:i], db.Books[i+1:]...)
		}
	}
}
