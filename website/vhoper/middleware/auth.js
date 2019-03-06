export default async function({ store, $axios, error, req, route, redirect }) {
  if (!store.state.user) {
    if (process.server) {
      if (req.headers.cookie) {
        const cookieparser = require('cookieparser')
        const parsed = cookieparser.parse(req.headers.cookie)
        if (parsed.token) {
          store.commit('SET_TOKEN', parsed.token)
          $axios.defaults.headers.common.Cookie = req.headers.cookie
          await $axios
            .$get(`/api/user/get`, { timeout: 3000 })
            .then(res => {
              if (res.code === 200) {
                const user = res.data
                store.commit('SET_USER', user)
              } else {
                store.commit('SET_TOKEN', null)
                redirect({
                  path: '/user/login?callbackUrl=' + route.path
                })
              }
            })
            .catch(function() {
              store.commit('SET_TOKEN', null)
              redirect({
                path: '/user/login?callbackUrl=' + route.path
              })
            })
        }
      } else {
        store.commit('SET_TOKEN', null)
        redirect({ path: '/user/login?callbackUrl=' + route.path })
      }
    } else {
      await $axios
        .$get(`/api/user/get`, { timeout: 3000 })
        .then(res => {
          if (res.code === 200) {
            const user = res.data
            store.commit('SET_USER', user)
          } else {
            store.commit('SET_TOKEN', null)
            route.push({
              path: '/user/login?callbackUrl=' + route.currentRoute.path
            })
          }
        })
        .catch(e => {
          console.log(e)
          // redirect({ path: '/' })
        })
    }
  }
}
