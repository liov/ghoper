import bodyParser from 'body-parser'

export default {
    router: {
        base: '/'
    },
    build: {
        vendor: ['axios'],
/*        babel: {
            presets: ['es2015', 'stage-0']
        }*/
    },

/*    modules: [
        '@nuxtjs/axios',
        '@nuxtjs/proxy'
    ],
    proxy: [
        ['/api', {target: 'https://hoper.xyz/', pathRewrite: { '^/api': '/api/breeds/image/random' }}]
    ],*/
    css: [

        { src: '~assets/common.scss', lang: 'scss' },
    ],
    plugins: [
        {src:'~plugins/filter/vue.js',ssr: true},
        {src:'~plugins/filter/hoper_axios.js'},
        { src: '~plugins/iview.js', ssr: false }
    ],

};
