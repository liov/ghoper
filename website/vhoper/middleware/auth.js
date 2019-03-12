export default async function({ store, $axios, error, req, route, redirect }) {
  if (!store.state.user) {
    if (process.server) {
      if (req.headers.cookie) {
        const cookieparser = require('cookieparser')
        const parsed = cookieparser.parse(req.headers.cookie)
        if (parsed.token) {
          store.commit('SET_TOKEN', parsed.token)
          $axios.defaults.headers.common.Cookie = req.headers.cookie
          const res = await $axios
            .$get(`/api/user/get`, { timeout: 3000 })
            .catch(() => {
              store.commit('SET_TOKEN', null)
              redirect({
                path: '/user/login?callbackUrl=' + route.path
              })
            })
          if (res !== undefined && res.code === 200) {
            const user = res.data
            store.commit('SET_USER', user)
          } else {
            store.commit('SET_TOKEN', null)
            redirect({
              path: '/user/login?callbackUrl=' + route.path
            })
          }
        }
      } else {
        store.commit('SET_TOKEN', null)
        redirect({ path: '/user/login?callbackUrl=' + route.path })
      }
    } else {
      const res = await $axios
        .$get(`/api/user/get`, { timeout: 3000 })
        .catch(() => {
          if (typeof window !== 'undefined') {
            localStorage.removeItem('token')
          }
          store.commit('SET_TOKEN', null)
          redirect({
            path: '/user/login?callbackUrl=' + route.path
          })
        })
      // axios的返回错误处理把这里直接废了
      if (res !== undefined && res.code === 200) {
        const user = res.data
        store.commit('SET_USER', user)
      } else {
        if (typeof window !== 'undefined') {
          localStorage.removeItem('token')
        }
        store.commit('SET_TOKEN', null)
        redirect({
          path: '/user/login?callbackUrl=' + route.path
        })
      }
    }
  }
}
