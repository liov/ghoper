<template>
  <div>
    <div>db</div>
    <div id="status" />
  </div>
</template>

<script>
export default {
  data() {
    return {
      db: null,
      msg: ''
    }
  },
  async asyncData({ $axios }) {
    const params = {
      pageNo: 0,
      pageSize: 5
    }
    const res = await $axios.$get(`/api/moment`, { params })
    return {
      momentList: res.data,
      total: res.count,
      topCount: res.top_count,
      user_action:
        res.user_action != null
          ? res.user_action
          : {
              collect: [],
              like: [],
              approve: [],
              comment: [],
              browse: []
            }
    }
  },
  mounted() {
    this.db = openDatabase('hoper', '1.0', 'hoper DB', 2 * 1024 * 1024)
    this.db.transaction(function(tx) {
      tx.executeSql('CREATE TABLE IF NOT EXISTS Moments (id unique, content)')
    })
    const vm = this
    this.db.transaction(function(tx) {
      for (const item of vm.momentList) {
        tx.executeSql(
          `INSERT INTO Moments (id, content) VALUES (${item.id},'${
            item.content
          }')`
        )
      }
    })

    this.db.transaction(function(tx) {
      tx.executeSql(
        'SELECT * FROM Moments',
        [],
        function(tx, results) {
          const len = results.rows.length
          let i
          let msg = '<p>查询记录条数: ' + len + '</p>'
          document.querySelector('#status').innerHTML += msg

          for (i = 0; i < len; i++) {
            msg = '<p><b>' + results.rows.item(i).content + '</b></p>'
            document.querySelector('#status').innerHTML += msg
          }
        },
        null
      )
    })
  },
  methods: {}
}
</script>

<style scoped>
</style>
