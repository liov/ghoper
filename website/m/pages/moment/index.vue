<!--suppress ALL -->
<template>

    <div class="hoper">
        <div class="head">
        <nuxt-link to="/moment/add"><van-button type="primary">添加</van-button></nuxt-link>

        <a class="first" href="javascript:;" @click="index">
            <van-button type="primary">瞬间首页</van-button>
        </a>
        <a class="previous" href="javascript:;" @click="previous">
            <van-button type="primary">上一页</van-button>
        </a>
        <a class="next" href="javascript:;" @click="next">
            <van-button type="primary">下一页</van-button>
        </a>
        <nuxt-link class="index" to="/"><van-button type="primary">首页</van-button></nuxt-link>
        </div>
        <div v-for="(item, index) in momentList.top_moments">
            <div  class="moment" v-if="item.content !==''">
            <nuxt-link :to="{ path: '/moment/'+item.id,query: { t:topNum ,index:pageNo*pageSize+index}}">
                <div>
                    <p><span>[置顶]</span>{{item.content}}</p>
                    <p>{{item.created_at|dateFormat}}</p>
                    <p>{{item.mood_name}}</p>
                    <p><span v-for="tag in item.tags">{{tag.name}}&nbsp;</span></p>
                </div>
            </nuxt-link>
            </div>
        </div>

        <div v-for="(item, index) in momentList.normal_moments">
            <div  class="moment" v-if="item.content !==''">
            <nuxt-link :to="{ path: '/moment/'+item.id,query: { t: '0',index:pageNo*pageSize+index}}">
                <div>
                    <p>{{item.content}}</p>
                    <p>{{item.created_at|dateFormat}}</p>
                    <p>{{item.mood_name}}</p>
                    <p><span v-for="tag in item.tags">{{tag.name}}&nbsp;</span></p>
                </div>
            </nuxt-link>
            </div>
        </div>


    </div>

</template>

<script>
    import {copy} from "../../plugins/utils/utils.js"
    import axios from "axios"
    export default {
        //middleware: 'auth',
        data() {
            return {
                pageNo:0,
                pageSize:5,
                topNum:0,
                momentList:{},
                lastFlag :false,
                firstFlag : true
            }
        },
        async asyncData() {
            let params = {
                pageNo : 0,
                pageSize:5,
            };
            let res = await axios.get(`/api/moment`,{params});
            return {oldMomentList: res.data.data}
        },
        created: function () {
            this.momentList  = copy(this.oldMomentList);
            /*  let vm = this
               request.getMonment().then(res =>{
                   vm.momentList = res
               })*/
        },
        mounted:function(){
            this.topNum = this.oldMomentList.top_moments===null?0:this.momentList.top_moments.length
            //this.momentList.normal_moments=  this.momentList.normal_moments.splice(0, this.size-(this.momentList.top_moments!==null?this.momentList.top_moments.length:0))
        },
        computed:{
            normalMomentsStart:function () {
                if (this.pageNo>0){
                    return this.pageNo*this.pageSize-this.topNum
                } else {
                    return 0
                }

            }
        },
        watch:{
            pageNo:function () {
                let vm =this
                if (this.pageNo>0) {
                    vm.momentList.top_moments = null;
                } else {
                    vm.momentList.top_moments = vm.oldMomentList.top_moments;
                }
                vm.momentList.normal_moments = vm.oldMomentList.normal_moments.slice(vm.normalMomentsStart,vm.normalMomentsStart+vm.pageSize);
            }
        },
        methods: {
            setMoment: function (moment) {
                localStorage.setItem("moment_" + moment.id, moment);
            },
            next:function(){
                if (this.lastFlag){
                    this.$toast("最后一页")
                    return
                }
                let vm = this
                let params = {
                    t:vm.topNum,
                    pageNo : vm.pageNo+1,
                    pageSize:vm.pageSize
                };

                axios.get(`/api/moment`,{params}).then((res1) => { //
                    // success
                   let momentList = res1.data.data;
                    if (momentList===undefined) {
                        vm.lastFlag = true
                        vm.$toast("最后一页")
                    } else {
                        vm.oldMomentList.normal_moments=vm.oldMomentList.normal_moments.concat(momentList.normal_moments);
                        //之所以放这里，用了vue的属性侦听watch
                        vm.pageNo = vm.pageNo +1;
                        vm.firstFlag = false
                    }

                }).catch(function (err) {
                        // error
                    });
            },
            previous:function () {
                if(this.firstFlag){
                    this.$toast("已经是首页")
                    return
                }
                this.lastFlag = false
                if (this.pageNo>0){
                    this.pageNo = this.pageNo -1;
                } else {
                    this.$toast("已经是首页")
                }
            },
            index:function () {
                this.pageNo = 0;
                this.firstFlag = true;
                this.lastFlag = false;
            }
        },
        filters: {}
    }
</script>

<style lang="scss" scoped>

    .head{
        margin-left: .2rem;
        button {
            padding: 0 .35rem;
        }
    }
    .moment {
        position: relative;
        background-color: aqua;
        margin-top: .5rem;
        height: 3rem;
        padding: 10px;
    }
</style>
