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
      <span style="padding:0 8px" @click="$emit('more',index,comment.id)">展开更多评论</span>
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
    <div class="sub_comment">
      <hoper-comment
        v-for="(subComment,subIndex) in comment.sub_comments"
        :key="subIndex"
        :comment="subComment"
        @reply="$emit('reply',subComment)"
        @del="$emit('del',comment.id)"
        @more="$emit('more',index,comment.id)"
      />
    </div>
  </a-comment>
</template>

<script>
// $emit必须重启 npm run
export default {
  name: 'HoperComment',
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
      this.likes = 1
      this.dislikes = 0
      this.action = 'liked'
    },
    dislike() {
      this.likes = 0
      this.dislikes = 1
      this.action = 'disliked'
    },
    more(show) {
      if (show) document.getElementById('someOne').style.display = 'none'
      else document.getElementById('someOne').style.display = 'inline'
    }
  }
}
</script>

<style scoped>
</style>
