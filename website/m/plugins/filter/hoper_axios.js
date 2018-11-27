import axios from "axios";
import cookie from "../utils/cookie"
let context = {};
axios.interceptors.request.use(
    config => {
            let token;
        if (typeof window !== 'undefined'){
             token = localStorage.getItem("token");
            config.baseURL = "http://"+window.location.host
        }else {
            token = context.store.state.token;
            if(!token){
                token = cookie.getCookie("token",context.req);
            }
            //config.baseURL = "http://"+context.req.host
            //坑，用nginx转发，目前只能写死，或者在ng上改？
            config.baseURL = "http://hoper.xyz";
        }

        if (token) {  // 判断是否存在token，如果存在的话，则每个http header都加上token
            config.headers.Authorization = token;
        }

        return config;
    },
    err => {
        return Promise.reject(err);
    });



axios.interceptors.response.use(
    response => {
        return response;  //请求成功的时候返回的data
    },
    error => {
        try {
            if (error.response) {
                switch (error.response.status) {
                    case 401://token过期，清除token并跳转到登录页面
                        context.app.router.push({path:'/user/signin'})

                        return;
                }
            }
            return Promise.reject(error.response.data)
        }
        catch (e) {
        }
    });

//奇淫巧技，这里可以获取context
/*export default ({ app: { router }, req, res }) => {
    router.afterEach((to, from) => {
        if (typeof window === 'undefined') {
            cookie.refreshTokenCookie(req, res)
        }else {
            window.router = router
        }
    })
}*/


//简单理解一下，对外导出以后，这个参数就会被赋值？而且在初始化的时候只调用一次
export default (function getApp({app}) {
       context = app.context
    })

