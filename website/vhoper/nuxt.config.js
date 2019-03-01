const pkg = require('./package')
const {join} = require('path')
module.exports = {
    mode: 'universal',

    /*
    ** Headers of the page
    */
    head: {
        title: 'hoper',
        meta: [
            {charset: 'utf-8'},
            {name: 'viewport', content: 'width=device-width, initial-scale=1'},
            {hid: 'description', name: 'description', content: pkg.description}
        ],
        link: [{rel: 'icon', type: 'image/x-icon', href: '/favicon.ico'}]
    },

    /*
    ** Customize the progress-bar color
    */
    loading: {color: '#fff'},

    /*
    ** Global CSS
    */
    css: ['ant-design-vue/dist/antd.css'],

    /*
    ** Plugins to load before mounting the App
    */
    plugins: [
        '@/plugins/filter/hoper_vue',
        '@/plugins/filter/hoper_axios',
        '@/plugins/markdown-it'
    ],

    /*
    ** Nuxt.js modules
    */
    modules: [
        // Doc: https://axios.nuxtjs.org/usage
        '@nuxtjs/axios',
        '@nuxtjs/pwa',
        '@nuxtjs/apollo',
        '@nuxtjs/markdownit'
    ],
    /*
    ** Axios module configuration
    */
    axios: {
        // See https://github.com/nuxt-community/axios-module#options
    },

    apollo: {
        tokenName: 'yourApolloTokenName', // optional, default: apollo-token
        tokenExpires: 10, // optional, default: 7 (days)
        includeNodeModules: true, // optional, default: false (this includes graphql-tag for node_modules folder)
        authenticationType: 'Basic', // optional, default: 'Bearer'
        // optional
        errorHandler(error) {
            console.log('%cError', 'background: red; color: white; padding: 2px 4px; border-radius: 3px; font-weight: bold;', error.message)
        },
        // required
        clientConfigs: {
            default: {
                // required
                httpEndpoint: 'http://hoper.xyz/api/graphql',
                // optional
                // See https://www.apollographql.com/docs/link/links/http.html#options
                httpLinkOptions: {
                    credentials: 'same-origin'
                },
                // You can use `wss` for secure connection (recommended in production)
                // Use `null` to disable subscriptions
                // wsEndpoint: 'ws://localhost/api/chat/ws', // optional
                // LocalStorage token
                tokenName: 'apollo-token', // optional
                // Enable Automatic Query persisting with Apollo Engine
                persisting: false, // Optional
                // Use websockets for everything (no HTTP)
                // You need to pass a `wsEndpoint` for this to work
                websocketsOnly: false // Optional
            },
            test: {
                httpEndpoint: 'http://hoper.xyz/api/graphql',
                //wsEndpoint: 'ws://localhost/ws/echo',
                tokenName: 'apollo-token'
            },
            // alternative: user path to config which returns exact same config options
            //test2: '~/plugins/graphql.js'
        }
    },
    /*
    ** Build configuration
    */
    build: {
        /*
        ** You can extend webpack config here
        */
        extend(config, ctx) {
            // Run ESLint on save
            if (ctx.isDev && ctx.isClient) {
                config.module.rules.push({
                    enforce: 'pre',
                    test: /\.(js|vue)$/,
                    loader: 'eslint-loader',
                    exclude: /(node_modules)/,
                    options: {
                        fix: true
                    }
                })
            }
        },
        postcss: {
            plugins: {
                // https://github.com/postcss/postcss-import
                //'postcss-import': false
                /*{
                  resolve: function (id,basedir,importOptions) {
                    if(id.match(/~/) && id.match(/~/).index === 0){
                      console.log(join(__dirname,'node_modules',id.substr(1))+"实验")
                      return join(__dirname,'node_modules',id.substr(1))
                    }
                  }
                }*/
            }
        }
    }
}
