<template>
  <div>
    <b-card
      class="BookCard m-2 shadow-sm"
      :title="source.title"
      :sub-title="`by ${source.author} - ${source.year}`"
    >
      <template #header>
        <div class="d-flex justify-content-between align-items-center">
          <BookRead :source="source" />
          <b-button @click="routeDetails" class="ml-2">Details</b-button>
        </div>
      </template>

      <b-card-text>
        <div v-if="source.tags.length">
          <BookTags :tags="source.tags" />
        </div>
      </b-card-text>
    </b-card>
  </div>
</template>
<script>
import BookRead from './shared/BookRead'
import BookTags from './shared/BookTags'

export default {
  name: 'BookCard',
  components: {
    BookRead,
    BookTags
  },
  props: {
    source: {
      type: Object,
      required: true
    }
  },
  methods: {
    routeDetails() {
      this.$router.push({
        name: 'book',
        params: { id: this.source.id, source: this.source }
      })
    }
  }
}
</script>
<style lang="scss" scoped>
.BookCard {
  height: 225px;
}
</style>