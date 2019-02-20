<template>
  <div>
    <!-- 淡入 -->
    <transition :name="transitionName">
      <div
        v-show="show"
        class="lio-hint"
      >
        {{ msg }}
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      transitionName: '',
      msg: ''
    }
  },
  mounted() {},
  methods: {
    animate(msg) {
      this.show = true
      this.msg = msg
      this.transitionName = 'van-fade'
      setTimeout(() => {
        this.show = false
      }, 500)
    }
  }
}
</script>

<style scoped>
.lio-hint {
  top: 50%;
  left: 50%;
  width: 100px;
  height: 30px;
  font-size: 0.4rem;
  line-height: 30px;
  text-align: center;
  position: fixed;
  border-radius: 3px;
  margin: -50px 0 0 -50px;
  background-color: #3000ff8c;
}
</style>
