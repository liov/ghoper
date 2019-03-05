<template>
  <a-row>
    <a-col :span="6" style="text-align: right">
      <a-avatar shape="square" :size="100" :src="user.avatar_url" /><br>
      <a-upload
        name="file"
        action="/api/upload/avatar"
        :before-upload="beforeUpload"
        @change="uploadChange"
      >
        <a-button>
          <a-icon type="upload" />
          上传头像
        </a-button>
      </a-upload>
    </a-col>
    <a-col :span="16">
      <a-form-item
        label="用户名"
        :label-col="{ span: 2 }"
        :wrapper-col="{ span: 12 }"
      >
        <a-input
          v-model="user.name"
          disabled
        >
          <a-icon slot="prefix" type="user" />
        </a-input>
      </a-form-item>

      <a-form-item
        label="性别"
        :label-col="{ span: 2 }"
        :wrapper-col="{ span: 12 }"
      >
        <a-radio-group
          v-model="user.sex"
        >
          <a-radio-button value="男">
            男
          </a-radio-button>
          <a-radio-button value="女">
            女
          </a-radio-button>
        </a-radio-group>
      </a-form-item>

      <a-form-item
        label="邮箱"
        :label-col="{ span: 2 }"
        :wrapper-col="{ span: 12 }"
      >
        <a-input
          v-model="user.email"
          disabled
        >
          <a-icon slot="prefix" type="mail" />
        </a-input>
      </a-form-item>

      <a-form-item
        label="手机"
        :label-col="{ span: 2 }"
        :wrapper-col="{ span: 12 }"
      >
        <a-input
          v-model="user.phone"
          disabled
        >
          <a-icon slot="prefix" type="phone" />
        </a-input>
      </a-form-item>

      <a-row>
        <a-col :span="12" />
        <a-col :span="12" />
      </a-row>
    </a-col>
    <a-col :span="2" />
  </a-row>
</template>

<script>
export default {
  middleware: 'auth',
  async asyncData({ $axios }) {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await $axios.$get(`/api/user/edit`, { params })
    return {
      user: res.data
    }
  },
  created() {},
  methods: {
    getStatus: function() {
      fetch('http://hoper.xyz/user/1', {
        method: 'get',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({})
      }).then(res => {
        return res.json().status
      })
    },
    uploadChange(info) {
      if (info.file.status === 'uploading') {
        this.loading = true
        return
      }
      if (info.file.status === 'done') {
        this.user.avatar_url = info.file.response.data.url
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
    }
  }
}
</script>

<style scoped>
</style>
