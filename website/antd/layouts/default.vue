<template>
  <a-locale-provider :locale="zh_CN">
    <a-layout>
      <a-layout-header style="text-align: center">
        header
      </a-layout-header>
      <a-layout-content>
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
import moment from 'moment'
import 'moment/locale/zh-cn'

moment.locale('zh-cn')
export default {
  data() {
    return {
      zh_CN
    }
  }
}
</script>
