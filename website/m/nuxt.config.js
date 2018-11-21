let bodyParser = require('body-parser')
let session = require('express-session')
module.exports={
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
        {src:'~plugins/filter/axios.js'},
        { src: '~plugins/iview.js', ssr: false }
    ],
    serverMiddleware: [
        // body-parser middleware
        bodyParser.json(),
        // session middleware
        session({
            secret: 'oihopBHIDUHFifGFEwehoa',
            resave: false,
            saveUninitialized: false,
            cookie: { maxAge: 60000 }
        }),
        // Api middleware
        // We add /api/login & /api/logout routes
        '~/api'
    ]
};
