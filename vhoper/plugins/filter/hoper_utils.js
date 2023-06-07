import Vue from 'vue'
import moment from 'moment'
//import dayjs from 'dayjs'
//import 'dayjs/locale/zh-cn'
import 'moment/locale/zh-cn'
import { upload } from '../utils/upload'
moment.locale('zh-cn')
//dayjs.locale('zh-cn')
Vue.filter('dateFormat', function (value) {
  return moment(value).format('YYYY-MM-DD HH:mm:ss')
})
Vue.prototype.$s2date = value => moment(value, 'YYYY-MM-DD HH:mm:ss.SSS Z')
Vue.prototype.$date2s = value => moment(value).format('YYYY-MM-DD HH:mm:ss')
Vue.prototype.$customUpload = async ({
  action,
  data,
  file,
  filename,
  headers,
  onError,
  onProgress,
  onSuccess,
  withCredentials,
  classify
}) => {
  const res = await upload(classify, file)
  onSuccess({ data: res, status: 200 }, file)
  file.status = 'done'
}

process.on('unhandledRejection', (reason, p) => {
  console.log('Unhandled Rejection at: Promise', p, 'reason:', reason)
  // application specific logging, throwing an error, or other logic here
})
