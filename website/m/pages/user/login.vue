<template>
<div class="hoper">
    <div class="black"></div>
    <div class="inp">
        <van-cell-group>
            <van-field v-model="user.input"
                       label="用户名"
                       required
                       placeholder="请输入邮箱或手机号" />
            <van-field v-model="user.password"
                       label="密码"
                       type="password"
                       required
                       placeholder="请输入密码" />
        </van-cell-group>
    </div>
    <div class="button">
        <van-button type="primary"  @click="commit">登录</van-button>
        <nuxt-link class="reg" to="/user/signup"> <van-button type="primary">注册</van-button></nuxt-link>
    </div>
</div>
</template>

<script>
    import axios from 'axios'
    export default {
        data(){
            return {
                user:{}
            }
        },
        methods:{
            commit:function () {
                let vm = this;
                let emailReg = new RegExp("^([a-zA-Z0-9]+[_\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$");
                let phonelReg =  /^1[0-9]{10}$/;
                if (emailReg.test(vm.user.input)){
                    vm.user.email = vm.input
                } else if(phonelReg.test(vm.user.input)) {
                      vm.user.phone = vm.user.input
                }

                axios.post(`/api/user/login`,vm.user)
                    .then((res)=> { //
                        // success
                        if (res.data.msg === '登录成功'){
                            localStorage.setItem("token",res.data.token);
                            vm.$store.commit("SET_USER", res.data.data);
                            vm.$store.commit("SET_TOKEN", res.data.token);
                            localStorage.setItem("user",res.data.data.id);
                            vm.$toast("登录成功");
                            vm.$router.replace("/")
                        }else if(res.data.msg === "账号未激活"){
                            vm.$toast(res.data.user.email)
                        }else {
                            vm.$toast(res.data.msg)
                        }
                    })
                    .catch(function(err) {
                        console.log(err)
                    });
            },
            async login() {
                try {
                    await this.$store.dispatch('login', {
                        username: this.formUsername,
                        password: this.formPassword
                    });
                    this.formUsername = '';
                    this.formPassword = '';
                    this.formError = null
                } catch (e) {
                    this.formError = e.message
                }
            },
            async logout() {
                try {
                    await this.$store.dispatch('logout')
                } catch (e) {
                    this.formError = e.message
                }
            }
        }
    }
</script>

<style lang="scss" scoped>
    .black{
        height: 3rem;
    }
    .button{
        position: absolute;
        width: 100%;
        top: 6rem;
        text-align: center;
    }

</style>
