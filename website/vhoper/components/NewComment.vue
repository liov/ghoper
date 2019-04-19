<template>
  <a-comment
    :author="comment.user.name"
    :avatar="comment.user.avatar_url"
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
      <span style="padding-right: 8px" @click="$emit('reply',comment)">回复</span>
      <span v-if="$store.state.user&&$store.state.user.id === comment.user.id " style="padding:0 8px" @click="$emit('del',comment.id)">删除</span>
    </span>
    <template slot="content">
      <div>
        {{ comment.content }}
      </div>
    </template>
    <a-tooltip slot="datetime" :title="comment.created_at|dateFormat">
      <span>{{ comment.created_at|dateFormat }}</span>
      <a-divider type="vertical" />
    </a-tooltip>
    <a-tooltip slot="datetime">
      <span>{{ $s2date(comment.created_at).fromNow() }}</span>
    </a-tooltip>
    <div v-for="(subComment,subIndex) in comment.sub_comments" :key="subIndex" class="sub_comment">
      <a-comment
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
          <span v-if="$store.state.user&&$store.state.user.id === subComment.user.id " style="padding:0 8px" @click="$emit('del',subComment.id)">删除</span>
        </span>
        <template slot="content">
          <div>
            {{ subComment.content }}
          </div>
        </template>
        <span v-if="subComment.parent_id !== 0" slot="datetime" :title="subComment.user.name">
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
    <span style="padding:0 8px;cursor:pointer" @click="$emit('more',index,comment.id)">展开更多评论</span>
    <span style="padding:0 8px;cursor:pointer" @click="$emit('more',index,comment.id)">收起评论</span>
  </a-comment>
</template>

<script>
// $emit必须重启 npm run
export default {
  name: 'NewComment',
  props: ['comment', 'index'],
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
    }
  }
}
</script>

<style scoped>
</style>
