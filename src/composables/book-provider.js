import { reactive, computed, readonly } from '@vue/composition-api'
import BookService from '@/services/'

const bookFields = ['title', 'author', 'year', 'tags', 'read', 'readDate']

const state = reactive({
  books: []
})

function init() {
  BookService.getAllBooks().then(res => {
    state.books = res.books
  })
}

function addBook(book) {
  BookService.addBook(book).then(res => {
    state.books.push(res.data.book)
  })
}

async function getBook(id) {
  const bookInState = state.books.find(book => book.id === id)
  if (state.books.length && bookInState) return await bookInState
  else
    return BookService.getBook(id)
      .then(res => {
        return res.book
      })
      .catch(() => {
        this.$router.push({ name: 'wild' })
        return
      })
}

function editBook(id, params) {
  state.books.forEach(book => {
    if (book.id === id) {
      assignFields(book, params)
      BookService.updateBook(book.id, book)
    }
  })
}

function removeBook(id) {
  state.books = state.books.filter(book => book.id !== id)
  BookService.deleteBook(id)
}

function assignFields(assignObj, assignee) {
  bookFields.forEach(field => {
    assignObj[field] = assignee[field]
  })
}

function filterAsc(field) {
  state.books = state.books.sort(function(a, b) {
    var textA = a[field].toUpperCase()
    var textB = b[field].toUpperCase()
    return textA < textB ? -1 : textA > textB ? 1 : 0
  })
}

function filterDesc(field) {
  state.books = state.books.sort(function(a, b) {
    var textA = a[field].toUpperCase()
    var textB = b[field].toUpperCase()
    return textA > textB ? -1 : textA < textB ? 1 : 0
  })
}

const totalLength = computed(() => state.books.length)
const getBooks = computed(() => state.books)

export const BookStore = readonly({
  init,
  totalLength,
  getBook,
  getBooks,
  addBook,
  editBook,
  filterAsc,
  filterDesc,
  removeBook
})
