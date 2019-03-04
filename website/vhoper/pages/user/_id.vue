<template>
  <a-row>
    <a-col :span="6" style="text-align: right">
      <a-avatar shape="square" :size="100" :src="user.avatar_url" />
    </a-col>
    <a-col :span="16">
      <a-form-item
        label="用户名"
        :label-col="{ span: 2 }"
        :wrapper-col="{ span: 12 }"
      >
        <a-input
          v-model="user.name"
        >
          <a-icon slot="prefix" type="user" />
        </a-input>
      </a-form-item>

      <a-form-item
        label="性别"
        :label-col="{ span: 2 }"
        :wrapper-col="{ span: 12 }"
      >
        <a-radio-group
          v-model="user.sex"
        >
          <a-radio-button value="男">
            男
          </a-radio-button>
          <a-radio-button value="女">
            女
          </a-radio-button>
        </a-radio-group>
      </a-form-item>
      <a-row>
        <a-col :span="12" />
        <a-col :span="12" />
      </a-row>
    </a-col>
    <a-col :span="2" />
  </a-row>
</template>

<script>
export default {
  middleware: 'auth',
  async asyncData({ $axios }) {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await $axios.$get(`/api/user/edit`, { params })
    return {
      user: res.data
    }
  },
  created() {},
  methods: {
    getStatus: function() {
      fetch('http://hoper.xyz/user/1', {
        method: 'get',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({})
      }).then(res => {
        return res.json().status
      })
    }
  }
}
</script>

<style scoped>
</style>
