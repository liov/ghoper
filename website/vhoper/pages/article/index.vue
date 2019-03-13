<template>
  <a-row>
    <a-col :span="2">
      <nuxt-link to="/article/add">
        添加
      </nuxt-link>
      <nuxt-link to="/article/edit">
        编辑
      </nuxt-link>
    </a-col>
    <a-col :span="20">
      <a-breadcrumb>
        <a-breadcrumb-item>主页</a-breadcrumb-item>
        <a-breadcrumb-item>
          <nuxt-link to="">
            博客
          </nuxt-link>
        </a-breadcrumb-item>
        <a-breadcrumb-item>列表</a-breadcrumb-item>
      </a-breadcrumb>
      <a-list
        item-layout="vertical"
        size="large"
        :data-source="articleList"
      >
        <div slot="footer">
          <b />
        </div>
        <a-list-item slot="renderItem" key="item.title" slot-scope="item">
          <span slot="actions" @click="star(item.id)">
            <a-icon type="star-o" style="margin-right: 8px" />
            {{ item.collect_count }}
          </span>
          <span slot="actions" @click="like(item.id)">
            <a-icon type="like-o" style="margin-right: 8px" />
            {{ item.like_count }}
          </span>
          <span slot="actions" @click="comment(item.id)">
            <a-icon type="message" style="margin-right: 8px" />
            {{ item.comment_count }}
          </span>

          <a-button-group slot="actions">
            <a-button v-for="(subitem,subindex) in item.categories" :key="subindex">
              {{ subitem.name }}
            </a-button>
          </a-button-group>

          <div slot="actions">
            <a-tag v-for="(subitem,subindex) in item.tags" :key="subindex" :color="color[subindex]">
              {{ subitem.name }}
            </a-tag>
          </div>

          <img v-if="item.image_url!==''" slot="extra" height="120" alt="logo" :src="item.image_url">
          <a-list-item-meta
            :description="item.intro"
          >
            <a-row slot="title">
              <a-col :span="3" style="font-size: 10px">
                <nuxt-link :to="'/user/'+item.user.id">
                  {{ item.user.name }}
                </nuxt-link>
              </a-col>
              <nuxt-link :to="'/article/'+item.id">
                <a-col :span="15" style="color:rgba(0, 0, 0, 0.85)">
                  {{ item.title }}
                </a-col>
              </nuxt-link>
              <a-col :span="6" style="font-size: 10px">
                <span> {{ item.created_at|dateFormat }}</span>
                <span>{{ $s2date(item.created_at).fromNow() }}</span>
              </a-col>
            </a-row>

            <nuxt-link slot="avatar" :to="'/user/'+item.user.id">
              <a-avatar :src="item.user.avatar_url" alt="头像" />
            </nuxt-link>
          </a-list-item-meta>
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
    <a-col :span="2">
      col-8
    </a-col>
  </a-row>
</template>

<script>
// nuxt-link组件不能包多个col
export default {
  data() {
    return {
      pageSizeOptions: ['5', '10', '15', '20'],
      pageSize: 5,
      current: 1,
      color: ['pink', 'red', 'orange', 'orange', 'cyan', 'blue', 'purple']
    }
  },
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
    const res = await $axios.$get(`/api/article`, { params })
    return {
      articleList: res.data,
      total: res.count
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
      const res = await this.$axios.$get(`/api/article`, { params })
      this.articleList = res.data
      this.total = res.count
    },
    star(id) {},
    like(id) {},
    comment(id) {}
  }
}
</script>

<style scoped>
</style>
