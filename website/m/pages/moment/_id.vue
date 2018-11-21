<template>
    <div>
        <p>{{moment.content}}</p>
        <p>{{moment.created_at|dateFormat}}</p>
        <p><span v-for="tag in moment.tags">{{tag.name}}&nbsp;</span></p>
        <p>{{moment.mood_name}}</p>
        <p>浏览量：{{moment.browse_count}},评论数：{{moment.comment_count}},喜欢：{{moment.love_count}},收藏：{{moment.collect_count}}</p>
        <nuxt-link v-if="belongFlag" :to="{ path: '/moment/edit',query: queryPamram}">修改</nuxt-link>
        <a v-if="belongFlag" href="javascript:;" @click="deleteMoment">删除</a>
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
                        if (res.data.msg === 'ok')
                            vm.$router.push({path: '/moment'});
                        else
                            alert(res.data.msg)
                    })
                    .catch(function (err) {
                        // error
                    });
            }
        }
    }
</script>

<style scoped>

</style>
