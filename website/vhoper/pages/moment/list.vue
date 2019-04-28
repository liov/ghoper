<template>
  <div>
    <a-radio-group v-model="mode" :style="{ marginBottom: '8px' }">
      <a-radio-button value="top">
        Horizontal
      </a-radio-button>
      <a-radio-button value="left">
        Vertical
      </a-radio-button>
    </a-radio-group>
    <a-tabs
      default-active-key="1"
      :tab-position="mode"
      @prevClick="callback"
      @nextClick="callback"
    >
      <a-tab-pane key="1" tab="Tab 1">
        <a-table
          :data-source="momentList"
          :row-key="record => record.id"
          :row-selection="rowSelection"
          :pagination="pagination"
          :loading="loading"
          @change="handleTableChange"
        >
          <a-table-column key="id" title="ID" data-index="id" />
          <a-table-column key="user.name" title="作者" data-index="user.name" />
          <a-table-column
            key="user.avatar_url"
            title="头像"
            data-index="user.avatar_url"
          />
          <a-table-column-group>
            <span slot="title" style="color: #1890ff">喜欢</span>
            <a-table-column key="collect_count" data-index="collect_count">
              <span slot="title" style="color: #1890ff">收藏</span>
            </a-table-column>
            <a-table-column
              key="like_count"
              title="喜欢"
              data-index="like_count"
            />
            <a-table-column
              key="approve_count"
              title="点赞"
              data-index="approve_count"
            />
            <a-table-column
              key="comment_count"
              title="评论"
              data-index="comment_count"
            />
          </a-table-column-group>

          <a-table-column key="tags" title="标签" data-index="tags">
            <template slot-scope="tags">
              <span>
                <a-tag v-for="tag in tags" :key="tag.name" color="blue">{{
                  tag.name
                }}</a-tag>
              </span>
            </template>
          </a-table-column>
          <a-table-column key="action" title="操作">
            <template slot-scope="text, record">
              <span>
                <a href="javascript:;">Action 一 {{ record.id }}</a>
                <a-divider type="vertical" />
                <a href="javascript:;">删除</a>
              </span>
            </template>
          </a-table-column>
        </a-table>
      </a-tab-pane>
      <a-tab-pane key="2" tab="Tab 2">
        Content of tab 2
      </a-tab-pane>
      <a-tab-pane key="3" tab="Tab 3">
        Content of tab 3
      </a-tab-pane>
      <a-tab-pane key="4" tab="Tab 4">
        Content of tab 4
      </a-tab-pane>
      <a-tab-pane key="5" tab="Tab 5">
        Content of tab 5
      </a-tab-pane>
      <a-tab-pane key="6" tab="Tab 6">
        Content of tab 6
      </a-tab-pane>
      <a-tab-pane key="7" tab="Tab 7">
        Content of tab 7
      </a-tab-pane>
      <a-tab-pane key="8" tab="Tab 8">
        Content of tab 8
      </a-tab-pane>
      <a-tab-pane key="9" tab="Tab 9">
        Content of tab 9
      </a-tab-pane>
      <a-tab-pane key="10" tab="Tab 10">
        Content of tab 10
      </a-tab-pane>
      <a-tab-pane key="11" tab="Tab 11">
        Content of tab 11
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script>
export default {
  data() {
    return {
      mode: 'top',
      shchqi2: undefined,
      pagination: {},
      loading: false
    }
  },
  computed: {
    rowSelection() {
      const { selectedRowKeys } = this
      return {
        onChange: (selectedRowKeys, selectedRows) => {
          console.log(
            `selectedRowKeys: ${selectedRowKeys}`,
            'selectedRows: ',
            selectedRows
          )
        },
        getCheckboxProps: record => ({
          props: {
            disabled: record.name === 'Disabled User', // Column configuration not to be checked
            name: record.name
          }
        })
      }
    }
  },
  async asyncData({ $axios }) {
    const params = {
      pageNo: 0,
      pageSize: 10
    }
    const res = await $axios.$get(`/api/moment`, { params })
    return {
      momentList: res.data,
      total: res.count,
      topCount: res.top_count
    }
  },
  created() {
    this.pagination.total = this.total
    this.shchqi2 = (function*(a) {
      for (let i = 0; i < a.length; i++) {
        yield i
      }
    })(this.momentList).next()
  },
  methods: {
    callback(val) {
      console.log(val)
    },
    handleTableChange(pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      this.pagination = pager
      const pageNo = pagination.current - 1
      fetch('/api/moment?pageNo=' + pageNo + '&pageSize=' + pagination.pageSize)
        .then(response => response.json())
        .then((data) => {
          this.loading = false
          this.momentList = data.data
          this.pagination.total = data.count
        })
        .catch(error => console.log('error is', error))
    },
    // 闭包，匿名函数，生成器
    bibao: function*() {
      for (let i = 0; i < this.momentList.length; i++) {
        yield i
      }
    },
    niming() {
      let i = 0
      ;(function f() {
        i++
      })()
      return i
    },
    shchqi() {
      let i = 0
      return function () {
        i++
        return i
      }
    }
  }
}
</script>

<style scoped></style>
