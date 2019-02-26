<template>
  <div />
</template>

<script>
import axios from 'axios'
export default {
  mounted() {
    const vm = this
    axios
      .get(
        `/api/user/active/${this.$route.params.id}/${this.$route.params.secret}`
      )
      .then(res => {
        vm.$message.info(res.data.msg)
        vm.$router.push({ path: '/user/sign' })
      })
  }
}
</script>

<style scoped>
</style>
