<template>
    <div>
        <ul>
            <li>
                <input type="text" v-model="message"></input>
            </li>
        </ul>
        <ul>
            <li>
                <input type="text" v-model="nsqMsg.func_name"></input>
            </li>
            <li>
                <input type="text" v-model="param"></input>
            </li>
        </ul>
            <div @click="commit">发送</div>
    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        data(){
            return {
                message:'',
                nsqMsg:{},
                param:''
            }
        },
        methods:{
            commit:function () {
                let vm =this;
                if (vm.message !== ''){
                    let formData = new FormData();
                    formData.append("message",vm.message);
                    axios.post(`/api/nsq?st=0`,formData,{headers: {
                            'Content-Type': 'application/x-www-form-urlencoded'
                        }})
                } else {
                    vm.nsqMsg.params = vm.param.split(",");
                    let i = 0;
                    for(let item of vm.nsqMsg.params) {
                        if(item.indexOf("\"")===-1){
                            vm.nsqMsg.params[i] = Number(item)
                        }
                        else {
                            vm.nsqMsg.params[i] = item.slice(1,-1)
                        }
                        i++
                    }
                    axios.post(`/api/nsq?st=1`,vm.nsqMsg)
                }
                vm.message = ''
            }
        }
    }
</script>

<style lang="scss" scoped>
    input{
        width: 50%;
        height: .8rem;
        font-size: .3rem;
    }
</style>
