<template>
    <div class="hoper">
        <van-nav-bar
                left-text="返回"
                left-arrow
                fixed
                @click-left="$router.push('/')"
                @click-right="$router.push('/')"
        >
            <span slot="title">修改瞬间</span>
            <van-icon name="home" size=".5rem" slot="right" />
        </van-nav-bar>

        <div class="black"></div>

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

            <div class="center"><van-button type="default">标签</van-button></div>
            <van-checkbox-group v-model="oldTags">
                <van-row v-for="(item,index) in tagsGroup" :key="index">
                    <van-cell-group>
                        <van-col span="8" v-for="(item1,index1) in item" :key="index1">
                            <van-cell
                                    clickable
                                    :title="item1"
                                    @click="toggle(index*3+index1)"
                            >
                                <van-checkbox :name="item1" ref="checkboxes" />
                            </van-cell>
                        </van-col>
                    </van-cell-group>
                </van-row>
            </van-checkbox-group>

            <div class="center"><van-button type="default">添加标签</van-button></div>
            <van-field v-model="tag" placeholder="请输入新标签"><van-button slot="button" size="small" type="primary" @click="addTag">添加</van-button></van-field>
            <van-radio-group v-model="moment.permission">
                <div class="center"><van-button type="default">权限</van-button></div>
                <van-row>
                    <van-cell-group>
                        <van-col span="12">
                            <van-cell title="全部可见" clickable @click="moment.permission = '0'">
                                <van-radio name="0" />
                            </van-cell>
                        </van-col>
                        <van-col span="12">
                            <van-cell title="自己可见" clickable @click="moment.permission = '1'">
                                <van-radio name="1" />
                            </van-cell></van-col>
                    </van-cell-group>
                </van-row>
                <van-row>
                    <van-cell-group>
                        <van-col span="12">
                            <van-cell title="好友可见" clickable @click="moment.permission = '2'">
                                <van-radio name="2" />
                            </van-cell>
                        </van-col>
                        <van-col span="12">
                            <van-cell title="陌生人可见" clickable @click="moment.permission = '3'">
                                <van-radio name="3" />
                            </van-cell></van-col>
                    </van-cell-group>
                </van-row>
            </van-radio-group>
        </van-cell-group>

        <div class="center"><van-button type="primary" @click="commit">提交</van-button></div>

    </div>
</template>

<script>
    import axios from 'axios'
    import cookie from '../../plugins/utils/cookie'
    export default {

/*        fetch ({ store, redirect }) {
            if (!store.state.authUser) {
                return redirect('/')
            }
        },*/
        data() {
            return {
                existTags: ["韩雪", "徐峥", "胡歌", "张卫健"],
                tagsGroup:[],
                tag: '',
                oldMoment:{},
                oldTags:[],
                moment:{
                    id:0,
                    tags:[]
                }
            }
        },
/*        async fetch({store,req}){
            let token =await cookie.getCookie("token",req)
            store.dispatch('SET_TOKEN', token)
        },*/
        asyncData({query}) {
        return axios.get(`/api/moment/${query.id}?t=${query.t}&index=${query.index}`).then((res) => {
            let tags = [];
            for (let v of res.data.data.tags) {
                tags.push(v.name)
            }
            return {moment: res.data.data, tags: tags||[]}
        })

        },
        created:function(){
            /*  let vm = this
               request.getMonment().then(res =>{
                   vm.momentList = res
               })*/
            this.tagsGroup =  this.tagGroup(this.existTags,3)
        },
        mounted:function(){
            this.oldMoment  = this.copy(this.moment);
            this.oldTags    = this.tags.concat();
            this.moment.permission =this.moment.permission.toString()
        },
        methods: {
            getMoment: function () {
                this.moment = localStorage.getItem("moment_" + this.$route.params.id);
            },
            commit: function () {
                let vm = this;
                let newMoment = {};

                if (vm.moment.content !== vm.oldMoment.content) {
                    newMoment.content = vm.moment.content;
                }
                if (vm.moment.permission !== vm.oldMoment.permission) {
                    newMoment.permission = vm.moment.permission;
                }
                if (vm.moment.mood_name !== vm.oldMoment.mood_name) {
                    newMoment.mood_name = vm.moment.mood_name;
                }
                if (vm.moment.image_url !== vm.oldMoment.image_url) {
                    newMoment.image_url = vm.moment.image_url;
                }
                if (vm.tags.toString() !== vm.oldTags.toString() ) {
                    newMoment.tags = [];
                    for (let v of vm.tags) {
                        newMoment.tags.push({name: v})
                    }
                }

                axios.put('/api/moment/' + vm.$route.query.id+'?t='+vm.$route.query.t+'&index='+vm.$route.query.index,newMoment)
                    .then(function (res) { //
                        // success
                        if (res.data.msg === '修改成功')
                            vm.$router.push({path: '/moment'});
                        else{
                            vm.$toast(res.data.msg)
                        }
                    })
                    .catch(function (err) {
                        // error
                        //vm.$router.push({path:'/user/signin'});
                    });

            },
            addTag: function () {
                if ((this.tag !== '') && (this.existTags.indexOf(this.tag) === -1)) {
                    this.existTags.push(this.tag);
                    if (this.tagsGroup[this.tagsGroup.length-1].length === 3){
                        this.tagsGroup.push([this.tag])
                    }else {
                        this.tagsGroup[this.tagsGroup.length-1].push(this.tag)
                    }
                    this.oldTags.push(this.tag);
                    this.tag = '';
                } else {
                    if (this.tag === '')
                        this.$toast("标签为空");
                    else
                        this.$toast("标签重复");
                }
            },
            copy: function (obj) {
                let newobj = {};
                for (let attr in obj) {
                    newobj[attr] = obj[attr];
                }
                return newobj;
            },
            toggle(index) {
                this.$refs.checkboxes[index].toggle()
            },
            tagGroup:function (arr, size) {
                let arr2=[];
                for(let i=0;i<arr.length;i=i+size){
                    arr2.push(arr.slice(i,i+size));
                }
                return arr2;
            }
        },
        filters: {}
    }
</script>

<style lang="scss" scoped>
    .black{
        margin-top: 1.5rem;
    }
    .center{
        text-align: center;
        .van-button--default {
            padding: 0 0;
            width: 100%;
        }
    }
</style>
