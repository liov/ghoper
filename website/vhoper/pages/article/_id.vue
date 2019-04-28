<template>
  <a-row>
    <a-col :span="3" />
    <a-col :span="18">
      <div v-html="article.html_content" />
      <div style="margin: 0 auto">
        <span @click="star(item.id)">
          <a-icon type="star-o" style="margin-right: 8px" />
          {{ article.collect_count }}
        </span>
        <a-divider type="vertical" />
        <span @click="like(item.id)">
          <a-icon type="like-o" style="margin-right: 8px" />
          {{ article.like_count }}
        </span>
        <a-divider type="vertical" />
        <span @click="comment(item.id)">
          <a-icon type="message" style="margin-right: 8px" />
          {{ article.comment_count }}
        </span>
        <a-divider type="vertical" />
        <a-button-group>
          <a-button v-for="(item, index) in article.categories" :key="index">
            {{ item.name }}
          </a-button>
        </a-button-group>
        <a-divider type="vertical" />
        <a-tag
          v-for="(item, index) in article.tags"
          :key="index"
          :color="color[index]"
        >
          {{ item.name }}
        </a-tag>
        <a-divider type="vertical" />
        <a-rate :default-value="2.5" allow-half />
      </div>
    </a-col>
    <a-col :span="3" />
  </a-row>
</template>

<script>
export default {
  data() {
    return {
      color: ['pink', 'red', 'orange', 'orange', 'cyan', 'blue', 'purple']
    }
  },
  async asyncData({ $axios, route }) {
    const res = await $axios.$get(`/api` + route.path)

    return { article: res.data !== null ? res.data : null }
  }
}
</script>

<style scoped></style>
