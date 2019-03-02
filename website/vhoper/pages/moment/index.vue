<template>
  <a-row>
    <a-col :span="6">
      <a-form-item
        label=""
        :label-col="{span: 4,offset:2}"
        :wrapper-col="{span: 24,offset:2}"
      >
        <nuxt-link to="/moment/add">
          <a-icon type="edit" />
        </nuxt-link>
      </a-form-item>
    </a-col>
    <a-col :span="12">
      <a-list
        class="comment-list"
        item-layout="horizontal"
        :data-source="momentList"
      >
        <a-list-item slot="renderItem" slot-scope="item">
          <a-comment
            :author="item.user.name"
            :avatar="item.user.avatar_url"
          >
            <!--   <template slot="actions">
              <span v-for="(action,index) in item.actions" :key="index">{{ action }}</span>
            </template>-->
            <p slot="content">
              {{ item.content }}
            </p>
            <a-tooltip slot="datetime" :title="item.created_at">
              <span>{{ item.created_at|dateFormat }}</span>
            </a-tooltip>
          </a-comment>
        </a-list-item>
      </a-list>
      <a-pagination
        v-model="current"
        :page-size-options="pageSizeOptions"
        :total="total"
        show-quick-jumper
        show-size-changer
        :page-size="pageSize"
        @showSizeChange="onShowSizeChange"
      >
        <template slot="buildOptionText" slot-scope="props">
          <span v-if="props.value!=='50'">{{ props.value }}条/页</span>
          <span v-if="props.value==='50'">全部</span>
        </template>
      </a-pagination>
    </a-col>
    <a-col :span="6" />
  </a-row>
</template>

<script>
import axios from 'axios'

export default {
  filters: {},
  data() {
    return {
      pageSizeOptions: ['5', '10', '15', '20'],
      current: 1,
      pageSize: 5
    }
  },
  computed: {},
  watch: {
    current: async function() {
      await this.next(this.current - 1)
    }
  },
  async asyncData() {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await axios.get(`/api/moment`, { params })
    return {
      momentList: res.data.data,
      total: res.data.count,
      topCount: res.data.top_count
    }
  },
  created: function() {},
  mounted: function() {},
  methods: {
    async onShowSizeChange(current, pageSize) {
      this.pageSize = pageSize
      await this.next(current - 1)
    },
    setMoment: function(moment) {
      localStorage.setItem('moment_' + moment.id, moment)
    },
    next: async function(pageNo) {
      const params = {
        pageNo: pageNo,
        pageSize: this.pageSize
      }

      // 这里可以这么写，async，await函数，或者 return axios().then((res)=>{})返回Promise
      const res = await axios.get(`/api/moment`, { params })

      const momentList = res.data.data

      this.momentList = momentList
    }
  }
}
</script>

<style scoped>
</style>
