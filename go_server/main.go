package main

import (
	"go-books/controllers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

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

	router.GET("/api/v1/books/", controllers.GetBooks)
	router.GET("/api/v1/books/:id/", controllers.GetBook)
	router.POST("/api/v1/books/", controllers.AddBook)
	router.PUT("/api/v1/books/:id", controllers.UpdateBook)
	router.DELETE("api/v1/books/:id", controllers.DeleteBook)

	router.Run(":5000")
}
