<template>
  <a-row>
    <a-col :span="2" />
    <a-col :span="20">
      <a-comment>
        <nuxt-link slot="author" :to="'/user/' + moment.user.id">
          <span>{{ moment.user.name }}</span>
        </nuxt-link>
        <nuxt-link slot="avatar" :to="'/user/' + moment.user.id">
          <a-avatar :src="moment.user.avatar_url" alt="头像" />
        </nuxt-link>
        <span slot="actions" style="margin-right: 10px" @click="star()">
          <a-icon
            type="star"
            :theme="
              user_action.collect.indexOf(moment.id) > -1
                ? 'twoTone'
                : 'outlined'
            "
            two-tone-color="#eb2f96"
          />
          收藏：{{ moment.collect_count }}
        </span>
        <span slot="actions" style="margin-right: 10px" @click="like()">
          <a-icon
            type="heart"
            :theme="
              user_action.like.indexOf(moment.id) > -1 ? 'twoTone' : 'outlined'
            "
            two-tone-color="#eb2f96"
          />
          喜欢：{{ moment.like_count }}
        </span>
        <span slot="actions" style="margin-right: 10px" @click="approve()">
          <a-icon
            type="like"
            :theme="
              user_action.approve.indexOf(moment.id) > -1
                ? 'twoTone'
                : 'outlined'
            "
          />
          点赞：{{ moment.approve_count }}
        </span>
        <span
          slot="actions"
          style="margin-right: 10px"
          @click="
            showModal({
              user: moment.user,
              user_id: moment.user_id,
              parent_id: 0
            })
          "
        >
          <a-icon
            type="message"
            :theme="moment.id > 0 ? 'twoTone' : 'outlined'"
          />
          评论：{{ moment.comment_count }}
        </span>
        <span slot="actions" style="margin-right: 10px">
          浏览：{{ moment.browse_count }}
        </span>

        <template slot="actions" style="margin:0 10px">
          <a-tag
            v-for="(subitem, subindex) in moment.tags"
            :key="subindex"
            :color="color[subindex]"
          >
            {{ subitem.name }}
          </a-tag>
        </template>

        <template slot="content">
          <div
            style="margin: .5rem .5rem 0 .5rem"
            v-html="$md.render(moment.content)"
          />
        </template>
        <img
          v-for="(subitem, subindex) in image_url"
          :key="subindex"
          slot="content"
          height="120"
          alt="logo"
          :src="subitem"
        >
        <a-tooltip slot="datetime" :title="moment.created_at | dateFormat">
          <span>{{ moment.created_at | dateFormat }}</span>
          <a-divider type="vertical" />
        </a-tooltip>
        <a-tooltip slot="datetime">
          <span>{{ $s2date(moment.created_at).fromNow() }}</span>
        </a-tooltip>
        <nuxt-link
          v-if="moment.user.id === user.id"
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
      <a-modal v-model="collectVisible" title="Title" on-ok="handleOk">
        <template slot="footer">
          <a-button key="back" @click="collectCancel">
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
            <a-select-option v-for="item in existFavorites" :key="item.id">
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
            <a-button style="margin-top: 5px" @click="addFavorite">
              添加
            </a-button>
          </a-col>
        </a-row>
      </a-modal>

      <new-comment
        kind="moment"
        :count="moment.comment_count"
        @reply="showModal"
      />

      <a-modal
        v-model="visible"
        :title="'Reply To: ' + comment.user.name"
        on-ok="commentHandleOk"
      >
        <template slot="footer">
          <a-button key="back" @click="handleCancel">
            返回
          </a-button>
          <a-button
            key="submit"
            type="primary"
            :loading="submitting"
            @click="handleSubmit"
          >
            评论
          </a-button>
        </template>
        <a-comment>
          <a-avatar slot="avatar" :src="user.avatar_url" alt="Han Solo" />
          <div slot="content">
            <a-form-item>
              <a-textarea :rows="4" :value="value" @change="handleChange" />
            </a-form-item>
          </div>
        </a-comment>
      </a-modal>
    </a-col>
    <a-col :span="2" />
  </a-row>
</template>

<script>
import NewComment from '../../components/NewComment'

export default {
  components: { NewComment },
  data() {
    return {
      // md: null,
      user: null,
      image_url: [],
      collectVisible: false,
      existFavorites: [],
      favorites: [],
      favorite: '',
      loading: false,
      visible: false,
      submitting: false,
      value: '',
      comment: { user: { name: '' } }
    }
  },
  async asyncData({ $axios, params, query }) {
    const res = await $axios.$get(
      `/api/moment/${params.id}?index=${query.index}`
    )
    return {
      moment: res.data,
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
    // this.md = require('markdown-it')()
    this.user = this.$store.state.user ? this.$store.state.user : { id: 0 }
    this.image_url =
      this.moment.image_url === '' ? [] : this.moment.image_url.split(',')
  },
  methods: {
    async star() {
      this.collectVisible = true
      if (this.existFavorites.length > 0) {
        this.favorites = [this.existFavorites[0].id]
        return
      }
      const res = await this.$axios.$get(`/api/favorites`)
      if (res !== undefined) {
        this.existFavorites = res.data
        this.favorites = [this.existFavorites[0].id]
      } else this.$message.error('无法获取收藏夹')
    },
    async handleOk(e) {
      this.loading = true
      const params = {
        ref_id: this.moment.id,
        kind: 'moment',
        favorites_ids: this.favorites
      }
      const res = await this.$axios.$post('/api/collection', params)
      if (res.code === 200) {
        this.$message.info('收藏成功')
        this.moment.collect_count += 1
        this.user_action.collect.push(this.moment.id)
      } else this.$message.error(res.msg)
      this.loading = false
      this.collectVisible = false
    },
    collectCancel(e) {
      this.collectVisible = false
    },
    async approve() {
      const idx = this.user_action.approve.indexOf(this.moment.id)
      if (idx > -1) {
        this.user_action.approve.splice(idx, 1)
        this.moment.approve_count -= 1
      }
      const params = {
        ref_id: this.moment.id,
        kind: 'moment'
      }
      const res = await this.$axios.$post('/api/approve', params)
      if (res.code === 200) {
        this.moment.approve_count += 1
        this.user_action.approve.push(this.moment.id)
      }
    },
    async like() {
      const idx = this.user_action.like.indexOf(this.moment.id)
      if (idx > -1) {
        this.user_action.like.splice(idx, 1)
        this.moment.like_count -= 1
      }
      const params = {
        ref_id: this.moment.id,
        kind: 'moment'
      }
      const res = await this.$axios.$post('/api/like', params)
      if (res.code === 200) {
        this.moment.like_count += 1
        // this.setLocalAction('like', id)
        this.user_action.like.push(this.moment.id)
      }
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
      if (res.code === 200) this.$message.info('添加收藏夹成功')
      this.existFavorites.push(res.data)
      this.favorites.push(res.data.id)
      this.favorite = ''
    },
    showModal: function (comment) {
      this.comment = comment
      this.visible = true
    },
    async handleSubmit(e) {
      if (!this.value) {
        return
      }
      this.submitting = true
      const comment = {
        content: this.value,
        recv_user_id: this.comment.user_id,
        parent_id: this.comment.id,
        root_id:
          this.comment.parent_id === 0 ? this.comment.id : this.comment.root_id,
        moment_id: this.moment.id
      }
      const res = await this.$axios.$post(
        `/api/comment/moment/${this.$route.params.id}`,
        comment
      )
      if (res.code === 200) {
        this.$message.info('评论成功')
        this.value = ''
      } else this.$message.error(res.msg)
      this.submitting = false
      this.visible = false
    },
    handleCancel(e) {
      this.visible = false
    },
    handleChange(e) {
      this.value = e.target.value
    }
  }
}
</script>

<style scoped></style>
