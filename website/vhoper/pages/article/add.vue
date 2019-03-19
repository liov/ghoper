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
            v-model="editorType"
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
                :custom-request="$customUpload"
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
              style="width: 200px"
            >
              <a-select-option v-for="item in existCategories" :key="item.id">
                {{ item.name }}
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
              style="width: 200px"
            >
              <a-select-option v-for="item in existTags" :key="item.name">
                {{ item.name }}
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
      <no-ssr placeholder="Loading...">
        <mavon-editor
          v-show="editorType==='markdown'"
          ref="md"
          style="height: 650px"
          @imgAdd="imgAdd"
          @save="save"
        />
      </no-ssr>
      <div v-show="editorType==='html'">
        <div id="editor_t" />
      </div>
    </div>
  </div>
</template>

<script>
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import 'tinymce/skins/ui/oxide/skin.min.css'
import 'tinymce/skins/ui/oxide/content.min.css'
import '../../static/css/content.css'
import { upload } from '../../plugins/utils/upload'

export default {
  middleware: 'auth',
  components: {
    mavonEditor
  },
  data() {
    return {
      editorType: 'html',
      article: {
        categories: [],
        tags: [],
        permission: 0
      },
      showImage: false,
      imageUrl: '',
      tag: '',
      categories: [],
      tags: [],
      init: {
        selector: '#editor_t',
        language_url: '../tinymce/lang/zh_CN.js',
        language: 'zh_CN',
        skin: 'oxide',
        height: 650,
        plugins: 'link lists image code table wordcount ',
        toolbar:
          'bold italic underline strikethrough | fontsizeselect | forecolor backcolor | alignleft aligncenter alignright alignjustify | bullist numlist | outdent indent blockquote | undo redo | link unlink image code | removeformat',
        branding: false,
        menubar: true,
        // 此处为图片上传处理函数，这个直接用了base64的图片形式上传图片，
        // 如需ajax上传可参考https://www.tiny.cloud/docs/configure/file-image-upload/#images_upload_handler
        images_upload_handler: async (blobInfo, success, failure) => {
          const data = await upload('article', blobInfo.blob())
          success(data.url)
        },
        convert_urls: false
        // images_upload_url: '/api/upload/article'
      }
    }
  },
  async asyncData({ $axios }) {
    const cres = await $axios.$get(`/api/category`)
    const params = { pageNO: 0, pageSize: 10 }
    const tres = await $axios.$get(`/api/tag`, { params })
    return {
      existCategories: cres.data,
      existTags: tres.data,
      categories: [cres.data[0].id]
    }
  },
  created() {},
  mounted() {
    /*
      this.component = () => ({
        component: import(`mavon-editor`)
      })
  */
    /*    Vue.component('mavon-editor', resolve => {
        // 这个特殊的 `require` 语法将会告诉 webpack
        // 自动将你的构建代码切割成多个包，这些包
        // 会通过 Ajax 请求加载
        require(['mavon-editor'], ({ mavonEditor }) => resolve(mavonEditor))
      }) */
    // 这个函数真是无敌
    this.$nextTick(function() {
      require('../../plugins/filter/tinymce')
      tinymce.init(this.init)
    })
  },
  beforeDestroy() {
    tinymce.activeEditor.destroy()
  },
  methods: {
    handleChange(e) {
      this.editorType = e.target.value

      // 手动创建有bug，切换路由回来不渲染了,得destroy()了,另一个组件测试可以用show，无语
      // 两种方式，1.插件客户端渲染，2.有window后引入
      /*      if (typeof window !== 'undefined') {
          /!*        const Vditor = require('vditor')
            const vditor = new Vditor('content', {})
            vditor.focus() *!/
          require('../../plugins/filter/tinymce')
          if (e.target.value === 'html' && tinymce.activeEditor === null) {
            tinymce.init(this.init)
          }
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
        this.imageUrl = info.file.response.data.url
        this.loading = false
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
      const vm = this
      if (this.tag === '') {
        this.$message.error('标签为空')
        return
      }
      for (const v of this.existTags) {
        if (v.name === vm.tag) {
          vm.$message.error('标签重复')
          return
        }
      }
      this.existTags.push({ name: this.tag })
      this.tags.push(this.tag)
      this.tag = ''
    },
    async imgAdd(pos, $file) {
      // 第一步.将图片上传到服务器.
      const data = await upload('article', $file)
      // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
      /**
       * $vm 指为mavonEditor实例，可以通过如下两种方式获取
       * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
       * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
       */
      this.$refs.md.$img2Url(pos, data.url)
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
      this.$axios
        .$post(`/api/article`, this.article)
        .then(function(res) {
          // success
          if (res.code === 200) vm.$router.push({ path: '/article' })
          else vm.$message.error(res.msg)
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
