<template>
  <div class="article">
    <a-row>
      <a-col :span="4">
        <a-form-item
          label=""
          :label-col="{span: 4,offset:4}"
          :wrapper-col="{span: 24,offset:4}"
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
          :label-col="{span: 6}"
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
          :wrapper-col="{span:18}"
        >
          <a-upload
            multiple
            name="file"
            action="/api/upload/article"
            :before-upload="beforeUpload"
            @change="uploadChange"
          >
            <a-button>
              <a-icon type="upload" /> 上传封面
            </a-button>
          </a-upload>
        </a-form-item>
      </a-col>

      <a-button id="display" @click="showImage=!showImage">
        <span v-if="!showImage">显示封面</span>
        <span v-if="showImage">不显示封面</span>
      </a-button>
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
              placeholder="请选择分类"
              :default-value="[]"
              style="width: 200px"
            >
              <a-select-option v-for="(item,index) in existTags" :key="index">
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
              mode="multiple"
              placeholder="请选择标签"
              :default-value="[]"
              style="width: 200px"
            >
              <a-select-option v-for="(item,index) in existTags" :key="index">
                {{ item }}
              </a-select-option>
            </a-select>
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
              placeholder="请选择权限"
              :default-value="['0']"
              style="width: 200px"
            >
              <a-select-option value="0">
                全部可见
              </a-select-option>
              <a-select-option value="1">
                自己可见
              </a-select-option>
              <a-select-option value="2" disabled>
                部分可见
              </a-select-option>
              <a-select-option value="3">
                陌生人可见
              </a-select-option>
            </a-select>
          </a-form-item>
        </a-col>
      </a-row>
    </div>
    <div id="editor">
      <mavon-editor v-show="editorType==='markdown'" style="height: 20%" />

      <div v-show="editorType==='rich'" id="weditor" />
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
import ARow from 'ant-design-vue/es/grid/Row'

function getBase64(img, callback) {
  const reader = new FileReader()
  reader.addEventListener('load', () => callback(reader.result))
  reader.readAsDataURL(img)
}
export default {
  components: {
    ARow,
    mavonEditor
    // or 'mavon-editor': mavonEditor
  },
  data() {
    return {
      editorType: 'markdown',
      article: {},
      showImage: false,
      imageUrl: '',
      existTags: ['韩雪', '徐峥', '胡歌', '张卫健'],
      tagsGroup: [],
      Tags: [],
      tag: ''
    }
  },
  created() {
    this.tagsGroup = this.tagGroup(this.existTags, 3)
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
    },
    addTag: function() {
      if (this.tag !== '' && this.existTags.indexOf(this.tag) === -1) {
        this.existTags.push(this.tag)
        if (this.tagsGroup[this.tagsGroup.length - 1].length === 3) {
          this.tagsGroup.push([this.tag])
        } else {
          this.tagsGroup[this.tagsGroup.length - 1].push(this.tag)
        }
        this.Tags.push(this.tag)
        this.tag = ''
      } else if (this.tag === '') this.$toast('标签为空')
      else this.$toast('标签重复')
    },
    toggle(index) {
      this.$refs.checkboxes[index].toggle()
    },
    tagGroup: function(arr, size) {
      const arr2 = []
      for (let i = 0; i < arr.length; i = i + size) {
        arr2.push(arr.slice(i, i + size))
      }
      return arr2
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
#display {
  margin-top: 3px;
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
