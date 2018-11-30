import bodyParser from 'body-parser'

export default {
    router: {
        base: '/'
    },
    build: {
        vendor: ['axios','~plugins/filter/hoper_axios.js','vant','mint-ui','iview'],
        'autoprefixer': {
            browsers: ['Android >= 4.0', 'iOS >= 7']
        },
        'postcss-pxtorem': {
            rootValue: 37.5,
            propList: ['*']
        },
        babel:{        //配置按需引入规则
            plugins:[
                [
                    "component",
                    {
                        "libraryName": "mint-ui",
                        "style": true
                    }
                ]
            ]
        },
        /*
         ** Run ESLINT on save
         */
/*        extend (config, { isClient }) {
            if (isClient) {
                config.module.rules.push({
                    enforce: 'pre',
                    test: /\.(js|vue)$/,
                    loader: 'eslint-loader',
                    exclude: /(node_modules)/
                })
            }
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
        'vant/lib/index.css',
        'mint-ui/lib/style.css'
    ],
    plugins: [
        {src:'~plugins/filter/hoper_vue.js'},
        {src:'~plugins/filter/hoper_axios.js'},
    ],

};
