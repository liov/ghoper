export default async function({ store, $axios, error, req, route, redirect }) {
  if (!store.state.user) {
    let login = false
    if (
      process.server &&
      req.headers.cookie &&
      req.headers.cookie.indexOf('token') !== -1
    ) {
      $axios.defaults.headers.common.Cookie = req.headers.cookie
    }

    const res = await $axios.$get(`/api/user/get`, { timeout: 3000 })

    if (res !== undefined && res.code === 200) {
      login = true
      const user = res.data
      store.commit('SET_USER', user)
    }

    if (!login) {
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
