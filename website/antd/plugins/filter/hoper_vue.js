import Vue from 'vue'
// import Vant from 'vant';
import Antd from 'ant-design-vue/lib'
import lioHint from '../../components/hint'
// import Vuetify from 'vuetify'

// import Vant from 'vant'
// import 'vant/lib/index.css';
// import Mint from 'mint-ui';
// import 'mint-ui/lib/style.css';

Vue.filter('dateFormat', function(value) {
  const date = new Date(value)
  return (
    date.getFullYear() +
    '-' +
    (date.getMonth() + 1) +
    '-' +
    date.getDate() +
    ' ' +
    date.getHours() +
    ':' +
    date.getMinutes() +
    ':' +
    date.getSeconds()
  )
})
Vue.component('lio-hint', lioHint)

Vue.use(Antd)
// Vue.use(Vuetify)
// Vue.use(Mint);
