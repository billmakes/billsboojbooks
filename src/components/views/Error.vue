<template>
  <div>
    <div>
      <h1 v-if="error">404</h1>
      <p>This book record either does not exist or it may have been deleted.</p>
      <router-link to="/">Click here to go back home.</router-link>
    </div>
  </div>
</template>