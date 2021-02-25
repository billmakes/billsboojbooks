<template>
  <div>
    Read on {{ computedDate }}
    <b-icon icon="check" />
  </div>
</template>
<script>
export default {
  name: 'BookRead',
  props: {
    readDate: {
      type: String,
      default: null
    }
  },
  computed: {
    computedDate() {
      var readDate = new Date(this.readDate)
      var dd = String(readDate.getDate()).padStart(2, '0')
      var mm = String(readDate.getMonth() + 1).padStart(2, '0') //January is 0!
      var yyyy = readDate.getFullYear()
      let date = mm + '/' + dd + '/' + yyyy
      return date
    }
  }
}
</script>