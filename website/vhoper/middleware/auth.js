import axios from 'axios'

export default async function({ store, error, req, route, redirect }) {
  if (!store.state.user) {
    if (process.server) {
      if (req.headers.cookie) {
        const cookieparser = require('cookieparser')
        const parsed = cookieparser.parse(req.headers.cookie)
        if (parsed.token) {
          store.commit('SET_TOKEN', parsed.token)
          axios.defaults.headers.common.Cookie = req.headers.cookie
          await axios
            .get(`/api/user/get`, { timeout: 3000 })
            .then(res => {
              if (res.data.code === 200) {
                const user = res.data.data
                store.commit('SET_USER', user)
              } else {
                redirect({
                  path: '/user/login?callbackUrl=' + route.currentRoute.path
                })
              }
            })
            .catch(function() {
              route.push({
                path: '/user/login?callbackUrl=' + route.currentRoute.path
              })
            })
        }
      } else {
        redirect({ path: '/user/login?callbackUrl=' + route.currentRoute.path })
      }
    } else {
      await axios
        .get(`/api/user/get`, { timeout: 3000 })
        .then(res => {
          if (res.data.code === 200) {
            const user = res.data.data
            store.commit('SET_USER', user)
          } else {
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
