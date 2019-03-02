<template>
  <div class="article">
    <a-row>
      <a-col :span="5" style="z-index: 0">
        <a-form-item
          label=""
          :label-col="{span: 4,offset:2}"
          :wrapper-col="{span: 24,offset:2}"
        >
          <a-radio-group
            default-value="markdown"
            @change="handleChange"
          >
            <a-radio-button value="markdown">
              markdown
            </a-radio-button>
            <a-radio-button value="html">
              富文本
            </a-radio-button>
          </a-radio-group>
          <a-button v-if="editorType==='html'" @click="save">
            <a-icon type="save" />
          </a-button>
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          label="标题"
          required
          :label-col="{span: 4}"
          :wrapper-col="{span: 18}"
        >
          <a-input
            v-model="article.title"
            placeholder="请输入标题！"
          />
        </a-form-item>
      </a-col>
      <a-col :span="6">
        <a-form-item
          label="封面"
          :label-col="{span: 6}"
          :wrapper-col="{span:15}"
        >
          <a-row>
            <a-col :span="16">
              <a-upload
                name="file"
                action="/api/upload/article"
                :before-upload="beforeUpload"
                @change="uploadChange"
              >
                <a-button>
                  <a-icon type="upload" />
                  上传封面
                </a-button>
              </a-upload>
            </a-col>
            <a-col :span="8">
              <a-button class="formbuttion" @click="showImage=!showImage">
                <span v-if="!showImage">显示封面</span>
                <span v-if="showImage">不显示封面</span>
              </a-button>
            </a-col>
          </a-row>
        </a-form-item>
      </a-col>
    </a-row>

    <div align="center">
      <img v-if="showImage" :src="imageUrl">
    </div>
    <div id="tag">
      <a-row>
        <a-col :span="6">
          <a-form-item
            label="分类"
            required
            :label-col="{span: 4}"
            :wrapper-col="{span:6}"
          >
            <a-select
              v-model="categories"
              mode="multiple"
              placeholder="请选择分类"
              :default-value="[]"
              style="width: 200px"
            >
              <a-select-option v-for="(item,index) in existCategories" :key="index">
                {{ item }}
              </a-select-option>
            </a-select>
          </a-form-item>
        </a-col>
        <a-col :span="6">
          <a-form-item
            label="标签"
            required
            :label-col="{span: 4}"
            :wrapper-col="{span: 6}"
          >
            <a-select
              v-model="tags"
              mode="multiple"
              placeholder="请选择标签"
              :default-value="[]"
              style="width: 200px"
            >
              <a-select-option v-for="item in existTags" :key="item">
                {{ item }}
              </a-select-option>
            </a-select>
          </a-form-item>
        </a-col>
        <a-col :span="6">
          <a-form-item
            label="新标签"
            :label-col="{span:6}"
            :wrapper-col="{span: 12}"
          >
            <a-row>
              <a-col :span="16">
                <a-input
                  v-model="tag"
                />
              </a-col>
              <a-col :span="6">
                <a-button class="formbuttion" @click="addTag">
                  添加
                </a-button>
              </a-col>
            </a-row>
          </a-form-item>
        </a-col>
        <a-col :span="6">
          <a-form-item
            label="权限"
            :label-col="{span: 4}"
            :wrapper-col="{span:6}"
          >
            <a-select
              v-model="article.permission"
              placeholder="请选择权限"
              :default-value="[0]"
              style="width: 200px"
            >
              <a-select-option :key="0">
                全部可见
              </a-select-option>
              <a-select-option :key="1">
                自己可见
              </a-select-option>
              <a-select-option :key="2" disabled>
                部分可见
              </a-select-option>
              <a-select-option :key="3" disabled>
                陌生人可见
              </a-select-option>
            </a-select>
          </a-form-item>
        </a-col>
      </a-row>
    </div>
    <div id="editor">
      <mavon-editor
        v-show="editorType==='markdown'"
        ref="md"
        v-model="html_code"
        style="height: 650px"
        @imgAdd="imgAdd"
        @save="save"
      />
      <!--    <div v-show="editorType==='html'" id="weditor" />-->
      <div v-show="editorType==='html'">
        <!--   <editor :init="init" />-->
        <!--        <editor-t />-->
        <div id="editor_t" />
      </div>
    </div>
  </div>
</template>

<script>
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import axios from 'axios'
// import Editor from '@tinymce/tinymce-vue'
import tinymce from 'tinymce/tinymce'
// import EditorT from '../../components/Tinymce'
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
import { upload, getBase64 } from '../../plugins/utils/upload'

export default {
  middleware: 'auth',
  components: {
    mavonEditor
    // EditorT
  },
  data() {
    return {
      editorType: 'markdown',
      article: {
        categories: [],
        tags: [],
        permission: 0
      },
      showImage: false,
      imageUrl: '',
      existCategories: ['小说', '散文', '戏剧', '诗歌'],
      existTags: ['韩雪', '徐峥', '胡歌', '张卫健'],
      tag: '',
      categories: [],
      tags: [],
      init: {
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
        },
        convert_urls: false
        // images_upload_url: '/api/upload/article'
      }
    }
  },
  created() {},
  mounted() {
    if (tinymce.activeEditor !== null) {
      tinymce.activeEditor.destroy()
    }
    tinymce.init(this.init)
  },
  methods: {
    handleChange(e) {
      this.editorType = e.target.value

      /* if (e.target.value === 'html' && tinymce.activeEditor === null) {
                  // 手动创建有bug，切换路由回来不渲染了,得destroy()了,另一个组件测试可以用show，无语
                  tinymce.init(this.init)
                } */
    },
    uploadChange(info) {
      if (info.file.status === 'uploading') {
        this.loading = true
        return
      }
      if (info.file.status === 'done') {
        this.article.image_url = info.file.response.data.url
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
        this.$message.error('只能上传图片!')
      }
      const isLt2M = file.size / 1024 / 1024 < 4
      if (!isLt2M) {
        this.$message.error('不能超过 4MB!')
      }
      return isImg && isLt2M
    },
    addTag: function() {
      if (this.tag !== '' && this.existTags.indexOf(this.tag) === -1) {
        this.existTags.push(this.tag)
        this.article.tags.push(this.tag)
        this.tag = ''
      } else if (this.tag === '') this.$message.error('标签为空')
      else this.$message.error('标签重复')
    },
    async imgAdd(pos, $file) {
      // 第一步.将图片上传到服务器.
      const url = await upload('article', $file)
      // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
      /**
       * $vm 指为mavonEditor实例，可以通过如下两种方式获取
       * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
       * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
       */
      this.$refs.md.$img2Url(pos, url)
    },
    save(value, render) {
      const vm = this
      this.article.contentType = this.editorType
      if (this.editorType === 'markdown') {
        this.article.html_content = render
        this.article.content = value
      } else {
        this.article.html_content = tinymce.activeEditor.getContent()
        this.article.content = tinymce.activeEditor.getContent({
          format: 'text'
        })
      }

      for (const i of this.tags) {
        this.article.tags.push({ name: i })
      }
      for (const i of this.categories) {
        this.article.categories.push({ id: i })
      }
      axios
        .post(`/api/article`, this.article)
        .then(function(res) {
          // success
          if (res.data.code === 200) vm.$router.push({ path: '/article' })
          else vm.$message.error(res.data.msg)
        })
        .catch(function(err) {
          vm.$message.error(err)
        })
    }
  }
}
</script>

<style scoped>
/*@import '~vditor/dist/index.classic.css';*/
/*@import '~vditor/dist/index.classic.css';*/
.article {
  width: 80%;
  margin-left: 10%;
}

#editor {
  margin: auto;
}

.formbuttion {
  margin-top: 4px;
}

#tag {
  position: relative;
  z-index: 2;
}

#editor {
  position: relative;
  z-index: 1;
}
</style>
