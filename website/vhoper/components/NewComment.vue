<template>
  <a-list
    class="comment-list"
    :header="`${count} 条评论`"
    item-layout="horizontal"
    :data-source="comments"
  >
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
        <div class="sub_comment">
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
        <span style="padding:0 8px;font-size:12px;cursor:pointer" @click="more($event,index)">展开更多评论</span>
        <span style="padding:0 8px;font-size:12px;cursor:pointer" @click="hide">收起评论</span>
      </a-comment>
    </a-list-item>
  </a-list>
</template>

<script>
// $emit必须重启 npm run
export default {
  name: 'NewComment',
  props: ['comments', 'kind', 'count'],
  data() {
    return {
      likes: 0,
      dislikes: 0,
      action: null
    }
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
    async more(e, index) {
      if (e.target.previousElementSibling.style.display === 'none') {
        e.target.previousElementSibling.style.display = 'block'
      } else {
        if (
          this.$props.comments[index].sub_comments !== null &&
          this.$props.comments[index].sub_comments.length ===
            this.$props.comments[index].comment_count
        ) {
          this.$message.info('没有更多评论')
          return
        }

        const commentRes = await this.$axios.$get(
          `/api/comments/${this.$props.kind}/${this.$route.params.id}?offset=${
            this.$props.comments[index].sub_comments
              ? this.$props.comments[index].sub_comments.length
              : 0
          }&limit=5&rootId=${this.$props.comments[index].id}`
        )
        if (commentRes.code === 200) {
          this.$props.comments[index].sub_comments = this.$props.comments[index]
            .sub_comments
            ? this.$props.comments[index].sub_comments.concat(commentRes.data)
            : (this.$props.comments[index].sub_comments = commentRes.data)
          this.$props.comments[index].comment_count = commentRes.comment_count
        }
      }
    },
    hide(e) {
      e.target.previousElementSibling.previousElementSibling.style.display =
        'none'
    },
    delComment(id) {
      console.log(id)
    }
  }
}
</script>

<style scoped>
</style>
