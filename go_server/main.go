package main

import (
	"database/sql"
	"go-books/models"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckId(result int64, err error) int64 {
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	// Connect to database
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	CheckId(models.DBAddBook(db, "Bill", "Testing", []string{"Space"}))
	//router.get("/api/v1/books/", controllers.findbooks)
	//router.get("/api/v1/books/:id", controllers.findbook)
	//router.post("/api/v1/books/", controllers.createbook)
	//router.put("/api/v1/books/:id", controllers.updatebook)
	//router.delete("api/v1/books/:id", controllers.deletebook)

	router.Run(":5000")
}
