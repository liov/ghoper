<template>
  <div>IndexDB</div>
</template>

<script>
export default {
  data() {
    return {
      DB: {
        name: 'hoper',
        version: 1,
        db: null
      }
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
  async mounted() {
    this.DB.db = await this.openDB(this.DB.name, this.DB.version)
    this.addData(this.DB.db, 'moments', this.momentList)
  },
  methods: {
    openDB: function(name, version = 1) {
      return new Promise(function(resolve, reject) {
        const request = window.indexedDB.open(name)
        request.onerror = function(e) {
          reject(e)
        }
        request.onsuccess = function(e) {
          resolve(e.target.result)
        }
        request.onupgradeneeded = function(e) {
          const db = e.target.result
          if (!db.objectStoreNames.contains('moments')) {
            db.createObjectStore('moments', { keyPath: 'id' })
          }
          console.log('DB version changed to ' + version)
        }
      })
    },
    closeDB: function(db) {
      db.close()
    },
    deleteDB: function(name) {
      indexedDB.deleteDatabase(name)
    },
    addData: function(db, storeName, data) {
      const transaction = db.transaction(storeName, 'readwrite')
      const store = transaction.objectStore(storeName)

      for (let i = 0; i < data.length; i++) {
        store.add(data[i])
      }
    },
    getDataByKey: function(db, storeName, key) {
      const transaction = db.transaction(storeName, 'readwrite')
      const store = transaction.objectStore(storeName)
      const request = store.get(key)
      request.onsuccess = function(e) {
        const data = e.target.result
        console.log(data.name)
      }
    },
    updateDataByKey: function(db, storeName, key, newdata) {
      const transaction = db.transaction(storeName, 'readwrite')
      const store = transaction.objectStore(storeName)
      const request = store.get(key)
      request.onsuccess = function(e) {
        let data = e.target.result
        data = newdata
        store.put(student)
      }
    },
    deleteDataByKey: function(db, storeName, key) {
      const transaction = db.transaction(storeName, 'readwrite')
      const store = transaction.objectStore(storeName)
      store.delete(key)
    },
    clearObjectStore: function(db, storeName) {
      const transaction = db.transaction(storeName, 'readwrite')
      const store = transaction.objectStore(storeName)
      store.clear()
    }
  }
}
</script>

<style scoped>
</style>
