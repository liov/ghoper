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
                <div class="input-field col s8">
                    <input type="text" v-model="newMsg" @keyup.enter="send">
                </div>
                <div class="input-field col s4">
                    <button class="waves-effect waves-light btn" @click="send">
                        <i class="material-icons right">消息</i>
                        发送
                    </button>
                </div>

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
        width: 8rem;
        height: 11rem;
        overflow: auto;
    }
    .row{
        width: 7rem;
        height: 1rem;
        font-size: .3rem;
        background-color: aliceblue;
        color: #000;
        line-height: .8rem;
        padding-left: .5rem;
        p{
            height: .1rem;
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
        left: 0;
        width: 8rem;
        height: 1rem;
        .input-field{
            float: left;
            margin-left: .2rem;
        }
        input{
            width: 5rem;
            height: .6rem;
            font-size: .3rem;
            line-height: .5rem;
            padding-left: .2rem;
            z-index: 1;
            border: 1px solid #000;
        }
        button{
            width: 2rem;
            height: .6rem;
            font-size: .3rem;
        }
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
