<template>
  <v-app>
    <default-bar />

    <default-view class="body"/>
  </v-app>
</template>

<script setup>
  import DefaultBar from './AppBar.vue'
	import DefaultView from './View.vue'
</script>

<style scoped>
.body {
  display: flex;
  justify-content: center;
  align-items: center;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto',
  'Helvetica Neue', 'Ubuntu', sans-serif;
  height: 100vh;
  margin: 0;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>