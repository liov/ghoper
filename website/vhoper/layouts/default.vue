<template>
  <a-locale-provider :locale="zh_CN">
    <a-layout>
      <div class="header" style="background-color: #fff">
        <div class="menu">
          <a-menu
            v-model="current"
            :theme="theme"
            mode="horizontal"
            :style="{ lineHeight: '64px' }"
          >
            <a-menu-item key="home">
              <nuxt-link to="/">
                <a-icon type="home" />
                主页
              </nuxt-link>
            </a-menu-item>
            <a-menu-item key="file-text">
              <nuxt-link to="/article">
                <a-icon type="file-text" />
                博客
              </nuxt-link>
            </a-menu-item>

            <a-menu-item key="message">
              <nuxt-link to="/chat/v2">
                <a-icon type="message" />
                聊天
              </nuxt-link>
            </a-menu-item>

            <a-menu-item key="picture">
              <nuxt-link to="/moment">
                <a-icon type="picture" />
                瞬间
              </nuxt-link>
            </a-menu-item>
            <a-menu-item key="user">
              <nuxt-link to="/user/login">
                <a-icon type="user" />
                <span v-if="$store.state.user">注销</span>
                <span v-else>登录</span>
              </nuxt-link>
            </a-menu-item>
            <!--<a-menu-item key="book">
              <nuxt-link to="/diary">
                <a-icon type="book" /> 日记
              </nuxt-link>
            </a-menu-item>-->
            <a-sub-menu>
              <span slot="title" class="submenu-title-wrapper"><a-icon type="setting" />设置</span>
              <a-menu-item-group title="初始化">
                <a-menu-item key="setting:1">
                  <nuxt-link to="/tpl/init">
                    数据库初始化
                  </nuxt-link>
                </a-menu-item>
                <a-menu-item key="setting:2">
                  设置初始化
                </a-menu-item>
              </a-menu-item-group>
              <a-menu-item-group title="危险操作">
                <a-menu-item key="setting:3">
                  重启
                </a-menu-item>
                <a-menu-item key="setting:4">
                  关闭
                </a-menu-item>
              </a-menu-item-group>
            </a-sub-menu>
            <a-menu-item key="app" disabled>
              <a-icon type="appstore" />
              Hoper
            </a-menu-item>
            <a-switch
              :default-checked="false"
              @change="changeTheme"
            />
            主题
          </a-menu>
        </div>
      </div>
      <a-layout-content style="background-color: #fff">
        <nuxt />
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        hoper ©2019 Created by JYB
      </a-layout-footer>
    </a-layout>
  </a-locale-provider>
</template>
<script>
import zh_CN from 'ant-design-vue/lib/locale-provider/zh_CN'

export default {
  data() {
    return {
      current: ['main'],
      theme: 'light',
      style: { backgroundColor: '#fff' },
      zh_CN
    }
  },
  destroyed() {
    sessionStorage.setItem('back_url', this.$route.path)
  },
  methods: {
    changeTheme(checked) {
      this.theme = checked ? 'dark' : 'light'
      this.style = checked
        ? { height: '48px' }
        : { backgroundColor: '#fff', height: '48px' }
    }
  }
}
</script>
<style scoped>
.header {
  height: 64px;
}

.menu {
  width: 100%;
  text-align: center;
  position: fixed;
  top: 0;
  z-index: 1;
}
</style>
