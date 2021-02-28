<template>
  <div>
    <div v-if="books">
      <b-alert variant="primary" :show="!books.length"
        >You have no books added to your list. Click the Add Book button to
        begin adding entries.</b-alert
      >
      <b-alert variant="danger" :show="books.length === 1">
        <b-icon icon="chat-right-quote" />
        Beware of the person of one book. -- Thomas Aquinas</b-alert
      >
    </div>
    <div class="d-flex justify-content-between">
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
      <div>
        <b-button class="mr-2" variant="primary" @click="addBook"
          >Add Book</b-button
        >
        <b-button v-b-toggle.filter-bar>View Filters</b-button>
      </div>
    </div>
    <div>
      <b-sidebar id="filter-bar" title="Filters" width="30em" shadow backdrop>
        <div class="px-3 py-2">
          <div class="FilterBar">
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
                :icon="
                  filter.value === 'asc' ? filter.iconAsc : filter.iconDesc
                "
              />
            </b-button>
            <b-button
              class="ml-2"
              @click="hideRead = !hideRead"
              :variant="hideRead ? 'warning' : ''"
              ><span>{{ hideRead ? 'Show Read' : 'Hide Read' }}</span>
              <b-icon class="ml-1" :icon="hideRead ? 'eye' : 'eye-slash'"
            /></b-button>
          </div>
        </div>
      </b-sidebar>
    </div>
    <div class="d-flex align-items-center justify-content-between">
      <div
        v-if="selectedFilter"
        class="border rounded bg-white p-2 mt-2 d-flex justify-content-between align-items-center"
      >
        <div>
          <span class="font-weight-bold w-100"> Current filter: </span>
          <span>
            {{ selectedFilter.label }} -
            {{ selectedFilter.value === 'asc' ? 'Ascending' : 'Descending' }}
          </span>
        </div>
        <b-button class="ml-2" variant="danger" @click="clearFilter"
          >Clear Filter</b-button
        >
      </div>
    </div>

    <transition-group
      name="cell"
      tag="div"
      class="container mt-2"
      :class="selectedView.value === 'card' ? 'BookItemCardContainer' : ''"
    >
      <div
        v-for="book in books"
        :key="book.id"
        v-show="!hideRead || (hideRead && !book.read)"
      >
        <BookController :source="book" :view="selectedView.value" />
      </div>
      BookStore.books
    </transition-group>
  </div>
</template>
<script>
import BookController from '@/components/books/BookController'
import { BookStore } from '@/composables/book-provider.js'
import { Filters } from '@/composables/filter-provider.js'

export default {
  name: 'MyBooks',
  components: {
    BookController
  },
  data() {
    return {
      viewSelectVariant: Filters.viewSelectVariant,
      changeFilter: Filters.changeFilter,
      clearFilter: Filters.clearFilter,
      ...Filters.state,
      BookStore
    }
  },
  computed: {
    books() {
      return BookStore.getBooks.value
    },
    selectedFilter() {
      return Filters.selectedFilter.value
    }
  },
  methods: {
    addBook() {
      this.$router.push({ path: '/add-book' })
    }
  }
}
</script>
<style lang="scss" scoped>
@import '../../styles/_breakpoints.scss';

.FilterBar {
  display: grid;
  grid-gap: 5px;
  grid-template-columns: repeat(2, 1fr);
}

.BookItemCardContainer {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  @media (max-width: $medium) {
    grid-template-columns: repeat(2, 1fr);
  }
  @media (max-width: $small) {
    grid-template-columns: repeat(1, 1fr);
  }
}

.cell-move {
  transition: transform 0.5s;
}
</style>