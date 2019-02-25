<template>
  <a-row>
    <a-col :span="6">
      col-8
    </a-col>
    <a-col :span="12">
      col-8
    </a-col>
    <a-col :span="6">
      col-8
    </a-col>
  </a-row>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      pageNo: 1,
      pageNum: 1,
      pageSize: 5,
      lastFlag: false,
      firstFlag: true
    }
  },
  async asyncData() {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await axios.get(`/api/article`, { params })
    return { articleList: res.data.data }
  }
}
</script>

<style scoped>
</style>
