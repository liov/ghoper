<template>
  <div>
    <div id="tag">
      <a-row>
        <a-col :span="5">
          <a-form-item
            label="心情"
            :label-col="{span:5,offset:6}"
            :wrapper-col="{span: 9}"
          >
            <a-input
              v-model="moment.mood_name"
            />
          </a-form-item>
        </a-col>
        <a-col :span="5">
          <a-form-item
            label="标签"
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
        <a-col :span="5">
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
        <a-col :span="5">
          <a-form-item
            label="权限"
            :label-col="{span: 4}"
            :wrapper-col="{span:6}"
          >
            <a-select
              v-model="moment.permission"
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
        <a-col :span="2">
          <a-button icon="save" style="margin-top: 3px" @click="commit">
            保存
          </a-button>
        </a-col>
      </a-row>

      <a-form-item style="width: 80%">
        <div id="vditor" style="margin: 0 10%" />
      </a-form-item>
      <a-form-item style="margin-left: 10%;">
        <a-upload
          action="/api/upload/moment"
          list-type="picture-card"
          :multiple="true"
          :file-list="imgList"
          :before-upload="beforeUpload"
          :custom-request="customUpload"
          @preview="handlePreview"
          @change="uploadChange"
        >
          <div v-if="imgList.length < 9">
            <a-icon type="plus" />
            <div class="ant-upload-text">
              图片
            </div>
          </div>
        </a-upload>
        <a-modal :visible="previewVisible" :footer="null" @cancel="handleCancel">
          <img alt="example" style="width: 100%" :src="previewImage">
        </a-modal>
      </a-form-item>
    </div>
  </div>
</template>

<script>
import { upload } from '../../plugins/utils/upload'
import 'vditor/dist/index.classic.css'
let vditor
export default {
  middleware: 'auth',
  data() {
    return {
      moment: {
        image_url: '',
        mood_name: '',
        tags: [],
        permission: 0
      },
      imgList: [],
      existTags: [],
      tag: '',
      categories: [],
      tags: [],
      previewVisible: false,
      previewImage: ''
    }
  },
  async asyncData({ $axios }) {
    const params = { pageNO: 0, pageSize: 10 }
    const tres = await $axios.$get(`/api/tag`, { params })
    return {
      existTags: tres.data
    }
  },
  created() {},
  mounted: function() {
    this.$nextTick(function() {
      const Vditor = require('vditor')
      vditor = new Vditor('vditor', {
        width: '100%',
        preview: { delay: 0 },
        upload: { url: '/api/upload/moment' }
      })
      vditor.focus()
    })
  },
  methods: {
    uploadChange({ fileList }) {
      // 果然如我所料，双向绑定的锅
      if (!fileList[fileList.length - 1].response) {
        this.imgList = fileList
        return
      }
      // 2. read from response and show file link
      fileList = fileList.map(file => {
        if (file.response) {
          // Component will show file.url as link
          file.url = file.response.data.url
          /* if (this.moment.image_url.indexOf(file.response.data.url) === -1) {
            if (this.moment.image_url === '')
              this.moment.image_url = file.response.data.url
            else
              this.moment.image_url =
                this.moment.image_url + ',' + file.response.data.url
          } */
        }
        return file
      })
      // 3. filter successfully uploaded files according to response from server
      /*      fileList = fileList.filter(file => {
          if (file.response) {
            return file.response.data.code === 200
          }
          return false
        }) */
      this.imgList = fileList
      this.loading = false
    },
    handleCancel() {
      this.previewVisible = false
    },
    handlePreview(file) {
      this.previewImage = file.url || file.thumbUrl
      this.previewVisible = true
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
    commit: function() {
      const vm = this
      this.moment.permission = parseInt(this.moment.permission)
      this.moment.content = vditor.getValue()
      if (this.moment.image_url !== '')
        this.moment.image_url.substring(0, this.moment.image_url.length - 1)
      this.moment.tags = []
      for (const i of this.tags) {
        this.moment.tags.push({ name: i })
      }
      this.$axios
        .$post(`/api/moment`, this.moment)
        .then(function(res) {
          // success
          if (res.code === 200) {
            vm.$router.push({ path: '/moment' })
            vditor.clearCache()
          } else vm.$message.error(res.msg)
        })
        .catch(function(err) {
          vm.$message.error(err)
        })
    },
    customUpload: async function({
      action,
      data,
      file,
      filename,
      headers,
      onError,
      onProgress,
      onSuccess,
      withCredentials
    }) {
      const res = await upload('moment', file)
      if (this.moment.image_url === '') this.moment.image_url = res.url
      else this.moment.image_url = this.moment.image_url + ',' + res.url
      onSuccess({ data: res, status: 200 }, file)
      file.status = 'done'
    }
  }
}
</script>

<style scoped>
</style>
