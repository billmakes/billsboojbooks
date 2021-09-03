package db

import (
	"go-books/models"
	"time"
)

var Books = []models.Book{
	{
		ID:       "1",
		Title:    "Pride and Prejudice",
		Author:   "Jane Austen",
		Year:     "1813",
		Tags:     []string{"fiction"},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "2",
		Title:    "To Kill a Mockingbird",
		Author:   "Harper Lee",
		Year:     "1960",
		Tags:     []string{"fiction", "classic"},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "3",
		Title:    "The Great Gatsby",
		Author:   "F. Scott Fitzgerald",
		Year:     "1925",
		Tags:     []string{"gatsby", "fiction", "classic"},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "4",
		Title:    "One Hundred Years of Solitude",
		Author:   "Gabriel García Márquez",
		Year:     "1967",
		Tags:     []string{"lonely", "classic"},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "5",
		Title:    "In Cold Blood",
		Author:   "Truman Capote",
		Year:     "1965",
		Tags:     []string{"fiction", "thriller"},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "6",
		Title:    "The C Programming Language. 2nd Edition",
		Author:   "Brian Kernighan and Dennis Ritchie",
		Year:     "1978",
		Tags:     []string{"programming", "technical", "C"},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "7",
		Title:    "Fantastic Mr. Fox",
		Author:   "Roald Dahl",
		Year:     "1970",
		Tags:     []string{"animals", "hunger"},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "8",
		Title:    "Z",
		Author:   "Bill",
		Year:     "8888",
		Tags:     []string{},
		Read:     false,
		ReadDate: time.Now(),
	},
	{
		ID:       "9",
		Title:    "B",
		Author:   "Also Bill",
		Year:     "3333",
		Tags:     []string{},
		Read:     false,
		ReadDate: time.Now(),
	},
}
