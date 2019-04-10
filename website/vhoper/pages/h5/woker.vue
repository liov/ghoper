<template>
  <div>
    <div>Web Woker</div>
    <div id="result" />
    <a-button @click="stopWorker">
      停止
    </a-button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      w: undefined
    }
  },
  mounted() {
    this.startWorker()
    const vm = this
    setTimeout(function() {
      vm.stopWorker()
    }, 3000)
  },
  methods: {
    startWorker: function() {
      if (typeof Worker !== 'undefined') {
        if (typeof this.w === 'undefined') {
          this.w = new Worker('../js/demo_workers.js')
        }
        this.w.onmessage = function(event) {
          document.getElementById('result').innerHTML = event.data
        }
      } else {
        document.getElementById('result').innerHTML =
          '抱歉，你的浏览器不支持 Web Workers...'
      }
    },
    stopWorker: function() {
      this.w.terminate()
      this.w = undefined
    }
  }
}
</script>

<style scoped>
</style>
