import bodyParser from 'body-parser'

export default {
    router: {
        base: '/'
    },
    build: {
        postcss:{
            'autoprefixer': {
                browsers: ['Android >= 4.0', 'iOS >= 7']
            },
            'postcss-pxtorem': {
                rootValue: 37.5,
                propList: ['*']
            }
        },
        babel:{        //配置按需引入规则
            plugins:[

/*                    ["import",
                    {
                        libraryName: "vant",
                        libraryDirectory: "es",
                        style: "index.css"
                    }, "vant"],*/
 /*               ["component", {
                    "libraryName": "vant",
                    "libraryDirectory": "es",
                    "style": "index.css"
                }],*/
                ["component",
                    {
                        libraryName: "mint-ui",
                        style: true
                    }, "mint-ui"]

/*                ['import',
                    {
                    libraryName: 'vant',
                    libraryDirectory: 'es',
                    style: true
                }]*/
            ]
        },
        /*
         ** Run ESLINT on save
         */
        /*        extend (config, { isClient,isServer }) {
                    if(isServer){
                        config.module.rules.push({
                            loader: 'babel',
                            query: {
                                plugins: [["import", { libraryName: 'vant',
                                    libraryDirectory: 'es',
                                    style: true
                              }]]
                            }
                        })
                    }
        /*            if (isClient) {
                        config.module.rules.push({
                            enforce: 'pre',
                            hoper: /\.(js|vue)$/,
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
        '@/assets/common.scss',
        '@/assets/agentClean.css',
        'vant/lib/index.css',
        //'mint-ui/lib/style.css'
    ],
    plugins: [
        {src:'~plugins/filter/hoper_axios.js'},
        {src:'~plugins/filter/hoper_antd.js',ssr: true}
    ],

};
