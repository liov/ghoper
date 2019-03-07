<template>
  <a-row>
    <a-col :span="8" />
    <a-col :span="8" />
    <a-col :span="8" />
  </a-row>
</template>

<script>
export default {
  mounted() {
    const client = new EventSource('/api/sse/events')
    client.onmessage = function(evt) {
      console.log(evt)
    }
  }
}
</script>
<style scoped>
</style>
