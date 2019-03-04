<template>
  <a-row>
    <a-col :span="3">
      <a-form-item
        label=""
        :label-col="{span: 4,offset:2}"
        :wrapper-col="{span: 24,offset:2}"
      >
        <nuxt-link to="/moment/add">
          <a-button icon="edit">
            添加
          </a-button>
        </nuxt-link>
      </a-form-item>
    </a-col>
    <a-col :span="21">
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
            <template slot="actions">
              <span>回复</span>
              <span v-if="item.user.id=user.id">编辑</span>
            </template>
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
  </a-row>
</template>

<script>
export default {
  filters: {},
  data() {
    return {
      pageSizeOptions: ['5', '10', '15', '20'],
      current: 1,
      pageSize: 5,
      user: null
    }
  },
  computed: {},
  watch: {
    current: async function() {
      await this.next(this.current - 1)
    }
  },
  async asyncData({ $axios }) {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await $axios.$get(`/api/moment`, { params })
    return {
      momentList: res.data,
      total: res.count,
      topCount: res.top_count
    }
  },
  created: function() {
    this.user = this.$store.state.user
  },
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
      const res = await this.$axios.$get(`/api/moment`, { params })

      const momentList = res.data

      this.momentList = momentList
    }
  }
}
</script>

<style scoped>
</style>
