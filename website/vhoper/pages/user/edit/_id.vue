<template>
  <a-row>
    <a-col :span="2" />
    <a-col :span="20" />
    <a-col :span="2" />
  </a-row>
</template>

<script>
export default {
  async asyncData({ $axios }) {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await $axios.$get(`/api/article`, { params })
    return {
      articleList: res.data,
      total: res.count
    }
  },
  methods: {
    getStatus: function() {
      fetch('http://hoper.xyz/user/1', {
        method: 'get',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({})
      }).then(res => {
        return res.json().status
      })
    }
  }
}
</script>

<style scoped>
</style>
