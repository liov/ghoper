<template>
    <div class="hoper">
        <van-nav-bar
                left-text="返回"
                left-arrow
                fixed
                @click-left="onClickLeft"
                @click-right="onClickRight"
        >
            <span slot="title">会话</span>
            <van-icon name="home" size=".5rem" slot="right" />
        </van-nav-bar>
        <div class="display">
            <div class="chat">
                <div class="row" v-for="item of Msg">
                    <van-row>
                        <van-col span="4"><van-cell class="name van-hairline--bottom">{{item.send_user.name}}</van-cell></van-col>
                        <van-col span="20" class="van-hairline--bottom"><van-field
                                v-model="item.content"
                                type="textarea"
                                rows="1"
                                disabled
                                autosize
                        /></van-col>
                    </van-row>
                    <van-row>
                        <van-col span="16"><van-cell>{{item.created_at|dateFormat}}</van-cell></van-col>
                        <van-col span="8"><van-cell>来自：{{item.device}}</van-cell></van-col>
                    </van-row>
                </div>
            </div>
        </div>
            <div class="inp">
                <van-cell-group>
                    <van-field
                            v-model="newMsg"
                            center
                            clearable
                            label="消息"
                            placeholder="请输入消息"
                            @keyup.enter="send"
                    >
                        <van-button slot="button" size="small" type="primary" @click="send">发送</van-button>
                    </van-field>
                </van-cell-group>
            </div>


    </div>
</template>

<script>
    import axios from "axios"
    export default {
        middleware: 'auth',
        data() {
            return{
                ws: null, // Our websocket
                newMsg: '', // Holds new messages to be sent to the server
                recipient: null, // Email address used for grabbing an avatar
                joined: false // True if email and username have been filled in
            }
        },
        async asyncData() {

            let res = await axios.get(`/api/chat/getChat`);

            return {Msg: res.data.data!==null?res.data.data:[]}
        },
        created: function() {
            //运行在服务端
        },
        mounted:function(){
          this.newWs();
         /* this.chatContent=JSON.parse(localStorage.getItem("chatContent"));
          if(this.chatContent === null) this.chatContent=[];*/
            let div = document.querySelector('.chat');
            div.scrollIntoView(false);
        },
 /*       beforeDestroy(){
            vm.ws.close()
        },*/
        methods: {
            newWs:function(){
                //不能放在created里
                let vm = this;
                vm.ws = new WebSocket('ws://'+window.location.host+'/api/chat/ws');
                vm.ws.addEventListener('message', function(e) {
                    let msg = JSON.parse(e.data);
                    if(msg.content!==undefined) {
                        vm.Msg.push(msg);
                        vm.$nextTick(function(){
                            let div = document.querySelector('.chat');
                            div.scrollIntoView(false)
                        })
                    }
                    //localStorage.setItem("chatContent",JSON.stringify(vm.chatContent))
                });
            },
            send: function () {
                let vm = this;
                if(vm.ws.readyState!==1){
                    this.newWs()
                }
                if (this.newMsg !== '') {
                    this.ws.send(
                        JSON.stringify({
                                recipient_user_id: vm.recipient,
                                sender_user_id: vm.$store.state.user!==null?vm.$store.state.user.id:parseInt(localStorage.getItem("user")),
                                content: vm.newMsg // Strip out html
                            }
                        ));
                    this.newMsg = ''; // Reset newMsg
                }
            },
            delChat:function(){
                localStorage.removeItem("chatContent");
                this.chatContent = []
            },
            onClickLeft() {
                this.$router.push('/')
            },
            onClickRight() {
                this.$router.push('/')
            }
        }
    }
</script>

<style lang="scss" scoped>

    .display{
        position: fixed;
        top: .2rem;
        width: 10rem;
        height: 13rem;
        overflow: auto;
        .chat{
            margin-top: 1rem;
            background-color: #c27c88;
            .row{
                margin-top: .2rem;
                background-color: #fff;
                .name{
                    padding: 10px 10px;
                    color: #2aa198;
                }
            }
        }
    }

    .inp{
        position: fixed;
        bottom: 0;
        width: 100%;
    }

</style>
