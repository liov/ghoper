<template>
    <div id="app">
        <div class="chat">
            <div class="row" v-for="item of chatContent">
                <p>{{item}}</p>
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
    export default {
        data() {
            return{
                ws: null, // Our websocket
                newMsg: '', // Holds new messages to be sent to the server
                chatContent: [], // A running list of chat messages displayed on the screen
                recipient: null, // Email address used for grabbing an avatar
                sender: null, // Our username
                joined: false // True if email and username have been filled in
            }
        },

        created: function() {
            //运行在服务端

        },
        mounted:function(){
            //不能放在created里
            let vm = this;
            this.chatContent=JSON.parse(localStorage.getItem("chatContent"));
            if(this.chatContent === null) vm.chatContent=[];
            vm.ws = new WebSocket('ws://'+window.location.host+`/pai/chat/ws/${params.id}`);
            vm.ws.addEventListener('message', function(e) {
                let msg = JSON.parse(e.data);
                vm.chatContent.push(msg.content);
                localStorage.setItem("chatContent",JSON.stringify(vm.chatContent))
            });
        },
        methods: {
            send: function () {
                let vm = this;
                if (this.newMsg !== '') {
                    this.ws.send(
                        JSON.stringify({
                                recipient_user_id: vm.recipient,
                                send_user_id: vm.$store.state.user!==null?vm.$store.state.user.id:'',
                                content: vm.newMsg // Strip out html
                            }
                        ));
                    this.newMsg = ''; // Reset newMsg
                }
            },

        }
    }
</script>

<style lang="scss" scoped>
    .row{
        width: 100%;
        height: 50px;
        background-color: aliceblue;
        color: blueviolet;
        line-height: 50px;
        padding-left: 50px;
        font-size: 20px;
    }
    .inp{
        position: fixed;
        bottom: 20%;
        left: 10%;
        width: 100%;
        height: 50px;
    }
</style>
