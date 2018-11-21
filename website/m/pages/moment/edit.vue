<template>
    <div class="hoper">
        <ul>
            <li>瞬间</li>
            <li><textarea rows="10" cols="30" v-model="moment.content"></textarea></li>

            <li>心情：<input type="text" v-model="moment.mood_name"/></li>

            <li>标签：
                <span v-for="item of existTags">
                    <input type="checkbox"  v-model="tags" :value="item"/>{{item}}
                </span>

                <div class="addTag">添加标签</div>
                <input type="text" v-model="tag"/>
                <div @click="addTag">添加</div>
            </li>
            <li>
                谁可以查看：
                <input type="radio" v-model="moment.permission" value="0"/>全部可见
                <input type="radio" v-model="moment.permission" value="1"/>自己可见
                <input type="radio" v-model="moment.permission" value="2"/>好友可见
                <input type="radio" v-model="moment.permission" value="3"/>陌生人可见
            </li>
        </ul>
        <div class="commit" @click="commit">提交</div>
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
            console.log(res.data.data);
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
        },
        mounted:function(){

            this.oldMoment  = this.copy(this.moment);
            this.oldTags    = this.tags.concat();
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
                            alert(res.data.msg)
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
                    this.Tags.push(this.tag);
                    this.tag = '';
                } else {
                    if (this.tag === '')
                        alert("标签为空");
                    else
                        alert("标签重复");
                }
            },
            copy: function (obj) {
                var newobj = {};
                for (var attr in obj) {
                    newobj[attr] = obj[attr];
                }
                return newobj;
            }
        },
        filters: {}
    }
</script>

<style scoped>
    .commit {
        width: 100px;
        height: 50px;
        background-color: #47494e;
        border-radius: 50%;
        margin-left: 50px;
        color: #f7f8fb;
        text-align: center;
        line-height: 50px;
    }
</style>
