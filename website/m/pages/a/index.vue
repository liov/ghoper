<template>
<div>
    {{msg.data.name}}
</div>
</template>
<script>
  import axios from "axios"
    export default {
        async asyncData() {
            let { data } = await axios.get(`/api/a`);
            return {msg:data }
            }
        }
</script>
