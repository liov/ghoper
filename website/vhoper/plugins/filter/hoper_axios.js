import cookie from '../utils/cookie'

// 奇淫巧技，这里可以获取context
/* export default ({ app: { router }, req, res }) => {
    router.afterEach((to, from) => {
        if (typeof window === 'undefined') {
            cookie.refreshTokenCookie(req, res)
        }else {
            window.router = router
        }
    })
} */

// 简单理解一下，对外导出以后，这个参数就会被赋值？而且在初始化的时候只调用一次
/* export default (function getApp({ app }) {
  context = app.context
}) */
export default function({ app, store, $axios, req, redirect }) {
  $axios.onRequest(config => {
    let token
    if (typeof window !== 'undefined') {
      token = localStorage.getItem('token')
      // config.baseURL = 'http://' + window.location.host
    } else {
      token = store.state.token
      if (!token) {
        token = cookie.getCookie('token', req)
      }
      // config.baseURL = "http://"+context.req.host
      // 坑，用nginx转发，目前只能写死，或者在ng上改？
      // config.baseURL = 'http://hoper.xyz'
    }

    if (token) {
      // 判断是否存在token，如果存在的话，则每个http header都加上token
      config.headers.Authorization = token
    }

    return config
  })
  $axios.onRequestError(error => {
    return Promise.reject(error)
  })
  $axios.onResponseError(error => {
    try {
      if (error.response) {
        switch (error.response.status) {
          case 401: // token过期，清除token并跳转到登录页面
            if (typeof window !== 'undefined') {
              localStorage.removeItem('token')
            }
            app.router.push({
              path: '/user/login?callbackUrl=' + app.router.currentRoute.path
            })

            return
        }
      }
      return Promise.reject(error.response.data)
    } catch (e) {}
  })
  $axios.onError(error => {
    const code = parseInt(error.response && error.response.status)
    if (code === 400) {
      redirect('/400')
    }
  })
}
