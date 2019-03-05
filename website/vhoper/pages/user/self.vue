<template>
  <a-row>
    <a-col :span="3" style="text-align: right">
      <a-avatar shape="square" :size="100" :src="user.avatar_url" />
      <br>
      <a-upload
        name="file"
        action="/api/upload/avatar"
        :before-upload="beforeUpload"
        @change="uploadAvatarChange"
      >
        <a-button>
          <a-icon type="upload" />
          上传头像
        </a-button>
      </a-upload>
      <a-upload
        name="file"
        action="/api/upload/cover"
        :before-upload="beforeUpload"
        @change="uploadCoverChange"
      >
        <a-button>
          <a-icon type="upload" />
          上传背景
        </a-button>
      </a-upload>
    </a-col>
    <a-col
      :span="15"
      :style="{background:'url('+user.cover_url+') no-repeat',
               backgroundSize: 'cover'}"
    >
      <a-form-item
        label="用户名"
        :label-col="formItemLayout.labelCol"
        :wrapper-col="formItemLayout.wrapperCol"
      >
        <a-input
          v-model="user.name"
          disabled
        >
          <a-icon slot="prefix" type="user" />
        </a-input>
      </a-form-item>
      <a-row>
        <a-col :span="10">
          <a-form-item
            label="性别"
            :label-col="{ span: 12 }"
            :wrapper-col="{ span: 10 }"
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
        </a-col>
        <a-col :span="13">
          <a-form-item
            label="生日"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span:10 }"
          >
            <a-date-picker show-time />
          </a-form-item>
        </a-col>
      </a-row>


      <a-form-item
        label="邮箱"
        :label-col="formItemLayout.labelCol"
        :wrapper-col="formItemLayout.wrapperCol"
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
        :label-col="formItemLayout.labelCol"
        :wrapper-col="formItemLayout.wrapperCol"
      >
        <a-input
          v-model="user.phone"
          disabled
        >
          <a-icon slot="prefix" type="phone" />
        </a-input>
      </a-form-item>
      <a-form-item
        label="个人简介"
        :label-col="formItemLayout.labelCol"
        :wrapper-col="formItemLayout.wrapperCol"
      >
        <a-input
          v-model="user.introduction"
        >
          <a-icon slot="prefix" type="profile" />
        </a-input>
      </a-form-item>
      <a-form-item
        label="个人签名"
        :label-col="formItemLayout.labelCol"
        :wrapper-col="formItemLayout.wrapperCol"
      >
        <a-input
          v-model="user.signature"
        >
          <a-icon slot="prefix" type="profile" />
        </a-input>
      </a-form-item>
      <a-form-item
        v-for="(item,index) in schools"
        :key="index"
        label="教育经历"
        :label-col="formItemLayout.labelCol"
        :wrapper-col="{ span: 18 }"
      >
        <a-input-group>
          <a-col :span="7">
            <a-input v-model="schools[index].name" placeholder="请输入学习名称!" />
          </a-col>
          <a-col :span="7">
            <a-input v-model="schools[index].speciality" placeholder="请输入专业!" />
          </a-col>
          <a-col :span="7">
            <a-range-picker>
              <a-icon slot="suffixIcon" type="smile" />
            </a-range-picker>
          </a-col>
          <a-button icon="plus" @click="addSchool" />
        </a-input-group>
      </a-form-item>
      <a-form-item
        v-for="(item,index) in careers"
        :key="index+schools.length"
        label="职业经历"
        :label-col="formItemLayout.labelCol"
        :wrapper-col="{ span: 18 }"
      >
        <a-input-group>
          <a-col :span="7">
            <a-input v-model="careers[index].company" placeholder="请输入公司!" />
          </a-col>
          <a-col :span="7">
            <a-input v-model="careers[index].title" placeholder="请输入职务!" />
          </a-col>
          <a-col :span="7">
            <a-range-picker>
              <a-icon slot="suffixIcon" type="smile" />
            </a-range-picker>
          </a-col>
          <a-button icon="plus" @click="addCareer" />
        </a-input-group>
      </a-form-item>
      <a-row>
        <a-col :span="12" />
        <a-col :span="12" />
      </a-row>
    </a-col>
    <a-col :span="4">
      关&nbsp;&nbsp;&nbsp;注：{{ user.follow_count }}<br>
      粉&nbsp;&nbsp;&nbsp;丝：{{ user.followed_count }}<br>
      积&nbsp;&nbsp;&nbsp;分：{{ user.Score }}<br>
      文&nbsp;&nbsp;&nbsp;章：{{ user.article_count }}<br>
      瞬&nbsp;&nbsp;&nbsp;间：{{ user.moment_count }}<br>
      日记本：{{ user.diary_book_count }}<br>
      日&nbsp;&nbsp;&nbsp;记：{{ user.diary_count }}
    </a-col>
  </a-row>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      formItemLayout: {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 }
      },
      schools: [{ name: '', speciality: '', start_time: '', end_time: '' }],
      careers: [{ company: '', title: '', start_time: '', end_time: '' }],
      loading: false
    }
  },
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
    uploadAvatarChange(info) {
      if (info.file.status === 'uploading') {
        this.loading = true
        return
      }
      if (info.file.status === 'done') {
        this.user.avatar_url = info.file.response.data.url
        this.loading = false
      }
    },
    uploadCoverChange(info) {
      if (info.file.status === 'uploading') {
        this.loading = true
        return
      }
      if (info.file.status === 'done') {
        this.user.cover_url = info.file.response.data.url
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
    addSchool() {
      if (this.schools.length < 5) {
        this.schools.push({
          name: '',
          speciality: '',
          start_time: '',
          end_time: ''
        })
        return
      }
      this.$message.warning('添加过多!')
    },
    addCareer() {
      if (this.careers.length < 5) {
        this.careers.push({
          company: '',
          title: '',
          start_time: '',
          end_time: ''
        })
        return
      }
      this.$message.warning('添加过多!')
    }
  }
}
</script>

<style scoped>
</style>
