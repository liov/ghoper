<template>
  <div class="article">
    <a-row>
      <a-col :span="5">
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
                  <a-icon type="upload" /> 上传封面
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
            required
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
            required
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
      <mavon-editor v-show="editorType==='markdown'" ref="md" style="height: 650px" @imgAdd="imgAdd" @save="save" />
      <div v-show="editorType==='html'" id="weditor" />
    </div>
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
  middleware: 'auth',
  components: {
    mavonEditor
    // or 'mavon-editor': mavonEditor
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
      editor: {}
    }
  },
  created() {},
  mounted() {
    // const editor = new MediumEditor('.editable', {})
    this.editor = new E('#weditor')
    this.editor.customConfig.uploadImgServer = '/api/upload_multiple/article'
    this.editor.customConfig.height = '550px'
    this.editor.create()
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
        this.article.imageUrl = info.file.response.data.url
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
    imgAdd(pos, $file) {
      // 第一步.将图片上传到服务器.
      const formdata = new FormData()
      formdata.append('file', $file)
      axios({
        url: '/api/upload/article',
        method: 'post',
        data: formdata,
        headers: { 'Content-Type': 'multipart/form-data' }
      }).then(res => {
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        /**
         * $vm 指为mavonEditor实例，可以通过如下两种方式获取
         * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
         * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
         */
        this.$refs.md.$img2Url(pos, res.data.data.url)
      })
    },
    save(value) {
      const vm = this
      this.article.contentType = this.editorType
      if (this.editorType === 'markdown') this.article.content = value
      else {
        this.article.html_content = this.editor.txt.html()
        this.article.content = this.editor.txt.text()
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
          if (res.data.msg === '保存成功') vm.$router.push({ path: '/article' })
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
