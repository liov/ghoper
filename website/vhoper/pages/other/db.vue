<template>

</template>

<script>
    export default {
      data(){
        return{
          DB:{
            name:'hoper',
            version:1,
            db:null
          }
        }
      },
      methods:{
        openDB: function (name) {
          var version=version || 1;
      const request=window.indexedDB.open(name);
      request.onerror=function(e){
        console.log('OPen Error!');
      };
      request.onsuccess=function(e){
        this.DB.db=e.target.result;
      };
          request.onupgradeneeded=function(e){
            var db=e.target.result;
            if(!db.objectStoreNames.contains('students')){
              db.createObjectStore('students',{keyPath:"id"});
            }
            console.log('DB version changed to '+version);
          };
    },
        closeDB:function(db){
          db.close();
        },
      }
        deleteDB:function (name) {
          indexedDB.deleteDatabase(name);
        },
      addData:function (db,storeName) {
        var transaction=db.transaction(storeName,'readwrite');
        var store=transaction.objectStore(storeName);

        for(var i=0;i<students.length;i++){
          store.add(students[i]);
        }
      },
      getDataByKey:function (db,storeName,value) {
        var transaction=db.transaction(storeName,'readwrite');
        var store=transaction.objectStore(storeName);
        var request=store.get(value);
        request.onsuccess=function(e){
          var student=e.target.result;
          console.log(student.name);
        };
      },
      updateDataByKey:function(db,storeName,value){
        var transaction=db.transaction(storeName,'readwrite');
        var store=transaction.objectStore(storeName);
        var request=store.get(value);
        request.onsuccess=function(e){
          var student=e.target.result;
          student.age=35;
          store.put(student);
        };
    },
      deleteDataByKey:function (db,storeName,value) {
        var transaction=db.transaction(storeName,'readwrite');
        var store=transaction.objectStore(storeName);
        store.delete(value);
      },
      clearObjectStore:function (db,storeName) {
        var transaction=db.transaction(storeName,'readwrite');
        var store=transaction.objectStore(storeName);
        store.clear();
      }
    }
</script>

<style scoped>

</style>
