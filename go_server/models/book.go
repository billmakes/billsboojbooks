package models

import (
	"database/sql"
	"time"
)

//Book
type Book struct {
	BookID   uint      `json:"id"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Year     string    `json:"year"`
	Read     bool      `json:"read"`
	ReadDate time.Time `json:"readDate"`
}

//Tags
type Tag struct {
	ID    uint
	Title string
}

func DBAddBook(db *sql.DB, title string, author string, tags []string) (int64, error) {
	r, err := db.Exec("insert into Book(title, author) values(?, ?)",
		title, author)
	if err != nil {
		return 0, err
	}
	bookID, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	// For each tag requested for this book, find the tag's ID in the DB and
	// create an entry in the BookTag linking table.
	for _, tag := range tags {
		row := db.QueryRow("select tagID from Tag where name = ?", tag)
		var tagID int
		err := row.Scan(&tagID)
		if err != nil {
			return 0, err
		}

		_, err = db.Exec("insert into BookTag(bookID, tagID) values(?, ?)", bookID, tagID)
		if err != nil {
			return 0, err
		}
	}

	return bookID, nil
}
