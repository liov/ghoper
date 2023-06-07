<template>
  <a-row :gutter="24">
    <a-col :span="12">
      <nuxt-link to="/article/add">
        ss
      </nuxt-link>
    </a-col>
    <a-col
      v-if="!isLogin"
      :span="12"
    >
      <a-form :form="user">
        <a-form-item
          label=""
          :label-col="{ span: 5, offset: 5 }"
          :wrapper-col="{ span: 6, offset: 5 }"
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

        <a-form-item
          v-show="formType === 'signup'"
          :label-col="formItemLayout.labelCol"
          :wrapper-col="formItemLayout.wrapperCol"
          label="用户名"
        >
          <a-input
            v-decorator="[
              'name',
              {
                rules: [
                  { required: formType === 'signup', message: '请输入用户名!' }
                ]
              }
            ]"
            placeholder="请输入用户名！"
          />
        </a-form-item>

        <a-form-item
          v-show="formType === 'login'"
          :label-col="formItemLayout.labelCol"
          :wrapper-col="formItemLayout.wrapperCol"
          label="邮箱/手机"
        >
          <a-input
            v-decorator="[
              'input',
              {
                rules: [
                  {
                    required: formType === 'login',
                    message: '请输入邮箱或手机!'
                  }
                ]
              }
            ]"
            type="email"
            placeholder="输入邮箱或手机号！"
          />
        </a-form-item>
        <a-form-item
          v-show="formType === 'signup'"
          :label-col="formItemLayout.labelCol"
          :wrapper-col="formItemLayout.wrapperCol"
          label="邮箱"
        >
          <a-input
            v-decorator="[
              'email',
              {
                rules: [
                  { required: formType === 'signup', message: '请输入邮箱!' }
                ]
              }
            ]"
            type="email"
            placeholder="请输入邮箱！"
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
              { rules: [{ required: true, message: '请输入密码!' }] }
            ]"
            type="password"
            placeholder="请输入密码！"
          />
        </a-form-item>

        <a-form-item
          v-show="formType === 'signup'"
          label="性别"
          required
          :label-col="formItemLayout.labelCol"
          :wrapper-col="formItemLayout.wrapperCol"
        >
          <a-radio-group
            v-decorator="[
              'sex',
              {
                rules: [
                  { required: formType === true, message: '请输入手机号!' }
                ]
              }
            ]"
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
          v-show="formType === 'signup'"
          :label-col="formItemLayout.labelCol"
          :wrapper-col="formItemLayout.wrapperCol"
          label="手机号"
        >
          <a-input
            v-decorator="[
              'phone',
              {
                rules: [
                  { required: formType === 'signup', message: '请输入手机号!' }
                ]
              }
            ]"
            type="phone"
            placeholder="请输入手机号!"
          />
        </a-form-item>
        <a-form-item
          :label-col="{ span: 5, offset: 5 }"
          :wrapper-col="{ span: 6, offset: 5 }"
        >
          <div
            class="l-captcha"
            data-site-key="ff3498d2c6ffa1178cbf4fb6b445a8b3"
            data-width="200"
          />
        </a-form-item>
        <a-form-item
          :label-col="{ span: 6 }"
          :wrapper-col="{ span: 8, offset: 5 }"
        >
          <a-button
            type="primary"
            @click="check"
          >
            {{ formType === 'login' ? '登录' : '注册' }}
          </a-button>
        </a-form-item>
      </a-form>
    </a-col>
    <a-col
      v-if="isLogin"
      :span="12"
    >
      <nuxt-link to="/user/self">
        <a-button>
          个人信息
        </a-button>
      </nuxt-link>
      <a-button @click="logout">
        注销
      </a-button>
    </a-col>
  </a-row>
</template>

<script>
const formItemLayout = {
  labelCol: { span: 5 },
  wrapperCol: { span: 7 }
}
const formTailLayout = {
  labelCol: { span: 5 },
  wrapperCol: { span: 8, offset: 6 }
}
export default {
  data() {
    return {
      formType: 'login',
      formItemLayout,
      formTailLayout,
      user: this.$form.createForm(this),
      isLogin: false
    }
  },
  created() {
    if (this.$store.state.user !== null) {
      this.isLogin = true
    }
  },
  mounted() {
    if (this.$route.query.email !== null) {
      this.user.setFieldsValue({ input: this.$route.query.email })
    }

    const c = document.createElement('script');c.type = 'text/javascript';c.async = true;
    c.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'captcha.luosimao.com/static/dist/captcha.js?v=201812141420';
    const s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(c, s);
    s.parentNode.removeChild(c)
  },
  methods: {
    check() {
      this.user.validateFields((err) => {
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
      } else if (phoneReg.test(this.user.getFieldsValue().input)) {
        this.user.setFieldsValue({ phone: this.user.getFieldsValue().input })
      }
      /* this.$nextTick(() => {
        this.user.validateFields(['password'], { force: true })
      }) */
    },
    // 尊重语法糖，抛弃恶心的回调
    commit: function () {
      const vm = this
      this.$axios
        .$post(`/api/user/` + vm.formType, {...vm.user.getFieldsValue(),luosimao:document.getElementsByName("luotest_response")[0].value})
        .then((res) => {
          // success
          if (res.code === 200) {
            if (res.msg === '登录成功') {
              localStorage.setItem('token', res.token)
              vm.$store.commit('SET_USER', res.data)
              vm.$store.commit('SET_TOKEN', res.token)
              localStorage.setItem('user', res.data.id)
              vm.$message.info('登录成功')
              if (vm.$route.query.callbackUrl !== undefined) {
                vm.$router.replace(vm.$route.query.callbackUrl)
              } else {vm.$router.replace('/')}
            } else if (res.msg === '注册成功') {
              vm.$message.info('注册成功，请到邮箱激活')
            }
            // vm.$router.replace('/')
          } else if (res.msg === '账号未激活') {
            vm.$message.warning(res.user.email)
          } else {
            vm.$message.error(res.msg)
          }
        })
        .catch(function (err) {
          vm.$message.error(err)
        })
    },
    async logout() {
      const res = await this.$axios.$get('/api/user/logout')
      localStorage.removeItem('token')
      localStorage.removeItem('user')
      this.$store.commit('SET_USER', null)
      this.$store.commit('SET_TOKEN', null)
      this.$message.info(res.msg)
      this.isLogin = false
    }
  }
}
</script>

<style scoped></style>
