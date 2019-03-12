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
          <a-comment>
            <nuxt-link slot="author" :to="'/user/'+item.send_user.id">
              <span>{{ item.send_user.name }}</span>
            </nuxt-link>
            <nuxt-link slot="avatar" :to="'/user/'+item.send_user.id">
              <a-avatar :src="item.send_user.avatar_url" alt="头像" />
            </nuxt-link>
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
    <a-col :span="2">
      <a-affix :offset-top="this.top">
        <a-button type="primary" @click="()=>{this.top += 10}">
          Affix top
        </a-button>
      </a-affix>
      <a-back-top />
    </a-col>
  </a-row>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      top: 80,
      comments: [],
      submitting: false,
      value: '',
      user: null,
      socket: null, // Our websocket
      newMsg: '', // Holds new messages to be sent to the server
      recipient: 0, // Email address used for grabbing an avatar
      scheme: ''
    }
  },
  async asyncData({ $axios }) {
    const res = await $axios.$get(`/api/chat/getChat`)

    return { msgs: res.data !== null ? res.data : [] }
  },
  created: function() {
    // 运行在服务端
    this.user = this.$store.state.user
  },
  mounted: function() {
    // fetch(`/tpl/iris-ws.js`)
    const vm = this
    const script = document.createElement('script')
    script.type = 'text/javascript' // 设置Type
    script.src = '/tpl/iris-ws.js' // 设置src
    document.head.appendChild(script) // 异步加载
    script.onload = function() {
      vm.scheme = document.location.protocol === 'https:' ? 'wss' : 'ws'
      // see app.Get("/echo", ws.Handler()) on main.go
      vm.newWs()
    }
    setTimeout(function() {
      document.querySelector('#bottom').scrollIntoView()
    }, 0)
    /* this.chatContent=JSON.parse(localStorage.getItem("chatContent"));
              if(this.chatContent === null) this.chatContent=[]; */
  },
  updated: function() {},
  beforeDestroy() {
    this.socket.Disconnect()
  },
  methods: {
    newWs: function() {
      // 不能放在created里
      const vm = this
      const wsURL = this.scheme + '://' + 'hoper.xyz' + '/ws/echo'
      this.socket = new Ws(wsURL)
      this.socket.OnConnect(function() {
        if (vm.value !== '') {
          vm.handleSubmit()
        }
      })

      this.socket.OnDisconnect(function() {
        // console.log('websocket连接关闭')
      })

      // read events from the server
      this.socket.On('chat', function(msg) {
        vm.submitting = false
        vm.msgs = [...vm.msgs, JSON.parse(msg)]
        vm.value = ''
        vm.$nextTick(function() {
          document.querySelector('#bottom').scrollIntoView()
        })
      })

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

      if (this.socket.conn.readyState !== 1) {
        this.newWs()
        return
      }
      this.submitting = true

      this.socket.Emit(
        'chat',
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
