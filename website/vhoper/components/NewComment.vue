<template>
  <a-list
    class="comment-list"
    :header="`${count} 条评论`"
    item-layout="horizontal"
    :data-source="comments"
    :locale="{ emptyText: '加载评论中' }"
    style="word-break: break-all;"
  >
    <div
      slot="loadMore"
      :style="{
        textAlign: 'center',
        marginTop: '12px',
        height: '32px',
        lineHeight: '32px'
      }"
    >
      <a-spin
        v-if="loading"
        id="loading"
      />
      <span v-else>
        已经到底
      </span>
    </div>
    <a-list-item
      slot="renderItem"
      slot-scope="item, index"
    >
      <!--      <img
        v-for="(src,sidx) in item"
        :key="sidx"
        slot="extra"
        width="186"
        :src="src.user.avatar_url"
        style="margin-right: 10px"
      >-->

      <a-row style="width: 1000px">
        <a-col
          :span="2"
          style="margin: 16px 0"
        >
          <a-avatar
            shape="square"
            :size="80"
            :src="item.user.avatar_url"
          />
        </a-col>
        <a-col :span="21">
          <a-comment
            :key="index"
            :author="item.user.name"
          >
            <span slot="actions">
              <span>
                <a-tooltip title="Like">
                  <a-icon
                    type="like"
                    :theme="action === 'liked' ? 'filled' : 'outlined'"
                    @click="like"
                  />
                </a-tooltip>
                <span style="padding:0 8px;cursor: auto">
                  {{ likes }}
                </span>
              </span>
              <span>
                <a-tooltip title="Dislike">
                  <a-icon
                    type="dislike"
                    :theme="action === 'disliked' ? 'filled' : 'outlined'"
                    @click="dislike"
                  />
                </a-tooltip>
                <span style="padding:0 8px;cursor: auto">
                  {{ dislikes }}
                </span>
              </span>
              <span
                style="padding-right: 8px"
                @click="reply(item)"
              >回复</span>
              <span
                v-if="
                  $store.state.user && $store.state.user.id === item.user.id
                "
                style="padding:0 8px"
                @click="delComment(item.id)"
              >删除</span>
            </span>
            <template slot="content">
              <div v-html="$md.render(item.content)" />
            </template>
            <span
              v-if="item.parent_id !== 0"
              slot="datetime"
              :title="item.user.name"
            >
              <span>@<nuxt-link :to="'/user/' + item.recv_user.id">{{
                item.user.name
              }}</nuxt-link></span>
              <a-divider type="vertical" />
            </span>
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
            <a-collapse
              v-if="item.sub_comments && item.sub_comments.length > 0"
              key="1"
              default-active-key="1"
              :bordered="false"
            >
              <a-collapse-panel
                key="1"
                header="收起评论"
              >
                <div class="sub-comments">
                  <sub-comment
                    :ref="'subComment' + index"
                    :sub-comments="item.sub_comments"
                    :index="index"
                    :controller="controller"
                    @reply="reply"
                    @more="onLoadMore"
                    @like="like"
                    @dislike="dislike"
                    @delComment="delComment"
                  />
                </div>
              </a-collapse-panel>
            </a-collapse>
          </a-comment>
        </a-col>
        <a-col
          :span="1"
          style="margin: 16px 0"
        >
          {{ index + 1 }}楼
        </a-col>
      </a-row>
    </a-list-item>
  </a-list>
</template>

<script>
// $emit必须重启 npm run

import SubComment from './SubComment'
export default {
  name: 'NewComment',
  components: { SubComment },
  // directives: { infiniteScroll },
  // components: { 'virtual-scroller': VirtualScroller },
  props: ['kind', 'count'],
  data() {
    return {
      likes: 0,
      dislikes: 0,
      action: null,
      loading: true,
      showLoading: true,
      scene: null,
      controller: null,
      comments: []
    }
  },
  mounted() {
    const vm = this
    const ScrollMagic = require('scrollmagic')
    this.controller = new ScrollMagic.Controller()
    this.$nextTick(() => {
      // build scene
      this.scene = new ScrollMagic.Scene({
        triggerElement: '#loading',
        triggerHook: 'onEnter'
      })
        .addTo(this.controller)
        .on('enter', function (e) {
          // simulate ajax call to add content using the function below
          vm.onLoadMore()
        })
    })
  },
  methods: {
    like(id) {
      this.likes += 1
      this.action = 'liked'
      console.log(id)
    },
    dislike(id) {
      this.dislikes += 1
      this.action = 'disliked'
      console.log(id)
    },

    delComment(id) {
      console.log(id)
    },
    reply(comment) {
      this.$emit('reply', comment)
    },
    async onLoadMore(index = -1) {
      // js中0是false
      if (index !== -1) {
        const commentRes = await this.$axios.$get(
          `/api/comments/${this.$props.kind}/${this.$route.params.id}?offset=${
            this.comments[index].sub_comments
              ? this.comments[index].sub_comments.length
              : 0
          }&limit=10&rootId=${this.comments[index].id}`
        )
        if (commentRes.code === 200) {
          if (commentRes.data.length === 0) {
            this.$refs['subComment' + index].scene.destroy()
            this.$refs['subComment' + index].loading = false
            return
          }

          this.comments[index].sub_comments = this.comments[index].sub_comments
            ? this.comments[index].sub_comments.concat(commentRes.data)
            : (this.comments[index].sub_comments = commentRes.data)
          this.comments[index].comment_count = commentRes.comment_count
        }
      } else {
        const commentRes = await this.$axios.$get(
          `/api/comments/${this.$props.kind}/${this.$route.params.id}?offset=${
            this.comments.length
          }&limit=10`
        )
        if (commentRes.code === 200) {
          this.comments = this.comments.concat(commentRes.data)
          if (commentRes.data.length < 10) {
            this.scene.destroy()
            this.loading = false
          }
        }
      }
    }
  }
}
</script>

<style scoped>
.sub-comments {
  overflow: auto;
  min-width: 1000px;
  max-height: 500px;
  /*word-wrap: break-word;*/
  word-break: break-all;
}
</style>
