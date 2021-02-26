<template>
  <b-navbar toggleable="lg" type="dark" variant="primary">
    <b-navbar-brand href="/">My Book List</b-navbar-brand>
  </b-navbar>
</template>
<script>
import routes from '@/router/routes'

export default {
  name: 'NavHeader',
  computed: {
    routeLinks() {
      return routes
    }
  }
}
</script>