<template>
    <div class="hoper">
        <div class="display">
            <div class="chat">
                <div class="row" v-for="item of Msg">
                    <p>
                        <span class="name">{{item.send_user.name}}</span>:
                        <span class="content">{{item.content}}</span>
                    </p>
                     <p>
                            <span class="time">{{item.created_at|dateFormat}}</span>
                       <span class="device"> 来自：{{item.device}}</span>
                    </p>
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
            }
        }
    }
</script>

<style lang="scss" scoped>
    .display{
        position: fixed;
        top: .2rem;
        left: .2rem;
        width: 10rem;
        height: 13rem;
        overflow: auto;
    }
    .row{
        width: 9rem;
        height: 1.2rem;
        font-size: .3rem;
        background-color: aliceblue;
        color: #000;
        line-height: .8rem;
        padding-left: .5rem;
        margin-top: .2rem;
        p{
            height: .6rem;
            width: 6.6rem;
        }
        .name{
            color: cornflowerblue;
        }
        .content{
            color: blueviolet;
            margin-left: .1rem;
        }
        .time{
            color: #000;
            font-size: .2rem;
        }
        .device{
            right: 1.5rem;
            position: absolute;
        }
    }
    .inp{
        position: fixed;
        bottom: 0;
        width: 100%;
    }
    .del{
        position: fixed;
        bottom: 7%;
        left: 36%;
        width: 27%;
        height: 50px;
        font-size: .3rem;
        background-color: #2baee9;
        line-height: 50px;
        text-align: center;
    }
</style>
