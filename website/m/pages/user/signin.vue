<template>
<div class="hoper">
    <ul>
        <li class="acct"><span>账号：</span><input type="email" placeholder="邮箱或手机号" v-model="user.input"/></li>
        <li class="pass"><span>密码：</span><input type="password" v-model="user.password"/></li>
        <a class="commit" href="javascript:;" @click="commit">
            登录
        </a>
        <nuxt-link class="reg" to="/user/signup">注册</nuxt-link>
    </ul>
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
                    console.log("email");
                    vm.user.email = vm.input
                } else if(phonelReg.test(vm.user.input)) {
                    console.log("phone");
                      vm.user.phone = vm.user.input
                }

                axios.post(`/api/user/signin`,vm.user)
                    .then((res)=> { //
                        // success
                        if (res.data.msg === 'ok'){
                            localStorage.setItem("token",res.data.token);
                            vm.$store.commit("SET_USER", res.data.user);
                            vm.$store.commit("SET_TOKEN", res.data.token);
                            localStorage.setItem("user",res.data.user.id);
                            //alert("登录成功");
                            vm.$router.replace("/chat")
                        }else {
                            alert(res.data.msg)
                        }
                    })
                    .catch(function(err) {
                        // error
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
.hoper{
    position: absolute;
    top: 20%;
    left: 10%;
    width: 100%;
    font-size: .8rem;
    .acct{
        position: absolute;
        top: 1rem;
        left: 10%;
        input{
            position: absolute;
            top: 1rem;
            left: 10%;
            height: 0.8rem;
            width: 5rem;
            font-size: .5rem;
            border: solid 1px;
        }
    }
    .pass{
        position: absolute;
        top:  3rem;
        left: 10%;
        input{
            position: absolute;
            top: 1rem;
            left: 10%;
            height: 0.8rem;
            width: 5rem;
            font-size: .5rem;
            border: solid 1px;
        }
    }
    .commit{
        position: absolute;
        top: 5rem;
        left: 1rem;
    }
    .reg{
        position: absolute;
        top: 5rem;
        left: 3rem;
    }
}
</style>
