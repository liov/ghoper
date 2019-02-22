<template>
  <div>
    <div ref="editor" />
  </div>
</template>

<script>
import E from 'wangeditor'
import 'wangeditor/release/wangEditor.css'
export default {
  name: 'Editor',
  mounted() {
    const editor = new E(this.$refs.editor)
    editor.customConfig.uploadImgServer = '/api/upload_multiple/article'
    editor.create()
  }
}
</script>

<style scoped>
</style>
