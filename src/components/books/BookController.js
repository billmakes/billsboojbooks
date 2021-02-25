import LOCard from './BookCard'
import LOTile from './BookTile'

export default {
  name: 'BookController',
  props: {
    source: {
      type: Object,
      required: true
    },
    view: {
      type: String,
      required: true
    }
  },
  computed: {
    Book() {
      return this.view === 'card' ? LOCard : LOTile
    }
  },
  render(h) {
    return h(this.Book, {
      props: {
        source: this.source
      }
    })
  }
}
