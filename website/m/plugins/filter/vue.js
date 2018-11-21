import Vue from 'vue'



Vue.filter('dateFormat', function (value) {
    let date = new Date(value);
    return date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate() + ' ' + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds();
});


