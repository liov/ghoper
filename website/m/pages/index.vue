<template>
    <div>
        <div>
        <nuxt-link to="/moment"><van-button type="primary">瞬间</van-button></nuxt-link>
        <a class="index" href="javascript:;" @click="signout">
            <van-button type="warning">注销</van-button>
        </a>
        <a href="/tpl/index"><van-button type="default">模板首页</van-button></a>
        <van-button type="primary" @click="hint('啥都没有发生')">默认按钮</van-button>
        </div>
        <nuxt-link class="chat" to="/chat"><van-button type="primary">点这个来聊天啊</van-button></nuxt-link>
        <lio-hint ref="lioHint"></lio-hint>
        <div class="van-hairline--top"></div>


    </div>
</template>

<script>
    import { delCookie } from "../plugins/utils/cookie"
    import axios from 'axios'
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
                        vm.$toast("注销成功")
                    });
            },
            hint:function (msg) {
                //this.$refs.lioHint.animate(msg)
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
