module.exports = {
    devServer: {
        open: process.platform === 'darwin',
        host: '127.0.0.1',
        port: 3000,
        https: false,
        hotOnly: false,
        proxy: null, // 设置代理
        before: app => {}
    },
}
