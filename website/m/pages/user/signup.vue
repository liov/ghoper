<template>
    <div class="hoper">
        <van-cell-group>
            <van-field v-model="user.name"
                       label="用户名"
                       required
                       placeholder="请输入用户名(必填)" />
            <van-field v-model="user.password"
                       label="密码"
                       type="password"
                       required
                       placeholder="请输入密码(必填（最小6位）)" />
            <van-field v-model="user.email"
                       label="邮箱"
                       type="email"
                       required
                       placeholder="请输入邮箱(必填)" />
            <van-row>
                <van-radio-group v-model="user.sex">
                    <van-col span="12">
                        <van-cell title="男" clickable @click="user.sex = '0'">
                            <van-radio name="0" />
                        </van-cell>

                    </van-col>
                    <van-col span="12">
                        <van-cell title="女" clickable @click="user.sex = '1'">
                            <van-radio name="1" />
                        </van-cell>

                    </van-col>
                </van-radio-group>
            </van-row>
            <van-field v-model="user.phone"
                       label="用户名"
                       type="tel"
                       placeholder="请输入手机号(选填)" />
        </van-cell-group>
        <div class="button">
            <van-button type="primary" @click="commit">注册</van-button>
        </div>
    </div>
</template>

<script>
    import axios from "axios"

    export default {
        data(){
            return {
                user:{
                    sex:'0'
                }
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
                            vm.$toast("注册成功");
                            vm.$router.replace("/user/login")
                        }else {
                            vm.$toast(res.data.msg)
                            vm.user.sex = vm.user.sex.toString()
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
  .button{
      text-align: center;
  }
</style>
