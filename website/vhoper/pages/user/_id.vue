<template>
  <a-row>
    <a-col :span="8" />
    <a-col :span="8">
      <a-avatar size="large" :src="user.avatar_url" />
    </a-col>
    <a-col :span="8" />
  </a-row>
</template>

<script>
export default {
  data() {
    return {
      user: null
    }
  },
  async asyncData({ $axios, route, redirect }) {
    const res = await $axios.$get(route.path).catch(() => {
      redirect({ path: '/user/self' })
    })
    return {
      user: res.data,
      address: res.data.address.split(' ')
    }
  }
}
</script>

<style scoped>
</style>
