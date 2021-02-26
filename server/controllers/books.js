const db = require('../db/db')

class BooksController {
  getAllBooks(req, res) {
    return res.status(200).send({
      success: 'true',
      message: 'books retrieved successfully',
      books: db
    })
  }

  getBook(req, res) {
    const id = parseInt(req.params.id, 10)
    db.map(book => {
      if (book.id === id) {
        return res.status(200).send({
          success: 'true',
          message: 'book retrieved successfully',
          book
        })
      }
    })
    return res.status(404).send({
      success: 'false',
      message: 'book does not exist'
    })
  }

  addBook(req, res) {
    if (!req.body.title) {
      return res.status(400).send({
        success: 'false',
        message: 'book title is required'
      })
    } else if (!req.body.author) {
      return res.status(400).send({
        success: 'false',
        message: 'book author is required'
      })
    } else if (!req.body.year) {
      return res.status(400).send({
        success: 'false',
        message: 'book year is required'
      })
    } else if (!req.body.read) {
      return res.status(400).send({
        success: 'false',
        message: 'book read status is required'
      })
    }
    const book = {
      id: db.length + 1,
      title: req.body.title,
      author: req.body.author,
      year: req.body.year,
      tags: req.body.tags,
      read: req.body.read || false,
      readDate: null
    }

    if (book.read) {
      book.readDate = new Date()
    }

    db.push(book)
    return res.status(201).send({
      success: 'true',
      message: 'book added successfully',
      book
    })
  }

  updateBook(req, res) {
    const id = parseInt(req.params.id, 10)
    let bookFound
    let itemIndex
    db.map((book, index) => {
      if (book.id === id) {
        bookFound = book
        itemIndex = index
      }
    })

    if (!bookFound) {
      return res.status(404).send({
        success: 'false',
        message: 'book not found'
      })
    }

    if (!req.body.title) {
      return res.status(400).send({
        success: 'false',
        message: 'book title is required'
      })
    } else if (!req.body.author) {
      return res.status(400).send({
        success: 'false',
        message: 'book author is required'
      })
    } else if (!req.body.year) {
      return res.status(400).send({
        success: 'false',
        message: 'book year is required'
      })
    }

    const updatedBook = {
      id: bookFound.id,
      title: req.body.title || bookFound.title,
      author: req.body.author || bookFound.author,
      year: req.body.year || bookFound.year,
      tags: req.body.tags || bookFound.tags,
      read: req.body.read || bookFound.read || false,
      readDate: req.body.readDate || bookFound.readDate || null
    }

    if (req.body.read && !bookFound.readDate) {
      updatedBook.readDate = new Date()
    }

    db.splice(itemIndex, 1, updatedBook)

    return res.status(201).send({
      success: 'true',
      message: 'book added successfully',
      updatedBook
    })
  }

  deleteBook(req, res) {
    const id = parseInt(req.params.id, 10)

    db.map((book, index) => {
      if (book.id === id) {
        db.splice(index, 1)
        return res.status(200).send({
          success: 'true',
          message: 'book deleted successfuly'
        })
      }
    })

    return res.status(404).send({
      success: 'false',
      message: 'book not found'
    })
  }
}

const bookController = new BooksController()
module.exports = bookController
