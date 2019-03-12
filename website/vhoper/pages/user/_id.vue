<template>
  <div>
    <div style="marginBottom:16px;text-align: center">
      <span style="marginRight:6px">Gutter (px): </span>
      <div style="width:50%; margin:0 auto;">
        <a-slider
          v-model="gutterKey"
          :min="0"
          :max="Object.keys(gutters).length - 1"
          :marks="gutters"
          :step="null"
        />
      </div>
      <span style="marginRight:6px;text-align: center">列数:</span>
      <div style="width:50%; margin:0 auto;">
        <a-slider
          v-model="colCountKey"
          :min="0"
          :max="Object.keys(colCounts).length - 1"
          :marks="colCounts"
          :step="null"
        />
      </div>
    </div>
    <a-row :gutter="gutters[gutterKey]">
      <a-col v-for="(item) in colCounts[colCountKey]" :key="item.toString()" :span="24/colCounts[colCountKey]">
        <div> <a-avatar size="large" :src="user.avatar_url" /></div>
        <div>{{ user.name }}</div>
        <div>{{ user.id }}</div>
        <div>{{ user.phone }}</div>
      </a-col>
    </a-row>
  </div>
</template>

<script>
export default {
  data() {
    const gutters = {}
    const arr = [8, 16, 24, 32, 40, 48]
    arr.forEach((value, i) => {
      gutters[i] = value
    })
    const colCounts = {}
    const arr1 = [2, 3, 4, 6, 8, 12]
    arr1.forEach((value, i) => {
      colCounts[i] = value
    })
    return {
      gutterKey: 1,
      colCountKey: 2,
      colCounts,
      gutters
    }
  },
  computed: {},
  async asyncData({ $axios, route, redirect }) {
    const res = await $axios.$get(`/api` + route.path).catch(() => {
      redirect({ path: '/user/self' })
    })
    return {
      user: res.data
    }
  }
}
</script>

<style scoped>
</style>
