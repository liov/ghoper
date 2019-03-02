<template>
  <a-row>
    <a-col :span="8">
      <nuxt-link to="/article/add">
        添加
      </nuxt-link>
      <nuxt-link to="/article/edit">
        编辑
      </nuxt-link>
    </a-col>
    <a-col :span="8">
      <a-list
        item-layout="vertical"
        size="large"
        :data-source="articleList"
      >
        <div slot="footer">
          <b />
        </div>
        <a-list-item slot="renderItem" key="item.title" slot-scope="item">
          <span slot="actions">
            <a-icon type="star-o" style="margin-right: 8px" />
            {{ item.collect_count }}
          </span>
          <span slot="actions">
            <a-icon type="like-o" style="margin-right: 8px" />
            {{ item.love_count }}
          </span>
          <span slot="actions">
            <a-icon type="message" style="margin-right: 8px" />
            {{ item.comment_count }}
          </span>
          <img v-if="item.image_url!==''" slot="extra" height="120" alt="logo" :src="item.image_url">
          <a-list-item-meta
            :description="item.intro"
          >
            <span slot="title" href="/">
              <a-row>
                <a-col :span="6" style="font-size: 10px">
                  {{ item.user.name }}
                </a-col>
                <a-col :span="12">
                  {{ item.title }}
                </a-col>
                <a-col :span="6">
                  <span style="font-size: 10px"> {{ item.created_at|dateFormat }}</span>
                </a-col>
              </a-row>
            </span>
            <a-avatar slot="avatar" :src="item.image_url" />
          </a-list-item-meta>
          {{ item.content }}
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
    <a-col :span="8">
      col-8
    </a-col>
  </a-row>
</template>

<script>
import axios from 'axios'
import ARow from 'ant-design-vue/es/grid/Row'
export default {
  components: { ARow },
  data() {
    return {
      pageSizeOptions: ['5', '10', '15', '20'],
      pageSize: 5,
      current: 1
    }
  },
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
    const res = await axios.get(`/api/article`, { params })
    return {
      articleList: res.data.data,
      total: res.data.count
    }
  },
  methods: {
    async onShowSizeChange(current, pageSize) {
      this.pageSize = pageSize
      await this.next(current - 1)
    },
    async next(pageNo) {
      const params = {
        pageNo: pageNo,
        pageSize: this.pageSize
      }
      const res = await axios.get(`/api/article`, { params })
      this.articleList = res.data.data
      this.total = res.data.count
    }
  }
}
</script>

<style scoped>
</style>
