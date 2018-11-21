<!--suppress ALL -->
<template>

    <div class="hoper">
        <div class="add">
            <nuxt-link to="/moment/add">添加</nuxt-link>
        </div>
        <a class="first" href="javascript:;" @click="index">
            首页
        </a>
        <a class="previous" href="javascript:;" @click="previous">
            上一页
        </a>
        <a class="next" href="javascript:;" @click="next">
            下一页
        </a>
        <nuxt-link class="index" to="/">首页</nuxt-link>
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
            //console.log(res.data.data);
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
                    alert("最后一页")
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
                        alert("最后一页")
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
                    alert("已经是首页")
                }
                this.lastFlag = false
                if (this.pageNo>0){
                    this.pageNo = this.pageNo -1;
                } else {
                    alert("已经是首页")
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

    .add {
        position: fixed;
        top: 0;
        height: 20px;
        background-color: aquamarine;
        z-index: 1;
    }
    .index{
        position: fixed;
        top: 0;
        left: 1rem;
        height: 20px;
        background-color: aquamarine;
        z-index: 1;
    }
    .next{
        position: fixed;
        top: 0;
        left: 2rem;
        height: 20px;
        background-color: aquamarine;
        z-index: 1;
    }
    .previous{
        position: fixed;
        top: 0;
        left: 3rem;
        height: 20px;
        background-color: aquamarine;
        z-index: 1;
    }
    .first{
        position: fixed;
        top: 0;
        left: 4rem;
        height: 20px;
        background-color: aquamarine;
        z-index: 1;
    }
    .moment {
        position: relative;
        background-color: aqua;
        margin-top: .5rem;
        height: 3rem;
        padding: 10px;
    }
</style>
