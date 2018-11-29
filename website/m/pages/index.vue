<template>
    <div>
        <nuxt-link to="/moment">瞬间</nuxt-link>
        <a class="index" href="javascript:;" @click="signout">
            注销
        </a>
        <a href="/tpl/index">模板首页</a>
        <nuxt-link class="chat" to="/chat">点这个来聊天啊</nuxt-link>
        <lio-hint ref="lioHint"></lio-hint>
        <div class="van-hairline--top"></div>
        <van-cell
                is-link
                title="Fade"
                @click="hint('登录成功')"
        />
    </div>
</template>

<script>
    import { delCookie } from "../plugins/utils/cookie"
    import axios from 'axios'
    import hint from "./common/hint";
    export default {
        data(){
            return {

            }
        },
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
            },
            hint:function (msg) {
                this.$refs.lioHint.animate(msg)
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
