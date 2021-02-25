<template>
  <div>
    <router-link
      v-for="link in routeLinks"
      :key="link.name"
      :to="{ name: link.name }"
      >{{ link.label }}</router-link
    >
  </div>
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