<template>
  <a-row>
    <a-col :span="3">
      <a-form-item
        label=""
        :label-col="{ span: 4, offset: 2 }"
        :wrapper-col="{ span: 24, offset: 2 }"
      >
        <nuxt-link to="/moment/add">
          <a-button icon="edit">
            添加
          </a-button>
        </nuxt-link>
        <a-affix :offset-top="top">
          <a-button
            type="primary"
            @click="showDrawer"
          >
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
        <a-list-item
          slot="renderItem"
          slot-scope="item, index"
          style="padding-bottom:0 "
        >
          <a-comment>
            <nuxt-link
              slot="author"
              :to="'/user/' + item.user.id"
            >
              <span>{{ item.user.name }}</span>
            </nuxt-link>
            <nuxt-link
              slot="avatar"
              :to="'/user/' + item.user.id"
            >
              <a-avatar
                :src="item.user.avatar_url"
                alt="头像"
              />
            </nuxt-link>
            <span
              slot="actions"
              style="margin-right: 10px"
              @click="star(item.id, index)"
            >
              <a-icon
                type="star"
                :theme="
                  user_action.collect.indexOf(item.id) > -1
                    ? 'twoTone'
                    : 'outlined'
                "
                two-tone-color="#eb2f96"
              />
              收藏：{{ item.collect_count }}
            </span>
            <span
              slot="actions"
              style="margin-right: 10px"
              @click="like(item.id, index)"
            >
              <a-icon
                type="heart"
                :theme="
                  user_action.like.indexOf(item.id) > -1
                    ? 'twoTone'
                    : 'outlined'
                "
                two-tone-color="#eb2f96"
              />
              喜欢：{{ item.like_count }}
            </span>
            <span
              slot="actions"
              style="margin-right: 10px"
              @click="approve(item.id, index)"
            >
              <a-icon
                type="like"
                :theme="
                  user_action.approve.indexOf(item.id) > -1
                    ? 'twoTone'
                    : 'outlined'
                "
              />
              点赞：{{ item.approve_count }}
            </span>
            <span
              slot="actions"
              style="margin-right: 10px"
              @click="comment(item.id, index)"
            >
              <a-icon
                type="message"
                :theme="item.id > 0 ? 'twoTone' : 'outlined'"
              />
              评论：{{ item.comment_count }}
            </span>
            <span
              slot="actions"
              style="margin-right: 10px"
            >
              浏览：{{ item.browse_count }}
            </span>

            <template
              slot="actions"
              style="margin:0 10px"
            >
              <a-tag
                v-for="(subitem, subindex) in item.tags"
                :key="subindex"
                :color="color[subindex]"
              >
                {{ subitem.name }}
              </a-tag>
            </template>

            <template slot="content">
              <div
                style="margin: 1rem 1rem 0 1rem"
                v-html="$md.render(item.content)"
              />
            </template>
            <img
              v-for="(subitem, subindex) in image_url[index]"
              :key="subindex"
              slot="content"
              height="120"
              alt="logo"
              :src="subitem"
            >
            <a-tooltip
              slot="datetime"
              :title="item.created_at | dateFormat"
            >
              <span>{{ item.created_at | dateFormat }}</span>
              <a-divider type="vertical" />
            </a-tooltip>
            <a-tooltip slot="datetime">
              <span>{{ $s2date(item.created_at).fromNow() }}</span>
            </a-tooltip>
            <a-tooltip
              v-if="item.mood.name !== ''"
              slot="datetime"
            >
              <a-divider type="vertical" />
              <span>心情：{{ item.mood.name }}</span>
            </a-tooltip>
            <nuxt-link
              v-if="item.user.id === user.id"
              slot="datetime"
              title="点击编辑"
              to="/moment/edit"
              style="color: #ccc"
            >
              <a-divider type="vertical" />
              <a-icon type="edit" />
              <span>编辑</span>
            </nuxt-link>
          </a-comment>
        </a-list-item>
      </a-list>
      <a-modal
        v-model="collectVisible"
        title="Title"
        on-ok="handleOk"
      >
        <template slot="footer">
          <a-button
            key="back"
            @click="collectCancel"
          >
            取消
          </a-button>
          <a-button
            key="submit"
            type="primary"
            :loading="loading"
            @click="handleOk"
          >
            确定
          </a-button>
        </template>
        <a-form-item
          label="标 签"
          required
          :label-col="{ span: 4 }"
          :wrapper-col="{ span: 6 }"
        >
          <a-select
            v-model="favorites"
            mode="multiple"
            placeholder="请选择收藏夹"
            style="width: 200px"
          >
            <a-select-option
              v-for="item in existFavorites"
              :key="item.id"
            >
              {{ item.name }}
            </a-select-option>
          </a-select>
        </a-form-item>

        <a-row>
          <a-col :span="16">
            <a-form-item
              label="新建收藏夹"
              :label-col="{ span: 6 }"
              :wrapper-col="{ span: 16 }"
            >
              <a-input v-model="favorite" />
            </a-form-item>
          </a-col>
          <a-col :span="8">
            <a-button
              style="margin-top: 5px"
              @click="addFavorite"
            >
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
        <template
          slot="buildOptionText"
          slot-scope="props"
        >
          <span v-if="props.value !== '50'">{{ props.value }}条/页</span>
          <span v-if="props.value === '50'">全部</span>
        </template>
      </a-pagination>
    </a-col>
  </a-row>
</template>

<script>
import Jsonp from '../../plugins/utils/jsonp'
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
      favorite: '',
      image_url: [],
      ref_id: 0,
      tmpIdx: 0,
      starIds: []
    }
  },
  computed: {},
  watch: {
    current: async function () {
      await this.next(this.current - 1)
    }
  },
  async asyncData({ $axios }) {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await $axios.$get(`/api/moment`, { params })
    console.log(res);
    return {
      momentList: res.data,
      total: res.count,
      topCount: res.top_count,
      user_action:
        res.user_action != null
          ? res.user_action
          : {
            collect: [],
            like: [],
            approve: [],
            comment: [],
            browse: []
          }
    }
  },
  created: function () {
    this.user = this.$store.state.user ? this.$store.state.user : { id: 0 }
    for (const i in this.momentList) {
      this.image_url.push(
        this.momentList[i].image_url === ''
          ? []
          : this.momentList[i].image_url.split(',')
      )
    }
  },
  mounted: function () {},
  methods: {
    async onShowSizeChange(current, pageSize) {
      this.pageSize = pageSize
      await this.next(current - 1)
    },
    setMoment: function (moment) {
      localStorage.setItem('moment_' + moment.id, moment)
    },
    next: async function (pageNo) {
      const params = {
        pageNo: pageNo,
        pageSize: this.pageSize
      }

      // 这里可以这么写，async，await函数，或者 return axios().then((res)=>{})返回Promise
      const res = await this.$axios.$get(`/api/moment`, { params })

      this.momentList = res.data
    },
    showDrawer() {
      this.visible = true
      Jsonp.jsonp('https://hoper.xyz/api/jsonp?callback=$nuxt.$message.info')
    },
    onClose() {
      this.visible = false
    },
    handleCancel(e) {
      this.visible = false
    },
    async star(id, index) {
      this.ref_id = id
      this.tmpIdx = index
      this.collectVisible = true
      if (this.existFavorites.length > 0) {
        this.favorites = [this.existFavorites[0].id]
        return
      }
      const res = await this.$axios.$get(`/api/favorites`)
      if (res !== undefined) {
        this.existFavorites = res.data
        this.favorites = [this.existFavorites[0].id]
      } else {
        this.$message.error('无法获取收藏夹')
      }
    },
    async handleOk(e) {
      this.loading = true
      const params = {
        ref_id: this.ref_id,
        kind: 'moment',
        favorites_ids: this.favorites
      }
      const res = await this.$axios.$post('/api/collection', params)
      if (res.code === 200) {
        this.$message.info('收藏成功')
        this.momentList[this.tmpIdx].collect_count += 1
        this.user_action.collect.push(this.ref_id)
        // this.setLocalAction('moment', this.ref_id)
      } else {
        this.$message.error(res.msg)
      }
      this.loading = false
      this.collectVisible = false
    },
    collectCancel(e) {
      this.collectVisible = false
    },
    async approve(id, index) {
      const idx = this.user_action.approve.indexOf(id)
      if (idx > -1) {
        this.user_action.approve.splice(idx, 1)
        this.momentList[index].approve_count -= 1
      }
      const params = {
        ref_id: id,
        kind: 'moment'
      }
      const res = await this.$axios.$post('/api/approve', params)
      if (res.code === 200) {
        this.momentList[index].approve_count += 1
        this.user_action.approve.push(id)
      }
    },
    async like(id, index) {
      const idx = this.user_action.like.indexOf(id)
      if (idx > -1) {
        this.user_action.like.splice(idx, 1)
        this.momentList[index].like_count -= 1
      }
      const params = {
        ref_id: id,
        kind: 'moment'
      }
      const res = await this.$axios.$post('/api/like', params)
      if (res.code === 200) {
        this.momentList[index].like_count += 1
        // this.setLocalAction('like', id)
        this.user_action.like.push(id)
      }
    },
    comment(id, index) {
      this.$router.push(
        `/moment/${id}?index=${(this.current - 1) * this.pageSize + index}`
      )
    },
    async addFavorite() {
      if (this.favorite === '') {
        this.$message.error('标签为空')
        return
      }
      for (const v of this.existFavorites) {
        if (v.name === this.favorite) {
          this.$message.error('标签重复')
          return
        }
      }
      const res = await this.$axios.$post('/api/favorites', {
        name: this.favorite
      })
      if (res.code === 200) {
        this.$message.info('添加收藏夹成功')
      }
      this.existFavorites.push(res.data)
      this.favorites.push(res.data.id)
      this.favorite = ''
    },
    setLocalAction(action, id) {
      const key = 'moment_' + action + '_' + this.user.id
      localStorage.setItem(
        key,
        localStorage.getItem(key) ? localStorage.getItem(key) + ',' + id : id
      )
    },
    getLocalAction() {
      for (const v of localStorage
        .getItem('moment_collect_' + this.user.id)
        .split(',')) {
        this.user_action.collect.push(parseInt(v))
      }
    }
  }
}
</script>

<style scoped></style>
