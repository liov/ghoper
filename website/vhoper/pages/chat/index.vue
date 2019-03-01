<template>
  <a-row>
    <a-col :span="6">
      ss
    </a-col>
    <a-col :span="12">
      <a-list
        class="comment-list"
        :header="`${msgs.length} 条消息`"
        item-layout="horizontal"
        :data-source="msgs"
      >
        <a-list-item slot="renderItem" slot-scope="item">
          <a-comment
            :author="item.send_user.name"
            :avatar="item.send_user.avatar_url"
          >
            <!--   <template slot="actions">
              <span v-for="(action,index) in item.actions" :key="index">{{ action }}</span>
            </template>-->
            <p slot="content">
              {{ item.content }}
            </p>
            <a-tooltip slot="datetime" :title="moment(item.created_at).format('YYYY-MM-DD HH:mm:ss')">
              <span>{{ item.created_at }}</span>
            </a-tooltip>
          </a-comment>
        </a-list-item>
      </a-list>
      <div>
        <a-list
          v-if="comments.length"
          :data-source="comments"
          :header="`${comments.length} ${comments.length > 1 ? 'replies' : 'reply'}`"
          item-layout="horizontal"
        >
          <a-list-item slot="renderItem" slot-scope="item">
            <a-comment
              :author="item.author"
              :avatar="item.avatar"
              :content="item.content"
              :datetime="item.datetime"
            />
          </a-list-item>
        </a-list>
        <a-comment>
          <a-avatar
            slot="avatar"
            :src="user.avatar_url"
            alt="Han Solo"
          />
          <div slot="content">
            <a-form-item>
              <a-textarea :rows="4" :value="value" @change="handleChange" />
            </a-form-item>
            <a-form-item>
              <a-button
                html-type="submit"
                :loading="submitting"
                type="primary"
                @click="handleSubmit"
              >
                发送消息
              </a-button>
            </a-form-item>
          </div>
        </a-comment>
      </div>
    </a-col>
    <a-col :span="6" />
  </a-row>
</template>

<script>
import moment from 'moment'
import axios from 'axios'
export default {
  middleware: 'auth',
  data() {
    return {
      comments: [],
      submitting: false,
      value: '',
      moment,
      user: null,
      ws: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      recipient: null, // Email address used for grabbing an avatar
      joined: false // True if email and username have been filled in
    }
  },
  async asyncData() {
    const res = await axios.get(`/api/chat/getChat`)

    return { msgs: res.data.data !== null ? res.data.data : [] }
  },
  created: function() {
    // 运行在服务端
    this.user = this.$store.state.user
  },
  mounted: function() {
    this.newWs()
    /* this.chatContent=JSON.parse(localStorage.getItem("chatContent"));
          if(this.chatContent === null) this.chatContent=[]; */
    /*    const div = document.querySelector('.chat')
    div.scrollIntoView(false) */
  },
  /*       beforeDestroy(){
            vm.ws.close()
        }, */
  methods: {
    newWs: function() {
      // 不能放在created里
      const vm = this
      this.ws = new WebSocket('ws://' + window.location.host + '/ws/chat')
      this.ws.onopen = function() {
        vm.$message.info('建立websocket连接')
      }
      this.ws.onmessage = function(evt) {
        vm.submitting = false
        vm.msgs = [evt.data, ...vm.msgs]
        vm.value = ''
      }
      this.ws.onclose = function() {
        vm.$message.info('websocket连接关闭')
      }
    },
    send: function() {
      const vm = this
      if (vm.ws.readyState !== 1) {
        this.newWs()
      }
      if (this.newMsg !== '') {
        this.ws.send(
          JSON.stringify({
            recipient_user_id: vm.recipient,
            sender_user_id:
              vm.$store.state.user !== null
                ? vm.$store.state.user.id
                : parseInt(localStorage.getItem('user')),
            content: vm.newMsg // Strip out html
          })
        )
        this.newMsg = '' // Reset newMsg
      }
    },
    delChat: function() {
      localStorage.removeItem('chatContent')
      this.chatContent = []
    },
    handleSubmit() {
      if (!this.value) {
        return
      }
      if (this.ws.readyState !== 1) {
        this.newWs()
      }
      this.submitting = true

      this.ws.send(
        JSON.stringify({
          recipient_user_id: this.recipient,
          sender_user_id:
            this.user !== null
              ? this.user.id
              : parseInt(localStorage.getItem('user')),
          content: this.value // Strip out html
        })
      )
    },
    handleChange(e) {
      this.value = e.target.value
    }
  }
}
</script>

<style scoped>
</style>
