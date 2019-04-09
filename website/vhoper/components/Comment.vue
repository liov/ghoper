<template>
  <a-comment
    :author="comment.user.name"
    :avatar="comment.user.avatar_url"
  >
    <span slot="actions">
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
    <hoper-comment v-for="(subComment,index) in comment.sub_comments" :key="index" :comment="subComment" @reply="$emit('reply',subComment)" @del="$emit('del',comment.id)" />
  </a-comment>
</template>

<script>
// $emit必须重启 npm run
export default {
  name: 'HoperComment',
  props: ['comment'],
  methods: {}
}
</script>

<style scoped>
</style>
