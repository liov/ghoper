import Vue from 'vue'
import { DynamicScroller } from 'vue-virtual-scroller'
import infiniteScroll from 'vue-infinite-scroll'

export default () => {
  Vue.use(infiniteScroll)
  Vue.component('DynamicScroller', DynamicScroller)
}
