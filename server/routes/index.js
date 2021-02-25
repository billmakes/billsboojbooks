const express = require('express')

const BookController = require('../controllers/books')

const router = express.Router()

router.get('/api/v1/books', BookController.getAllBooks)
router.get('/api/v1/books/:id', BookController.getBook)
router.post('/api/v1/books', BookController.addBook)
router.put('/api/v1/books/:id', BookController.updateBook)
router.delete('/api/v1/books/:id', BookController.deleteBook)

module.exports = router
