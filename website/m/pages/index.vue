<template>
    <div>
        <div>
        <nuxt-link to="/moment"><van-button type="primary">瞬间</van-button></nuxt-link>
        <a class="index" href="javascript:;" @click="logout">
            <van-button type="warning">注销</van-button>
        </a>
        <a href="/api/tpl/index"><van-button type="default">模板首页</van-button></a>
        <van-button type="primary" @click="hint('啥都没有发生')">默认按钮</van-button>
            <nuxt-link to="/user/signup"><van-button type="primary">注册</van-button></nuxt-link>
        </div>
        <div class="center">
        <nuxt-link class="chat" to="/chat"><van-button type="primary">点这个来聊天啊</van-button></nuxt-link>
        </div>
        <lio-hint ref="lioHint"></lio-hint>
        <div class="van-hairline--top"></div>
        <div >
            <van-button type="primary" @click="test">hoper</van-button>
        </div>
        <van-cell>{{testdata}}</van-cell>
    </div>
</template>

<script>
    import { delCookie } from "../plugins/utils/cookie"
    import axios from 'axios'
    export default {
        data(){
            return {
                testdata:''
            }
        },
        methods:{
            logout:function () {
                let vm =this;
                axios.get(`/api/user/logout`).then(
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
            },
            test:function () {
                let vm =this;
                let Init = { method: 'GET',
                    headers: new Headers(),
                    mode: 'cors',
                    cache: 'default' };

                fetch("http://hoper.xyz/api/tpl/index",Init).then((res)=>
                    res.text()).then((text)=>{
                        vm.testdata = text
                    }
                )
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
    .center{
        position: fixed;
        top: 50%;
        width: 100%;
        text-align: center;
        .chat{
            font-size: .6rem;
            color: cadetblue;
        }
    }

</style>
