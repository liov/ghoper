<template>
<div>
{{Msg}}
</div>
</template>
<script>
  import axios from "axios"
    export default {
        async asyncData() {

                let res = await axios.get(`/a`);

                return {Msg: res.data.data}
            }
        }
</script>
