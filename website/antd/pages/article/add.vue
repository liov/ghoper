<template>
  <div>
    <div id="vditor" />
    <div id="editor">
      <mavon-editor style="height: 100%" />
    </div>
  </div>
</template>

<script>
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import Vditor from 'vditor'
import 'vditor/dist/index.classic.css'
export default {
  components: {
    mavonEditor
    // or 'mavon-editor': mavonEditor
  },
  mounted() {
    const vditor = new Vditor('vditor', { focus })
    vditor.focus()
  },
  methods: {}
}
</script>

<style scoped>
/*@import '~vditor/dist/index.classic.css';*/
/*@import '~vditor/dist/index.classic.css';*/
#editor {
  margin: auto;
  width: 80%;
  height: 580px;
}
</style>
