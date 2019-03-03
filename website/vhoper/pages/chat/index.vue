<template>
  <a-row>
    <a-col :span="2" style="padding: 24px 0">
      <div style="text-align: center">
        死<br>生<br>契<br>阔<br>，<br>与<br>子<br>成<br>说<br>。
      </div>
      <div style="text-align: center">
        执<br>子<br>之<br>手<br>，<br>与<br>子<br>偕<br>老<br>。
      </div>
    </a-col>
    <a-col :span="20">
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
            <a-tooltip slot="datetime" :title="item.created_at">
              <span>{{ item.created_at|dateFormat }}</span>
            </a-tooltip>
          </a-comment>
        </a-list-item>
      </a-list>

      <div id="bottom">
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
    <a-col :span="2" />
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
    // 这是什么黑科技？？？，本来以为DOM没有渲染完就执行，所以没效果，
    // 加了个定时器，时间一直从500减到0，都一直有效
    setTimeout(function() {
      document.querySelector('#bottom').scrollIntoView()
    }, 0)
    /* this.chatContent=JSON.parse(localStorage.getItem("chatContent"));
            if(this.chatContent === null) this.chatContent=[]; */
  },
  updated: function() {
    this.$nextTick(function() {
      document.querySelector('#bottom').scrollIntoView()
    })
  },
  beforeDestroy() {
    this.ws.close()
  },
  methods: {
    newWs: function() {
      // 不能放在created里
      const vm = this
      this.ws = new WebSocket('ws://' + window.location.host + '/ws/chat')
      this.ws.onopen = function() {
        console.log('建立websocket连接')
        if (vm.value !== '') {
          vm.handleSubmit()
        }
      }
      this.ws.onmessage = function(evt) {
        vm.submitting = false
        vm.msgs = [...vm.msgs, JSON.parse(evt.data)]
        vm.value = ''
      }
      this.ws.onerror = function() {
        vm.newWs()
      }
      this.ws.onclose = function() {
        console.log('websocket连接关闭')
      }

      document.scrollingElement.scrollTop =
        document.scrollingElement.scrollHeight
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
        return
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
