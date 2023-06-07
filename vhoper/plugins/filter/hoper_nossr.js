import Vue from 'vue'
//import { DynamicScroller } from 'vue-virtual-scroller'
// import infiniteScroll from 'vue-infinite-scroll'
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css';
//Vue.component('DynamicScroller', DynamicScroller)

Vue.use(mavonEditor)
Vue.prototype.$md =mavonEditor.markdownIt
