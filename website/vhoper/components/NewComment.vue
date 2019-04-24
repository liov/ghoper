<template>
  <a-list
    class="comment-list"
    :header="`${count} 条评论`"
    item-layout="horizontal"
    :data-source="comments"
  >
    <div slot="loadMore" :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }">
      <a-spin v-if="loading" id="loading" />
      <a-button v-else>
        已经到底
      </a-button>
    </div>
    <a-list-item slot="renderItem" slot-scope="item,index">
      <a-comment
        :author="item.user.name"
        :avatar="item.user.avatar_url"
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
          <span style="padding-right: 8px" @click="$emit('reply',item)">回复</span>
          <span v-if="$store.state.user&&$store.state.user.id === item.user.id " style="padding:0 8px" @click="delComment(item.id)">删除</span>
        </span>
        <template slot="content">
          <div>
            {{ item.content }}
          </div>
        </template>
        <a-tooltip slot="datetime" :title="item.created_at|dateFormat">
          <span>{{ item.created_at|dateFormat }}</span>
          <a-divider type="vertical" />
        </a-tooltip>
        <a-tooltip slot="datetime">
          <span>{{ $s2date(item.created_at).fromNow() }}</span>
        </a-tooltip>
        <div v-if="item.sub_comments.length>0" class="sub-comment">
          <a-comment
            v-for="(subComment,subIndex) in item.sub_comments"
            :key="subIndex"
            :author="subComment.user.name"
            :avatar="subComment.user.avatar_url"
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
              <span style="padding-right: 8px" @click="$emit('reply',subComment)">回复</span>
              <span v-if="$store.state.user&&$store.state.user.id === subComment.user.id " style="padding:0 8px" @click="delComment(subComment.id)">删除</span>
            </span>
            <template slot="content">
              <div>
                {{ subComment.content }}
              </div>
            </template>
            <span v-if="subComment.parent_id !== subComment.root_id" slot="datetime" :title="subComment.user.name">
              <span>@<nuxt-link :to="'/user/'+subComment.recv_user.id">{{ subComment.user.name }}</nuxt-link></span>
              <a-divider type="vertical" />
            </span>
            <a-tooltip slot="datetime" :title="subComment.created_at|dateFormat">
              <span>{{ subComment.created_at|dateFormat }}</span>
              <a-divider type="vertical" />
            </a-tooltip>
            <a-tooltip slot="datetime">
              <span>{{ $s2date(subComment.created_at).fromNow() }}</span>
            </a-tooltip>
          </a-comment>
        </div>
        <a-button @click="onLoadMore($event,index)">
          展开更多评论
        </a-button>
        <a-button @click="hide">
          收起评论
        </a-button>
      </a-comment>
    </a-list-item>
  </a-list>
</template>

<script>
// $emit必须重启 npm run

export default {
  name: 'NewComment',
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
      comments: []
    }
  },
  mounted() {
    const vm = this
    const ScrollReveal = require('scrollreveal')
    const ScrollMagic = require('scrollmagic')
    ScrollReveal.default().reveal('.sub-comment')
    const controller = new ScrollMagic.Controller()
    // build scene
    const scene = new ScrollMagic.Scene({
      triggerElement: '#loading',
      triggerHook: 'onEnter'
    })
      .addTo(controller)
      .on('enter', function(e) {
        // simulate ajax call to add content using the function below
        vm.onLoadMore()
      })
  },
  methods: {
    like() {
      this.likes += 1
      this.action = 'liked'
    },
    dislike() {
      this.dislikes += 1
      this.action = 'disliked'
    },
    hide(e) {
      e.target.previousElementSibling.previousElementSibling.style.display =
        'none'
    },
    delComment(id) {
      console.log(id)
    },
    async onLoadMore(e, index) {
      console.log(1)
      if (index) {
        if (e.target.previousElementSibling.style.display === 'none') {
          e.target.previousElementSibling.style.display = 'block'
          return
        }
        if (
          this.comments[index].sub_comments !== null &&
          this.comments[index].sub_comments.length ===
            this.comments[index].comment_count
        ) {
          this.$message.info('没有更多评论')
          return
        }

        const commentRes = await this.$axios.$get(
          `/api/comments/${this.$props.kind}/${this.$route.params.id}?offset=${
            this.comments[index].sub_comments
              ? this.comments[index].sub_comments.length
              : 0
          }&limit=5&rootId=${this.comments[index].id}`
        )
        if (commentRes.code === 200) {
          if (commentRes.data.length === 0) {
            this.loading = false
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
          }&limit=5&rootId=0`
        )
        if (commentRes.code === 200) {
          if (commentRes.data.length === 0) {
            this.loading = false
            return
          }
          this.comments = this.comments.concat(commentRes.data)
        }
      }
    }
  }
}
</script>

<style scoped>
.sub-comment {
  overflow: auto;
  width: 971px;
  max-height: 300px;
}
</style>
