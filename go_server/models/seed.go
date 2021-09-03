package models

import (
	"time"
)

var Books = []Book{
	{
		Title:    "Pride and Prejudice",
		Author:   "Jane Austen",
		Year:     "1813",
		Read:     false,
		Tags:     []Tag{{Title: "test"}, {Title: "Another"}},
		ReadDate: time.Now(),
	},
	{
		Title:    "To Kill a Mockingbird",
		Author:   "Harper Lee",
		Year:     "1960",
		Read:     false,
		Tags:     []Tag{{Title: "killing"}, {Title: "harper lee"}},
		ReadDate: time.Now(),
	},
	{
		Title:    "The Great Gatsby",
		Author:   "F. Scott Fitzgerald",
		Year:     "1925",
		Read:     false,
		Tags:     nil,
		ReadDate: time.Now(),
	},
	{
		Title:    "One Hundred Years of Solitude",
		Author:   "Gabriel García Márquez",
		Year:     "1967",
		Read:     false,
		Tags:     nil,
		ReadDate: time.Now(),
	},
	{
		Title:    "In Cold Blood",
		Author:   "Truman Capote",
		Year:     "1965",
		Read:     false,
		Tags:     nil,
		ReadDate: time.Now(),
	},
	{
		Title:    "The C Programming Language. 2nd Edition",
		Author:   "Brian Kernighan and Dennis Ritchie",
		Year:     "1978",
		Read:     false,
		Tags:     nil,
		ReadDate: time.Now(),
	},
	{
		Title:    "Fantastic Mr. Fox",
		Author:   "Roald Dahl",
		Year:     "1970",
		Read:     false,
		Tags:     nil,
		ReadDate: time.Now(),
	},
	{
		Title:    "Z",
		Author:   "Bill",
		Year:     "8888",
		Read:     false,
		Tags:     nil,
		ReadDate: time.Now(),
	},
	{
		Title:    "B",
		Author:   "Also Bill",
		Year:     "3333",
		Read:     false,
		Tags:     nil,
		ReadDate: time.Now(),
	},
}
