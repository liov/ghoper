import bodyParser from 'body-parser'

export default {
    router: {
        base: '/'
    },
    build: {
        vendor: ['axios','~plugins/filter/hoper_axios.js','vant'],
        'autoprefixer': {
            browsers: ['Android >= 4.0', 'iOS >= 7']
        },
        'postcss-pxtorem': {
            rootValue: 37.5,
            propList: ['*']
        }
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
        {src:'~plugins/filter/hoper_vue.js'},
        {src:'~plugins/filter/hoper_axios.js'},
        {src: '~plugins/iview.js', ssr: false },
    ],

};
