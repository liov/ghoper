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
        <a-affix :offset-top="top">
          <a-button type="primary" @click="showDrawer">
            固钉
          </a-button>
        </a-affix>
        <a-back-top />
        <a-drawer
          title="Basic Drawer"
          placement="right"
          :closable="false"
          :visible="visible"
          @close="onClose"
        >
          <p>Some contents...</p>
          <p>Some contents...</p>
          <p>Some contents...</p>
        </a-drawer>
      </a-form-item>
    </a-col>
    <a-col :span="21">
      <a-list
        class="comment-list"
        item-layout="horizontal"
        :data-source="momentList"
      >
        <a-list-item slot="renderItem" slot-scope="item,index">
          <a-comment>
            <nuxt-link slot="author" :to="'/user/'+item.user.id">
              <span>{{ item.user.name }}</span>
            </nuxt-link>
            <nuxt-link slot="avatar" :to="'/user/'+item.user.id">
              <a-avatar :src="item.user.avatar_url" alt="头像" />
            </nuxt-link>
            <span slot="actions" style="margin-right: 10px" @click="star(item.id,index)">
              <a-icon type="star-o" />
              收藏：{{ item.collect_count }}
            </span>
            <span slot="actions" style="margin-right: 10px" @click="like(item.id,index)">
              <a-icon type="like-o" />
              喜欢：{{ item.like_count }}
            </span>
            <span slot="actions" style="margin-right: 10px" @click="comment(item.id,index)">
              <a-icon type="message" />
              回复：{{ item.comment_count }}
            </span>

            <div slot="actions" style="margin:0 10px">
              <a-tag v-for="(subitem,subindex) in item.tags" :key="subindex" :color="color[subindex]">
                {{ subitem.name }}
              </a-tag>
            </div>
            <template slot="actions">
              <span>回复</span>
              <span v-if="item.user.id=user.id">编辑</span>
            </template>
            <p slot="content">
              {{ item.content }}
            </p>
            <img v-if="item.user.avatar_url!==''" slot="content" height="120" alt="logo" :src="item.user.avatar_url">
            <a-tooltip slot="datetime" :title="item.created_at">
              <span>{{ item.created_at|dateFormat }}</span>
              <a-divider type="vertical" />
              <span>{{ $s2date(item.created_at).fromNow() }}</span>
            </a-tooltip>
          </a-comment>
        </a-list-item>
      </a-list>
      <a-modal
        v-model="collectVisible"
        title="Title"
        on-ok="handleOk"
      >
        <template slot="footer">
          <a-button key="back" @click="chandleCancel">
            取消
          </a-button>
          <a-button key="submit" type="primary" :loading="loading" @click="handleOk">
            确定
          </a-button>
        </template>
        <a-form-item
          label="标 签"
          required
          :label-col="{span: 4}"
          :wrapper-col="{span: 6}"
        >
          <a-select
            v-model="favorites"
            mode="multiple"
            placeholder="请选择收藏夹"
            style="width: 200px"
          >
            <a-select-option v-for="item in existFavorites" :key="item.id">
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-row>
          <a-col :span="16">
            <a-form-item
              label="新标签"
              :label-col="{span:6}"
              :wrapper-col="{span: 16}"
            >
              <a-input
                v-model="favorite"
              />
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-button style="margin-top: 5px" @click="addFavorite">
              添加
            </a-button>
          </a-col>
        </a-row>
      </a-modal>
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
      user: null,
      top: 20,
      visible: false,
      color: ['pink', 'red', 'orange', 'orange', 'cyan', 'blue', 'purple'],
      loading: false,
      collectVisible: false,
      favorites: [],
      existFavorites: [],
      favorite: ''
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

      this.momentList = momentList
    },
    showDrawer() {
      this.visible = true
    },
    onClose() {
      this.visible = false
    },
    handleCancel(e) {
      this.visible = false
    },
    async star(id, index) {
      this.collectVisible = true
      const res = await this.$axios.$get(`/api/favorite`)
      if (res !== undefined) {
        this.existFavorites = res.data
        this.favorites.push(this.existFavorites[0].id)
      } else this.$message.error('无法获取收藏夹')
    },
    async handleOk(e) {
      this.loading = true

      const params = {
        pageNo: pageNo,
        favorites: this.favorites
      }
      const res = await this.$axios.$post('/api/favorite', favorites)
      if (res.code === 200) this.$message.info('收藏成功')
      this.loading = false
      this.visible = false
    },
    chandleCancel(e) {
      this.collectVisible = false
    },
    like(id, index) {},
    comment(id, index) {},
    addFavorite() {
      const vm = this
      if (this.favorite === '') {
        this.$message.error('标签为空')
        return
      }
      for (const v of this.existFavorites) {
        if (v.name === vm.tag) {
          vm.$message.error('标签重复')
          return
        }
      }
      this.existFavorites.push({ name: this.tag })
      this.favorites.push(this.favorite)
      this.favorite = ''
    }
  }
}
</script>

<style scoped>
</style>
