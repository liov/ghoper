import Vue from 'vue'
import Vant from 'vant'
import 'vant/lib/index.css';


Vue.filter('dateFormat', function (value) {
    let date = new Date(value);
    return date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate() + ' ' + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds();
});
import lioHint from '../../pages/common/hint'
Vue.component('lio-hint',lioHint);
Vue.use(Vant);
