<template>
  <div>
    <a-row>
      <a-col :span="4">
        <a-form-item
          label=""
        >
          <a-radio-group
            default-value="markdown"
            @change="handleChange"
          >
            <a-radio-button value="markdown">
              markdown
            </a-radio-button>
            <a-radio-button value="rich">
              富文本
            </a-radio-button>
          </a-radio-group>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="标题"
          required
          :label-col="{span: 2}"
          :wrapper-col="{span: 18}"
        >
          <a-input
            v-model="article.title"
            size="large"
            placeholder="请输入标题！"
          />
        </a-form-item>
      </a-col>
      <a-col :span="6">
        <a-form-item
          label="封面"
          :label-col="{span: 2}"
          :wrapper-col="{span: 18}"
        >
          <a-upload
            name="file"
            list-type="picture-card"
            class="avatar-uploader"
            :show-upload-list="false"
            action="/api/upload/article"
            :before-upload="beforeUpload"
            @change="uploadChange"
          >
            <img v-if="imageUrl" :src="imageUrl" alt="avatar">
            <div v-else>
              <a-icon :type="loading ? 'loading' : 'plus'" />
              <div class="ant-upload-text">
                Upload
              </div>
            </div>
          </a-upload>
        </a-form-item>
      </a-col>
    </a-row>



    <div v-show="editorType==='markdown'" id="editor">
      <mavon-editor style="height: 20%" />
    </div>
    <div v-show="editorType==='rich'" id="weditor" />
  </div>
</template>

<script>
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
// import MediumEditor from 'medium-editor'
// import 'medium-editor/dist/css/medium-editor.min.css'
import E from 'wangeditor'
import 'wangeditor/release/wangEditor.css'
import axios from 'axios'

function getBase64(img, callback) {
  const reader = new FileReader()
  reader.addEventListener('load', () => callback(reader.result))
  reader.readAsDataURL(img)
}
export default {
  components: {
    mavonEditor
    // or 'mavon-editor': mavonEditor
  },
  data() {
    return {
      editorType: 'markdown',
      article: {},
      loading: false,
      imageUrl: '',
      formItemLayout: {
        labelCol: { span: 4 },
        wrapperCol: { span: 8 }
      }
    }
  },
  mounted() {
    // const editor = new MediumEditor('.editable', {})
    const editor = new E('#weditor')
    editor.create()
  },
  methods: {
    handleChange(e) {
      this.editorType = e.target.value
    },
    uploadChange(info) {
      if (info.file.status === 'uploading') {
        this.loading = true
        return
      }
      if (info.file.status === 'done') {
        // Get this url from response in real world.
        getBase64(info.file.originFileObj, imageUrl => {
          this.imageUrl = imageUrl
          this.loading = false
        })
      }
    },
    beforeUpload(file) {
      const isImg = /image\//.test(file.type)
      if (!isImg) {
        this.$message.error('You can only upload JPG file!')
      }
      const isLt2M = file.size / 1024 / 1024 < 3
      if (!isLt2M) {
        this.$message.error('Image must smaller than 4MB!')
      }
      return isImg && isLt2M
    }
  }
}
</script>

<style scoped>
/*@import '~vditor/dist/index.classic.css';*/
/*@import '~vditor/dist/index.classic.css';*/
#editor {
  margin: auto;
  width: 100%;
  height: 300px;
}
</style>
