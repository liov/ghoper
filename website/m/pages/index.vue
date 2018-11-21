<template>
    <div>
        <nuxt-link to="/moment">瞬间</nuxt-link>
        <a class="index" href="javascript:;" @click="signout">
            注销
        </a>
        <nuxt-link class="chat" to="/chat">点这个来聊天啊</nuxt-link>
    </div>
</template>

<script>
    import { delCookie } from "../plugins/utils/cookie"
    import axios from 'axios'
    export default {
        methods:{
            signout:function () {
                let vm =this;
                axios.get(`/api/user/signout`).then(
                    (res)=>{
                        if(res.data.code===200)
                            localStorage.removeItem("token");
                        //delCookie("token");
                        vm.$store.commit("SET_USER", null);
                        alert("注销成功")
                    });

            }
        }
    }
</script>

<style lang="scss" scoped>

    .add {
        position: fixed;
        top: 0;
        height: 20px;
        background-color: aquamarine;
        z-index: 1;
    }
    .chat{
        position: fixed;
        top: 50%;
        text-align: center;
        font-size: .6rem;
        color: cadetblue;
    }
</style>
