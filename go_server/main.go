package main

import (
	"go-books/controllers"
	"go-books/models"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET,PATCH, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	// Connect to database
	models.ConnectDatabase()

	router.GET("/api/v1/books/", controllers.FindBooks)
	router.GET("/api/v1/books/:id", controllers.FindBook)
	router.POST("/api/v1/books/", controllers.CreateBook)
	router.PUT("/api/v1/books/:id", controllers.UpdateBook)
	router.DELETE("api/v1/books/:id", controllers.DeleteBook)

	router.Run(":5000")
}
