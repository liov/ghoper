<template>
  <a-row>
    <a-col :span="8" />
    <a-col :span="8">
      <a-form-item
        label="事件"
        :label-col="{ span: 3 }"
        :wrapper-col="{ span: 6 }"
      >
        <a-input v-model="event" />
      </a-form-item>
      <a-button @click="setEvent">
        设置事件
      </a-button>
    </a-col>
    <a-col :span="8" />
  </a-row>
</template>

<script>
export default {
  data() {
    return {
      event: ''
    }
  },
  mounted() {
    const client = new EventSource('/api/get/events')
    client.onmessage = function (evt) {
      console.log(evt)
    }
  },
  methods: {
    setEvent() {
      this.$axios.$get('/api/set/events?event=' + this.event)
    }
  }
}
</script>
<style scoped></style>
