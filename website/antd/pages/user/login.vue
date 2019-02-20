<template>
  <a-row :gutter="24">
    <a-col :span="12">
      ss
    </a-col>
    <a-col :span="12">
      <a-form :form="user">
        <a-form-item
          :label-col="formItemLayout.labelCol"
          :wrapper-col="formItemLayout.wrapperCol"
          label="邮箱或手机"
        >
          <a-input
            v-decorator="[
              'input',
              {rules: [{ required: true, message: '请输入邮箱或手机!' }]}
            ]"
            placeholder="输入邮箱或手机号"
          />
        </a-form-item>
        <a-form-item
          :label-col="formItemLayout.labelCol"
          :wrapper-col="formItemLayout.wrapperCol"
          label="密码"
        >
          <a-input
            v-decorator="[
              'password',
              {rules: [{ required: true, message: '请输入密码!' }]}
            ]"

            placeholder="输入密码"
          />
        </a-form-item>
        <a-form-item
          :label-col="formTailLayout.labelCol"
          :wrapper-col="formTailLayout.wrapperCol"
        >
          <a-button
            type="primary"
            @click="check"
          >
            登录
          </a-button>
        </a-form-item>
      </a-form>
    </a-col>
  </a-row>
</template>

<script>
import axios from 'axios'
const formItemLayout = {
  labelCol: { span: 4 },
  wrapperCol: { span: 8 }
}
const formTailLayout = {
  labelCol: { span: 4 },
  wrapperCol: { span: 8, offset: 4 }
}
export default {
  data() {
    return {
      formItemLayout,
      formTailLayout,
      user: this.$form.createForm(this)
    }
  },
  methods: {
    check() {
      this.user.validateFields(err => {
        if (!err) {
          console.info(this.user.getFieldsValue())
          this.commit()
        }
      })
    },
    handleChange(e) {
      this.$nextTick(() => {
        this.user.validateFields(['password'], { force: true })
      })
    },
    commit: function() {
      const vm = this
      /*      const emailReg = new RegExp(
        '^([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+.[a-zA-Z]{2,3}$'
      )
      const phoneReg = /^1[0-9]{10}$/
      if (emailReg.test(vm.user.getFieldsValue().input)) {
        vm.user.setFieldsValue({ email: vm.user.getFieldsValue().input })
      } else if (phoneReg.test(vm.user.input)) {
        vm.user.setFieldsValue({ phone: vm.user.getFieldsValue().input })
      } */
      axios
        .post(`/api/user/login`, vm.user.getFieldsValue())
        .then(res => {
          //
          // success
          if (res.data.msg === '登录成功') {
            localStorage.setItem('token', res.data.token)
            vm.$store.commit('SET_USER', res.data.data)
            vm.$store.commit('SET_TOKEN', res.data.token)
            localStorage.setItem('user', res.data.data.id)
            vm.$toast('登录成功')
            vm.$router.replace('/')
          } else if (res.data.msg === '账号未激活') {
            vm.$toast(res.data.user.email)
          } else {
            vm.$toast(res.data.msg)
          }
        })
        .catch(function(err) {
          console.log(err)
        })
    },
    async login() {
      try {
        await this.$store.dispatch('login', {
          username: this.formUsername,
          password: this.formPassword
        })
        this.formUsername = ''
        this.formPassword = ''
        this.formError = null
      } catch (e) {
        this.formError = e.message
      }
    },
    async logout() {
      try {
        await this.$store.dispatch('logout')
      } catch (e) {
        this.formError = e.message
      }
    }
  }
}
</script>

<style scoped>
</style>
