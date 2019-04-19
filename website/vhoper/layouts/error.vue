<template>
  <div class="container" style="text-align: center">
    <h1 v-if="error.statusCode === 404">
      页面不存在
    </h1>
    <h1 v-else>
      {{ error }}
    </h1>
    <nuxt-link to="/">
      首 页
    </nuxt-link>
  </div>
</template>

<script>
export default {
  props: ['error']
}
</script>

<style scoped>
</style>
