import Vue from 'vue'
import moment from 'moment'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'
import 'moment/locale/zh-cn'
moment.locale('zh-cn')
dayjs.locale('zh-cn')
Vue.filter('dateFormat', function(value) {
  return moment(value).format('YYYY-MM-DD HH:mm:ss')
})
Vue.prototype.$s2date = value => moment(value, 'YYYY-MM-DD HH:mm:ss.SSS Z')
Vue.prototype.$date2s = value => moment(value).format('YYYY-MM-DD HH:mm:ss')
