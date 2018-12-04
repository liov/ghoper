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
            <div class="center"><van-button type="default">标签</van-button></div>
            <van-checkbox-group v-model="Tags">
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
                tagsGroup:[],
                Tags:[],
                tag:''
            }
        },

        created:function () {
            /*  let vm = this
               request.getMonment().then(res =>{
                   vm.momentList = res
               })*/
          this.tagsGroup =  this.tagGroup(this.existTags,3)
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
                    if (this.tagsGroup[this.tagsGroup.length-1].length === 3){
                        this.tagsGroup.push([this.tag])
                    }else {
                        this.tagsGroup[this.tagsGroup.length-1].push(this.tag)
                    }
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

<style lang="scss" scoped>
    .hoper{
        li{
            background-color: #95a5a6;
        }
        .center{
            text-align: center;
            .van-button--default {
                padding: 0 0;
                width: 100%;
            }
        }
    }
</style>
