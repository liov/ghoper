<template>
  <div>
    <a-row>
      <a-col :span="12">
        <a-form-item
          label="封面"
          :label-col="{span: 6}"
          :wrapper-col="{span:15}"
        >
          <a-row>
            <a-col :span="12">
              <a-upload
                name="file"
                action="/api/upload/moment"
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
            label="心情"
            :label-col="{span:6,offset:6}"
            :wrapper-col="{span: 9}"
          >
            <a-input
              v-model="moment.mood_name"
            />
          </a-form-item>
        </a-col>
        <a-col :span="9">
          <a-form-item
            label="标签"
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
      </a-row>
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
      <a-form-item style="width: 80%">
        <a-textarea v-model="moment.content" placeholder="请输入" autosize style="margin: 0 10%" />
      </a-form-item>
      <a-button icon="save" @click="commit">
        保存
      </a-button>
    </div>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      moment: {
        mood_name: '',
        tags: [],
        permission: 0
      },
      showImage: false,
      imageUrl: '',
      existCategories: ['小说', '散文', '戏剧', '诗歌'],
      existTags: ['韩雪', '徐峥', '胡歌', '张卫健'],
      tag: '',
      categories: [],
      tags: []
    }
  },
  methods: {
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
    commit: function() {
      const vm = this
      this.moment.permission = parseInt(this.moment.permission)
      this.moment.tags = []
      for (const i of this.tags) {
        this.moment.tags.push({ name: i })
      }
      this.$axios
        .$post(`/api/moment`, this.moment)
        .then(function(res) {
          // success
          if (res.code === 200) vm.$router.push({ path: '/moment' })
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
</style>
