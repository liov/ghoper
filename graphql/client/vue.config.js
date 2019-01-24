module.exports = {
    devServer: {
        open: process.platform === 'darwin',
        host: '127.0.0.1',
        port: 3000,
        https: false,
        hotOnly: false,
        proxy: {
            '/api': {    //将www.exaple.com印射为/apis
                target: 'http://localhost:8888',  // 接口域名
                secure: false,  // 如果是https接口，需要配置这个参数
                changeOrigin: true,  //是否跨域
                pathRewrite: {
                    '^/api': ''   //需要rewrite的,
                }
            }
        }, // 设置代理
        before: app => {}
    },
}
