<template>
    <div></div>

</template>

<script>
    import axios from "axios";
    export default {

        mounted(){
            let vm =this;
            axios.get(`/api/user/active/${this.$route.params.id}/${this.$route.params.secret}`)
                .then((res)=>{
                    //vm.$router.push({path: '/'});
                    alert(res.data.msg)
                    vm.$router.push({path: '/user/sign'})
                });

        }
    }
</script>

<style scoped>

</style>
