<template>
    <div class="hoper">
        <ul>
            <li>瞬间</li>
            <li><textarea rows="10" cols="30" v-model="moment.content"></textarea></li>

            <li>心情：<input type="text" v-model="moment.mood_name"/></li>

            <li>标签：
                <span v-for="item of existTags">
                    <input  type="checkbox" v-model="Tags" :value="item" />{{item}}
                </span>

                <div class="addTag">添加标签</div> <input type="text" v-model="tag"/><div  @click="addTag">添加</div>
            </li>
            <li>
                谁可以查看：
                <input type="radio" v-model="moment.permission" value="0" />全部可见
                <input type="radio" v-model="moment.permission" value="1" />自己可见
                <input type="radio" v-model="moment.permission" value="2" />好友可见
                <input type="radio" v-model="moment.permission" value="3" />陌生人可见
            </li>
        </ul>
        <van-button type="primary" @click="commit">提交</van-button>
    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        middleware: 'auth',
        data() {
            return {
                moment:{
                    content:'',
                    mood_name:'',
                    tags:[],
                    permission:0,
                },
                existTags:["韩雪","徐峥","胡歌","张卫健"],
                Tags:[],
                tag:''
            }
        },

        created:function () {
            /*  let vm = this
               request.getMonment().then(res =>{
                   vm.momentList = res
               })*/
        },
        methods:{
            commit:function(){
                let vm =this;
                this.moment.permission = parseInt(this.moment.permission);
                vm.moment.tags = [];
                for( let i of this.Tags){
                    this.moment.tags.push({name:i})
                }
               axios.post(`/api/moment`,this.moment)
                   .then(function(res) { //

                   // success
                       if (res.data.msg === '新建成功')
                           vm.$router.push({path:'/moment'});
                       else
                           vm.$toast(res.data.msg)
               })
                   .catch(function(err) {
                       // error
                   });


            },
            addTag:function () {
                if ((this.tag !== '')&&(this.existTags.indexOf(this.tag)===-1)){
                    this.existTags.push(this.tag);
                    this.Tags.push(this.tag);
                    this.tag = '';
                } else {
                    if (this.tag === '')
                        this.$toast("标签为空");
                    else
                        this.$toast("标签重复");
                }
            }
        }
    }
</script>

<style lang="scss" scoped>

</style>
