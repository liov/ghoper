<template>
  <div>
    <div id="vditor" />
  </div>
</template>

<script>
import Vditor from 'vditor'

export default {
  mounted() {
    const vditor = new Vditor('vditor', { focus })
    vditor.focus()
  },
  methods: {}
}
</script>

<style scoped>
/*@import '~vditor/dist/index.classic.css';*/
@import '~vditor/dist/index.classic.css';
</style>
