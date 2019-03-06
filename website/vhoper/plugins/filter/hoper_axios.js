import cookie from '../utils/cookie'

export default function({ app, store, $axios, req }) {
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
  /*  $axios.onRequestError(error => {
    return Promise.reject(error)
  }) */
  $axios.onResponseError(error => {
    if (typeof window !== 'undefined') {
      return (error.response.status = 300)
    }
    if (error.response) {
      switch (error.response.status) {
        case 401: // token过期，清除token并跳转到登录页面
          if (typeof window !== 'undefined') {
            localStorage.removeItem('token')
          }
          app.router.push({
            path: '/user/login?callbackUrl=' + app.router.currentRoute.path
          })
      }
    }
    return Promise.reject(error.response.data)
  })
  /*  $axios.onError(error => {
    const code = parseInt(error.response && error.response.status)
    if (code === 400) {
      redirect('/400')
    }
  }) */
}
