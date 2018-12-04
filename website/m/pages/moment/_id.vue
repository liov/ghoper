<template>
    <div>
        <van-nav-bar
                left-text="返回"
                left-arrow
                fixed
                @click-left="onClickLeft"
                @click-right="onClickRight"
        >
            <span slot="title" >瞬间</span>
            <van-icon name="home" size=".6rem" slot="right" />
        </van-nav-bar>

        <van-cell-group>
        <van-cell class="content">{{moment.content}}</van-cell>
            <van-cell>{{moment.created_at|dateFormat}}</van-cell>
            <van-cell><van-tag plain v-for="(tag,index) in moment.tags" :key="index">{{tag.name}}</van-tag></van-cell>
            <van-cell>{{moment.mood_name}}</van-cell>
            <van-cell>浏览量：{{moment.browse_count}},评论数：{{moment.comment_count}},喜欢：{{moment.love_count}},收藏：{{moment.collect_count}}</van-cell>
        </van-cell-group>
        <div class="button">
        <nuxt-link v-if="belongFlag" :to="{ path: '/moment/edit',query: queryPamram}"><van-button type="primary">修改</van-button></nuxt-link>
       <van-button v-if="belongFlag" type="danger"  @click="deleteMoment">删除</van-button>
        </div>
    </div>
</template>

<script>
    import {copy} from "../../plugins/utils/utils.js"
    import axios from 'axios'
    export default {
        middleware: 'auth',
        data(){
            return {
                queryPamram:{
                    id: this.$route.params.id,
                    t :  this.$route.query.t,
                    index :  this.$route.query.index
                }
            }
        },
        async asyncData ({ params,query}) {
            let res = await axios.get(`/api/moment/${params.id}?t=${query.t}&index=${query.index}`)
            //console.log(res.data.data);
            return {
                moment: res.data.data,
                belongFlag: res.data.msg==='belong'
            }
        },
        methods:{
            getMoment:function () {
               this.moment = localStorage.getItem("moment_"+this.$route.params.id);
            },
            removeMoment:function () {
                localStorage.removeItem("moment_"+this.$route.params.id);
            },
            deleteMoment:function () {
                let vm =this;
                axios.delete('/api/moment/' + vm.$route.params.id+'?t='+vm.$route.query.t+'&index='+vm.$route.query.index)
                    .then(function (res) { //
                        // success
                        if (res.data.msg === '删除成功')
                            vm.$router.push({path: '/moment'});
                        else
                            vm.$toast(res.data.msg)
                    })
                    .catch(function (err) {
                        // error
                    });
            },
            onClickLeft() {
                this.$router.go(-1)
            },
            onClickRight() {
                this.$router.push('/')
            },
            tagGroup:function (arr, size) {
                let arr2=[];
                for(let i=0;i<arr.length;i=i+size){
                    arr2.push(arr.slice(i,i+size));
                }
                return arr2;
            }
        }
    }
</script>

<style scoped>
    .content{
        margin-top: 1.5rem;
    }
    .button{
        text-align: center;
    }
</style>
