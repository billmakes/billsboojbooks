<template>
  <div>
    <h1>My Books</h1>
    <div class="d-flex">
      <b-form-group class="mr-3 mb-0">
        <b-form-radio-group
          id="selectedView-radio1"
          v-model="selectedView"
          buttons
          name="radio-btn-stacked"
          class="border rounded bg-white"
        >
          <b-form-radio
            v-for="option in viewOptions"
            :key="option.value"
            :value="option"
            :button-variant="viewSelectVariant(option)"
            class="px-1"
            ><b-icon :icon="option.icon" class="mx-1"
          /></b-form-radio>
        </b-form-radio-group>
      </b-form-group>
      <div class="d-flex align-items-center">
        <div v-if="selectedFilter" class="border rounded bg-white p-2">
          <span class="font-weight-bold"> Current filter: </span>
          <span>
            {{ selectedFilter.label }} -
            {{ selectedFilter.value === 'asc' ? 'Ascending' : 'Descending' }}
          </span>
        </div>
        <b-button
          v-for="filter in filters"
          class="ml-2"
          :key="filter.field"
          @click="changeFilter(filter)"
          :variant="filter.value ? 'primary' : ''"
        >
          <span>{{ filter.label }}</span>
          <span v-if="filter.field === 'year' && filter.value">{{
            filter.value === 'asc' ? ': Oldest' : ': Newest'
          }}</span>

          <b-icon
            v-if="filter.value"
            class="ml-1"
            :icon="filter.value === 'asc' ? filter.iconAsc : filter.iconDesc"
          />
        </b-button>
        <b-button
          class="ml-2"
          @click="hideRead = !hideRead"
          :variant="hideRead ? 'warning' : ''"
          ><span>{{ hideRead ? 'Show Read' : 'Hide Read' }}</span>
          <b-icon class="ml-1" :icon="hideRead ? 'eye' : 'eye-slash'"
        /></b-button>
        <b-button class="ml-2" variant="danger" @click="clearFilters"
          >Clear Filters</b-button
        >
      </div>
    </div>
    <div v-if="books" class="d-flex flex-wrap">
      <BookController
        v-for="book in books"
        :source="book"
        :view="selectedView.value"
        :key="book.id"
        v-show="!hideRead || (hideRead && !book.read)"
      />
    </div>
  </div>
</template>
<script>
import BookService from '@/services'
import BookController from '@/components/books/BookController'

const viewOptions = [
  { value: 'card', icon: 'grid-3x3-gap-fill' },
  { value: 'tile', icon: 'list-ul' }
]

const filters = [
  {
    field: 'title',
    label: 'Title',
    value: null,
    iconAsc: 'sort-alpha-down',
    iconDesc: 'sort-alpha-up'
  },
  {
    field: 'author',
    label: 'Author',
    value: null,
    iconAsc: 'sort-alpha-down',
    iconDesc: 'sort-alpha-up'
  },
  {
    field: 'year',
    label: 'Year Published',
    value: null,
    iconAsc: 'sort-numeric-down',
    iconDesc: 'sort-numeric-up'
  }
]
export default {
  name: 'MyBooks',
  components: {
    BookController
  },
  data() {
    return {
      books: null,
      selectedView: viewOptions[0],
      viewOptions,
      filters,
      hideRead: false
    }
  },
  created() {
    this.getBooks()
  },
  computed: {
    selectedFilter() {
      return this.filters.find(filter => filter.value)
    }
  },
  methods: {
    getBooks() {
      BookService.getAllBooks().then(res => {
        this.books = res.books
      })
    },
    clearFilters() {
      this.filters.forEach(f => {
        f.value = null
      })
    },
    changeFilter(filter) {
      this.filters.forEach(f => {
        if (f.field !== filter.field) f.value = null
      })
      if (!filter.value) {
        filter.value = 'asc'
        this.filterAsc(filter.field)
        return
      }
      if (filter.value === 'asc') {
        filter.value = 'desc'
        this.filterDesc(filter.field)
      } else {
        filter.value = 'asc'
        this.filterAsc(filter.field)
      }
    },
    filterAsc(field) {
      this.books = this.books.sort(function (a, b) {
        var textA = a[field].toUpperCase()
        var textB = b[field].toUpperCase()
        return textA < textB ? -1 : textA > textB ? 1 : 0
      })
    },
    filterDesc(field) {
      this.books = this.books.sort(function (a, b) {
        var textA = a[field].toUpperCase()
        var textB = b[field].toUpperCase()
        return textA > textB ? -1 : textA < textB ? 1 : 0
      })
    },
    viewSelectVariant(option) {
      return this.selectedView === option ? 'primary' : 'none'
    }
  }
}
</script>