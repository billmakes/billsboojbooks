<template>
  <div>
    <b-badge
      v-for="(tag, index) in tags"
      :key="index"
      class="mr-1"
      variant="primary"
      >{{ tag }}</b-badge
    >
  </div>
</template>
<script>
export default {
  name: 'BookTags',
  props: {
    tags: {
      type: Array,
      default: () => []
    }
  }
}
</script>