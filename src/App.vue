<template>
  <div id="app">
    <router-view name="nav-header" />
    <b-container>
      <router-view />
    </b-container>
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
body {
  background-color: $gray-100;
}
</style>
