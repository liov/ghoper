<!--suppress ALL -->
<template>

    <div class="hoper">
        <van-nav-bar
                left-text="返回"
                left-arrow
                fixed
                @click-left="onClickLeft"
                @click-right="onClickRight"
        >
            <span slot="title" @click="index">瞬间</span>
            <van-icon name="edit" size=".6rem" slot="right" />
        </van-nav-bar>

    <div class="all-moment">
        <div v-for="(item, index) in momentList.top_moments">
            <div  class="moment" v-if="item.content !==''">
            <nuxt-link :to="{ path: '/moment/'+item.id,query: { t:topNum ,index:(pageNo-1)*pageSize+index}}">
                <div>
                    <van-cell><span><van-button type="primary">[置顶]</van-button></span>
                        <van-field
                            :value="item.content"
                            type="textarea"
                            :rows="item.content.length/20"
                            disabled
                    />
                    </van-cell>
                    <van-row>
                        <van-col span="8"><van-cell><van-button plain type="primary">作者</van-button>{{item.user.name}}</van-cell></van-col>
                        <van-col span="16"><van-cell class="date"><van-button plain type="primary">日期</van-button>{{item.created_at|dateFormat}}</van-cell></van-col>
                    </van-row>
                    <van-cell-group>
                        <van-cell><van-button plain type="primary">心情</van-button>{{item.mood_name}}</van-cell>
                        <van-cell><van-button plain type="primary">标签</van-button><van-tag plain v-for="(tag,index) in item.tags" :key="index">{{tag.name}}</van-tag></van-cell>
                    </van-cell-group>
                </div>
            </nuxt-link>
            </div>
        </div>

        <div v-for="(item, index) in momentList.normal_moments">
            <div  class="moment" v-if="item.content !==''">
            <nuxt-link :to="{ path: '/moment/'+item.id,query: { t: '0',index:(pageNo-1)*pageSize+index}}">
                <div>
                    <van-cell>
                        <van-field
                                :value="item.content"
                                type="textarea"
                                :rows="Math.ceil(item.content.length/30)"
                                disabled
                        />
                    </van-cell>
                    <van-row>
                        <van-col span="8"><van-cell><van-button plain type="primary">作者</van-button>{{item.user.name}}</van-cell></van-col>
                        <van-col span="16"><van-cell class="date"><van-button plain type="primary">日期</van-button>{{item.created_at|dateFormat}}</van-cell></van-col>
                    </van-row>
                    <van-cell-group>
                        <van-cell><van-button plain type="primary">心情</van-button>{{item.mood_name}}</van-cell>
                        <van-cell><van-button plain type="primary">标签</van-button><van-tag plain v-for="(tag,index) in item.tags" :key="index">{{tag.name}}</van-tag></van-cell>
                    </van-cell-group>
                </div>
            </nuxt-link>
            </div>
        </div>
</div>
        <div class="pagination">
        <van-pagination
                v-model="pageNo"
                :total-items="125"
                :show-page-size="pageSize"
                force-ellipses
        />
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
                pageNo:1,
                pageNum:1,
                pageSize:5,
                topNum:0,
                momentList:{},
                lastFlag :false,
                firstFlag : true,
                loading: false,
                finished: false
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
                if (this.pageNo>1){
                    return (this.pageNo-1)*this.pageSize-this.topNum
                } else {
                    return 0
                }
            }
        },
        watch:{
            pageNo:async function () {
                if(this.pageNo>this.pageNum){
                    if (this.lastFlag){
                        this.$toast("最后一页")
                        this.pageNo = this.pageNo -1
                        this.pageNum = this.pageNo
                        return
                    }
                    await this.next()
                }
                if (this.pageNo>1) {
                    this.momentList.top_moments = null;
                } else {
                    this.momentList.top_moments = this.oldMomentList.top_moments;
                }
                this.momentList.normal_moments = this.oldMomentList.normal_moments.slice(this.normalMomentsStart,this.normalMomentsStart+this.pageSize);
            }
        },
        methods: {
            setMoment: function (moment) {
                localStorage.setItem("moment_" + moment.id, moment);
            },
            next:async function(){
                let vm = this
                    let params = {
                        t: vm.topNum,
                        pageNo: vm.pageNo - 1,
                        pageSize: vm.pageSize
                    };
                    //这里可以这么写，async，await函数，或者 return axios().then((res)=>{})返回Promise
                    let res = await axios.get(`/api/moment`, {params})

                        let momentList = res.data.data;
                        if (momentList === undefined) {
                            vm.lastFlag = true
                            vm.$toast("最后一页")
                            vm.pageNo = vm.pageNo - 1
                        } else {
                            vm.oldMomentList.normal_moments = vm.oldMomentList.normal_moments.concat(momentList.normal_moments);
                            //之所以放这里，用了vue的属性侦听watch
                            vm.pageNum = vm.pageNo;
                            vm.firstFlag = false;
                        }
            },
            previous:function () {
                if(this.firstFlag){
                    this.$toast("已经是首页")
                    return
                }
                this.lastFlag = false
                if (this.pageNo>1){
                    this.pageNo = this.pageNo -1;
                } else {
                    this.$toast("已经是首页")
                }
            },
            index:function () {
                this.pageNo = 1;
                this.firstFlag = true;
                this.lastFlag = false;
            },
            onLoad() {
                // 异步更新数据
                setTimeout(() => {
                    for (let i = 0; i < 10; i++) {
                        this.list.push(this.list.length + 1);
                    }
                    // 加载状态结束
                    this.loading = false;

                    // 数据全部加载完成
                    if (this.list.length >= 40) {
                        this.finished = true;
                    }
                }, 500);
            },
            onClickLeft() {
                this.$router.push('/')
            },
            onClickRight() {
               this.$router.push('/moment/add')
            }
        },
        filters: {}
    }
</script>

<style lang="scss" scoped>


    .date{
        font-size: .2rem;
    }
    .all-moment{
        margin-top: 1.5rem;
        margin-bottom: 1.5rem;
        .moment {
            margin-top: .5rem;
            .van-button{
                height: 20px;
                line-height: 20px;
                padding: 0 3px;
                margin-right: 3px;
            }
        }
    }
    .pagination{
        position: fixed;
        bottom: 0;
        width: 100%;
    }
</style>
