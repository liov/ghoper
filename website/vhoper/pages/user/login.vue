<template>
  <a-row :gutter="24">
    <a-col :span="12">
      <nuxt-link to="/article/add">
        ss
      </nuxt-link>
    </a-col>
    <a-col :span="12">
      <a-form :form="user">
        <a-form-item
          label=""
          :label-col="{span: 3,offset:6}"
          :wrapper-col="{span: 6,offset:6}"
        >
          <a-radio-group
            default-value="login"
            @change="handleChange"
          >
            <a-radio-button value="login">
              登录
            </a-radio-button>
            <a-radio-button value="signup">
              注册
            </a-radio-button>
          </a-radio-group>
        </a-form-item>
        <div v-show="formType==='login'">
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="邮箱或手机"
          >
            <a-input
              v-decorator="[
                'input',
                {rules: [{ required: formType==='login', message: '请输入邮箱或手机!' }]}
              ]"
              type="email"
              placeholder="输入邮箱或手机号！"
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
                {rules: [{ required: formType==='login', message: '请输入密码!' }]}
              ]"
              type="password"
              placeholder="请输入密码！"
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
        </div>
        <div v-show="formType==='signup'">
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="用户名"
          >
            <a-input
              v-decorator="[
                'name',
                {rules: [{ required: formType==='signup', message: '请输入用户名!' }]}
              ]"
              placeholder="请输入用户名！"
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
                {rules: [{ required: formType==='signup', message: '请输入密码!' }]}
              ]"
              type="password"
              placeholder="请输入密码！"
            />
          </a-form-item>
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="邮箱"
          >
            <a-input
              v-decorator="[
                'email',
                {rules: [{ required: formType==='signup', message: '请输入邮箱!' }]}
              ]"
              type="email"
              placeholder="请输入邮箱！"
            />
          </a-form-item>
          <a-form-item
            label="性别"
            required
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
          >
            <a-radio-group
              default-value="1"
            >
              <a-radio-button value="1">
                男
              </a-radio-button>
              <a-radio-button value="0">
                女
              </a-radio-button>
            </a-radio-group>
          </a-form-item>
          <a-form-item
            :label-col="formItemLayout.labelCol"
            :wrapper-col="formItemLayout.wrapperCol"
            label="手机号"
          >
            <a-input
              v-decorator="[
                'phone',
                {rules: [{ required: formType==='signup', message: '请输入手机号!' }]}
              ]"
              type="phone"
              placeholder="请输入手机号!"
            />
          </a-form-item>
          <a-form-item>
            <a-button
              type="primary"
              @click="check"
            >
              注册
            </a-button>
          </a-form-item>
        </div>
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
      formType: 'login',
      formItemLayout,
      formTailLayout,
      user: this.$form.createForm(this)
    }
  },
  created() {},
  mounted() {
    if (this.$route.query.email !== null) {
      this.user.setFieldsValue({ input: this.$route.query.email })
    }
  },
  methods: {
    check() {
      this.user.validateFields(err => {
        if (!err) {
          this.commit()
        }
      })
    },
    handleChange(e) {
      this.formType = e.target.value
      const emailReg = new RegExp(
        '^([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_.]?)*[a-zA-Z0-9]+.[a-zA-Z]{2,3}$'
      )
      const phoneReg = /^1[0-9]{10}$/
      if (emailReg.test(this.user.getFieldsValue().input)) {
        this.user.setFieldsValue({ email: this.user.getFieldsValue().input })
      } else if (phoneReg.test(this.user.input)) {
        this.user.setFieldsValue({ phone: this.user.getFieldsValue().input })
      }
      /* this.$nextTick(() => {
        this.user.validateFields(['password'], { force: true })
      }) */
    },
    commit: function() {
      const vm = this
      axios
        .post(`/api/user/` + vm.formType, vm.user.getFieldsValue())
        .then(res => {
          //
          // success
          if (res.data.code === 200) {
            if (res.data.msg === '登录成功') {
              localStorage.setItem('token', res.data.token)
              vm.$store.commit('SET_USER', res.data.data)
              vm.$store.commit('SET_TOKEN', res.data.token)
              localStorage.setItem('user', res.data.data.id)
              vm.$message.info('登录成功')
              if (vm.$route.query.callbackUrl !== undefined) {
                vm.$router.replace(vm.$route.query.callbackUrl)
              } else vm.$router.replace('/')
            } else if (res.data.msg === '注册成功') {
              vm.$message.info('注册成功，请到邮箱激活')
            }
            // vm.$router.replace('/')
          } else if (res.data.msg === '账号未激活') {
            vm.$message.warning(res.data.user.email)
          } else {
            vm.$message.error(res.data.msg)
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
