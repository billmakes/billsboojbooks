<template>
  <div id="app">
    <router-view name="nav-header" />
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="scss">
@import '~bootstrap/scss/_functions.scss';
@import '~bootstrap/scss/_variables.scss';
#app {
  background-color: $gray-100;
}
</style>
