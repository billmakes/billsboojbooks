<template>
  <div>
    <b-card
      :title="source.title"
      :sub-title="`by ${source.author} - ${source.year}`"
      img-top
      tag="article"
      style="max-width: 20rem"
      class="m-2 shadow-sm"
    >
      <template #header>
        <div class="d-flex justify-content-between align-items-center">
          <BookRead :source="source" />
          <b-button :href="`book/${source.id}`" class="ml-2">Details</b-button>
        </div>
      </template>

      <b-card-text>
        <div v-if="source.tags.length">
          <b-badge
            v-for="(tag, index) in source.tags"
            :key="index"
            class="mr-1"
            variant="primary"
            >{{ tag }}</b-badge
          >
        </div>
      </b-card-text>
    </b-card>
  </div>
</template>
<script>
import BookRead from './shared/BookRead'
export default {
  name: 'BookCard',
  components: {
    BookRead
  },
  props: {
    source: {
      type: Object,
      required: true
    }
  }
}
</script>