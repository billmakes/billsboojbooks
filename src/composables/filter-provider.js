import { reactive, computed, readonly } from '@vue/composition-api'
import { BookStore } from './book-provider'

const state = reactive({
  filters: [
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
  ],
  viewOptions: [
    { value: 'card', icon: 'grid-3x3-gap-fill' },
    { value: 'tile', icon: 'list-ul' }
  ],
  selectedView: { value: 'card', icon: 'grid-3x3-gap-fill' },
  hideRead: false
})

function changeFilter(filter) {
  state.filters.forEach(f => {
    if (f.field !== filter.field) f.value = null
  })
  if (!filter.value) {
    filter.value = 'asc'
    BookStore.filterAsc(filter.field)
    return
  }
  if (filter.value === 'asc') {
    filter.value = 'desc'
    BookStore.filterDesc(filter.field)
  } else {
    filter.value = 'asc'
    BookStore.filterAsc(filter.field)
  }
}

function toggleHideRead() {
  state.hideRead = !state.hideRead
}

function viewSelectVariant(option) {
  return state.selectedView === option ? 'primary' : 'none'
}
function clearFilter() {
  state.hideRead = false
  state.filters.forEach(f => {
    f.value = null
  })
}

const selectedFilter = computed(() =>
  state.filters.find(filter => filter.value)
)

export const Filters = readonly({
  toggleHideRead,
  changeFilter,
  viewSelectVariant,
  clearFilter,
  selectedFilter,
  state
})
