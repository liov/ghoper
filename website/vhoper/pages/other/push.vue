<template>
  <div />
</template>

<script>
import Push from 'push.js'

export default {
  mounted() {
    const promise = Push.create('Hello world!', {
      body: "How's it hangin'?",
      icon: '/icon.png',
      timeout: 4000,
      onClick: function() {
        window.focus()
        this.close()
      }
    })
    promise.then(function(notification) {
      notification.close()
    })
    Push.clear()
  },
  methods: {}
}
</script>

<style scoped>
</style>
