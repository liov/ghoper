<template>
    <div class="hoper">
        <van-cell-group>

            <van-field
                    v-model="moment.content"
                    label="瞬间"
                    type="textarea"
                    placeholder="请输入内容"
                    rows="1"
                    autosize
            />

            <van-field
                    v-model="moment.mood_name"
                    label="心情"
                    type="textarea"
                    placeholder="请输入心情"
                    rows="1"
                    autosize
            />

            <van-checkbox-group v-model="Tags">
                <van-cell-group>
                    <van-cell
                            v-for="(item,index) of existTags"
                            clickable
                            :key="item"
                            :title="item"
                            @click="toggle(index)"
                    >
                        <van-checkbox :name="item" ref="checkboxes" />
                    </van-cell>
                </van-cell-group>
            </van-checkbox-group>

            <van-field v-model="tag" placeholder="请输入新标签"><van-button slot="button" size="small" type="primary" @click="addTag">添加</van-button></van-field>

            <li>
                谁可以查看：
                <input type="radio" v-model="moment.permission" value="0" />全部可见
                <input type="radio" v-model="moment.permission" value="1" />自己可见
                <input type="radio" v-model="moment.permission" value="2" />好友可见
                <input type="radio" v-model="moment.permission" value="3" />陌生人可见
            </li>
        </van-cell-group>

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
            },
            toggle(index) {
                this.$refs.checkboxes[index].toggle();
            }
        }
    }
</script>

<style lang="scss" scoped>
    .hoper{
        li{
            background-color: #95a5a6;
        }

    }
</style>
