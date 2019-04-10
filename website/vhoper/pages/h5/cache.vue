<template>
  <div>
    <div>CacheStorage</div>
    <div id="status" />
  </div>
</template>

<script>
export default {
  mounted() {
    if ('serviceWorker' in navigator) {
      document.querySelector(
        '#status'
      ).innerHTML += `<p>浏览器是否支持：支持</p>`
      navigator.serviceWorker
        .register('../js/demo_cache.js', {
          scope: '/js/'
        })
        .then(function(registration) {
          document.querySelector(
            '#status'
          ).innerHTML += `<p>service worker是否注册成功：注册成功</p>`

          let serviceWorker
          if (registration.installing) {
            serviceWorker = registration.installing
            document.querySelector(
              '#status'
            ).innerHTML += `<p>当前注册状态：installing</p>`
          } else if (registration.waiting) {
            serviceWorker = registration.waiting
            document.querySelector(
              '#status'
            ).innerHTML += `<p>当前注册状态：waiting</p>`
          } else if (registration.active) {
            serviceWorker = registration.active
            document.querySelector(
              '#status'
            ).innerHTML += `<p>当前注册状态：active</p>`
          }
          if (serviceWorker) {
            document.querySelector(
              '#status'
            ).innerHTML += `<p>当前service worker状态：&emsp; ${
              serviceWorker.state
            }</p>`
            serviceWorker.addEventListener('statechange', function(e) {
              document.querySelector('#status').innerHTML += `<p>状态变化为:${
                e.target.state
              }</p>`
            })
          }
        })
        .catch(function(error) {
          console.log(error)
        })
    } else {
      document.querySelector(
        '#status'
      ).innerHTML += `<p>浏览器是否支持：不支持</p>`
    }
  },
  methods: {}
}
</script>

<style scoped>
</style>
