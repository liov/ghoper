<template>
    <div class="hoper">
        <ul>
        <li><span>昵称：</span><input type="text" placeholder="必填" v-model="user.name"/></li>
            <li><span>密码：</span><input type="password" placeholder="必填（最小6位）" v-model="user.password"/></li>
            <li><span>邮箱：</span><input type="email" placeholder="必填" v-model="user.email"/></li>
            <li><span>性别：</span> <input class="sex" type="radio" v-model="user.sex" value="0" />男
                <input class="sex" type="radio" v-model="user.sex" value="1" />女</li>
            <li><span>手机号：</span><input type="tel" placeholder="选填" v-model="user.phone"/></li>
            <a class="commit" href="javascript:;" @click="commit">
                注册
            </a>
        </ul>
    </div>
</template>

<script>
    import axios from "axios"

    export default {
        data(){
            return {
                user:{}
            }
        },
        methods:{
            commit:function () {
                let vm = this;
                if (vm.user.phone === ""){
                    vm.user.phone = null
                }
                vm.user.sex =  parseInt(vm.user.sex);
                axios.post(`/api/user/signup`,vm.user)
                    .then((res)=> { //
                        // success
                        if (res.data.msg === 'ok'){
                            vm.$toast("注册成功")
                            vm.$router.replace("/user/signin")
                        }else {
                            vm.$toast(res.data.msg)
                        }
                    })
                    .catch(function(err) {
                        // error
                    });
            }
        }
    }
</script>

<style lang="scss" scoped>
    .hoper{
        font-size: .5rem;
        ul {
            li{
                margin-top: .5rem;
                margin-left: .2rem;
                span{
                    width: 2rem;
                    height: 1rem;
                    background-color: aquamarine;
                    text-align: center;
                    color: cadetblue;
                }
                input{
                    width: 5rem;
                    height: 1rem;
                    border: solid 1px;
                    font-size: .5rem;
                    margin-left: .2rem;
                }
                .sex{
                    width: .5rem;
                    height: .5rem;
                    top:.2rem;
                }
            }
        }
    }
</style>
