<template>
  <div id="editor_t" />
</template>

<script>
import tinymce from 'tinymce/tinymce'
import 'tinymce/skins/ui/oxide/skin.min.css'
import 'tinymce/skins/ui/oxide/content.min.css'
import 'tinymce/skins/content/default/content.css'
import 'tinymce/themes/silver/theme'
import 'tinymce/plugins/image'
import 'tinymce/plugins/link'
import 'tinymce/plugins/code'
import 'tinymce/plugins/table'
import 'tinymce/plugins/lists'
import 'tinymce/plugins/contextmenu'
import 'tinymce/plugins/wordcount'
import 'tinymce/plugins/colorpicker'
import 'tinymce/plugins/textcolor'
export default {
  name: 'EditorT',
  mounted() {
    if (tinymce.activeEditor == null) {
      tinymce.init({
        selector: '#editor_t',
        language_url: '../tinymce/lang/zh_CN.js',
        language: 'zh_CN',
        skin: 'oxide',
        height: 650,
        plugins:
          'link lists image code table colorpicker textcolor wordcount contextmenu',
        toolbar:
          'bold italic underline strikethrough | fontsizeselect | forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist | outdent indent blockquote | undo redo | link unlink image code | removeformat',
        branding: false,
        menubar: true,
        // 此处为图片上传处理函数，这个直接用了base64的图片形式上传图片，
        // 如需ajax上传可参考https://www.tiny.cloud/docs/configure/file-image-upload/#images_upload_handler
        images_upload_handler: async (blobInfo, success, failure) => {
          const url = await upload('article', blobInfo.blob())
          success(url)
        }
        // images_upload_url: '/api/upload/article'
      })
    } else {
      tinymce.activeEditor.show()
    }
  }
}
</script>

<style scoped>
</style>
