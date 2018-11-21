<template>
    <div>哈哈哈</div>

</template>

<script>
    import axios from "axios";
    export default {

        mounted(){
            axios.get(`/api/user/active/${this.$route.params.id}/${this.$route.params.secret}`)
                .then((res)=>{
                    console.log(res.data.data);
                    //vm.$router.push({path: '/'});
                    alert(res.data.data)
                });

        }
    }
</script>

<style scoped>

</style>
