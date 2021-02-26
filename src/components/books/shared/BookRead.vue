<template>
  <div class="d-flex align-items-center justify-content-between">
    <template v-if="markedRead || source.read">
      <b-icon class="mb-0 mr-1 h3" icon="check" variant="success" />

      <div v-if="markedRead || source.read">Read on {{ computedDate }}</div>
    </template>
    <template v-else>
      <div
        class="MarkReadBtn d-flex justify-content-center align-items-center"
        style="cursor: pointer"
        @mouseenter="showReadOption = true"
        @mouseleave="showReadOption = false"
        @click="markRead"
      >
        <b-icon icon="book" class="h4 mb-0" />
        <transition name="slide-fade">
          <span v-show="showReadOption" class="ml-1">Mark read?</span>
        </transition>
      </div>
    </template>
  </div>
</template>
<script>
import BookService from '@/services/'

export default {
  name: 'BookRead',
  props: {
    source: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      showReadOption: false,
      markedRead: null
    }
  },
  computed: {
    computedDate() {
      var readDate = new Date(this.source.readDate)
      var dd = String(readDate.getDate()).padStart(2, '0')
      var mm = String(readDate.getMonth() + 1).padStart(2, '0') //January is 0!
      var yyyy = readDate.getFullYear()
      let date = mm + '/' + dd + '/' + yyyy
      return date
    }
  },
  methods: {
    markRead() {
      const { id, author, title, year } = this.source
      this.markedRead = true
      BookService.updateBook(id, {
        author,
        title,
        year,
        read: true
      })
    }
  }
}
</script>
<style>
.MarkReadBtn:hover {
  text-decoration: underline;
}
.slide-fade-enter-active {
  transition: all 0.3s ease;
}
.slide-fade-leave-active {
  transition: all 0.3s ease;
}
.slide-fade-enter,
.slide-fade-leave-to {
  transform: translateX(-10px);
  opacity: 0;
}
</style>