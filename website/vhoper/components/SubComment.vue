<template>
  <a-list item-layout="horizontal" :data-source="subComments">
    <div
      slot="loadMore"
      :style="{
        textAlign: 'center',
        marginTop: '12px',
        height: '32px',
        lineHeight: '32px'
      }"
    >
      <a-spin v-if="loading" :id="'loading' + index" />
      <span v-else>已经到底</span>
    </div>
    <a-list-item slot="renderItem" slot-scope="subComment, subIndex">
      <a-comment
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
                @click="$emit('like', subComment.id)"
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
                @click="$emit('dislike', subComment.id)"
              />
            </a-tooltip>
            <span style="padding:0 8px;cursor: auto">
              {{ dislikes }}
            </span>
          </span>
          <span
            style="padding-right: 8px"
            @click="$emit('reply', subComment)"
          >回复</span>
          <span
            v-if="
              $store.state.user && $store.state.user.id === subComment.user.id
            "
            style="padding:0 8px"
            @click="$emit('delComment', subComment.id)"
          >删除</span>
        </span>
        <template slot="content">
          <div>
            {{ subComment.content }}
          </div>
        </template>
        <span
          v-if="subComment.parent_id !== 0"
          slot="datetime"
          :title="subComment.user.name"
        >
          <span>@<nuxt-link :to="'/user/' + subComment.recv_user.id">{{
            subComment.user.name
          }}</nuxt-link></span>
          <a-divider type="vertical" />
        </span>
        <a-tooltip slot="datetime" :title="subComment.created_at | dateFormat">
          <span>{{ subComment.created_at | dateFormat }}</span>
          <a-divider type="vertical" />
        </a-tooltip>
        <a-tooltip slot="datetime">
          <span>{{ $s2date(subComment.created_at).fromNow() }}</span>
        </a-tooltip>
      </a-comment>
    </a-list-item>
  </a-list>
</template>

<script>
export default {
  name: 'SubComment',
  props: ['subComments', 'index', 'controller'],
  data() {
    return {
      likes: 0,
      dislikes: 0,
      action: null,
      loading: true,
      showLoading: true,
      scene: null
    }
  },
  mounted() {
    const vm = this

    this.$nextTick(() => {
      // build scene
      this.scene = new ScrollMagic.Scene({
        triggerElement: '#loading' + vm.$props.index,
        triggerHook: 'onEnter'
      })
        .addTo(this.$props.controller)
        .on('enter', function (e) {
          // simulate ajax call to add content using the function below
          vm.$emit('more', vm.$props.index)
        })
    })
  }
}
</script>

<style scoped></style>
