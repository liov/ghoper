import Vue from 'vue'
import moment from 'moment'
import Antd from 'ant-design-vue/lib'
import lioHint from '../../components/hint'

Vue.filter('dateFormat', function(value) {
  return moment(value).format('YYYY-MM-DD HH:mm:ss')
  /* date.getFullYear() +
    '-' +
    (date.getMonth() + 1) +
    '-' +
    date.getDate() +
    ' ' +
    date.getHours() +
    ':' +
    date.getMinutes() +
    ':' +
    date.getSeconds() */
})
Vue.component('lio-hint', lioHint)

Vue.use(Antd)
