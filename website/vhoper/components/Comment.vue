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
        <span style="padding-left: '8px';cursor: 'auto'">
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
        <span style="padding-left: '8px';cursor: 'auto'">
          {{ dislikes }}
        </span>
      </span>
      <span @click="$emit('reply',comment)">回复</span>
      <span v-if="$store.state.user&&$store.state.user.id === comment.user.id " @click="$emit('del',comment.id)">删除</span>
    </span>
    <p slot="content">
      {{ comment.content }}
    </p>
    <a-tooltip slot="datetime" :title="comment.created_at|dateFormat">
      <span>{{ comment.created_at|dateFormat }}</span>
      <a-divider type="vertical" />
    </a-tooltip>
    <a-tooltip slot="datetime">
      <span>{{ $s2date(comment.created_at).fromNow() }}</span>
    </a-tooltip>
    <a-collapse default-active-key="1" :bordered="false">
      <a-collapse-panel :key="comment.id" header="展开更多评论">
        <hoper-comment v-for="(subComment,index) in comment.sub_comments" :key="index" :comment="subComment" @reply="$emit('reply',subComment)" @del="$emit('del',comment.id)" />
      </a-collapse-panel>
    </a-collapse>
  </a-comment>
</template>

<script>
// $emit必须重启 npm run
export default {
  name: 'HoperComment',
  props: ['comment'],
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
    }
  }
}
</script>

<style scoped>
</style>
