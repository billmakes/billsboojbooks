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
          <h4 class="mb-0">
            <b-icon :icon="source.read ? 'book-fill' : 'book'" />
          </h4>
          <BookRead v-if="source.read" :read-date="source.readDate" />
          <b-button :href="`book/${source.id}`">Details</b-button>
        </div>
      </template>

      <b-card-text>
        card
        {{ `${source.title} by ${source.author} (${source.year})` }}
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