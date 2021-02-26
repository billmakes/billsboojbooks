<template>
  <div class="border rounded my-2 p-2">
    <div class="d-flex justify-content-between align-items-center">
      <div>
        <span class="font-weight-bold">{{ source.title }}</span>
        <span>{{ ` by ${source.author} - ${source.year}` }}</span>
      </div>
      <div class="d-flex">
        <BookRead :source="source" />
        <b-button :href="`book/${source.id}`" class="ml-2">Details</b-button>
      </div>
    </div>
    <div class="d-flex justify-content-between"></div>
  </div>
</template>
<script>
import BookRead from './shared/BookRead'
export default {
  name: 'tile',
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
<style lang="scss" scoped>
</style>