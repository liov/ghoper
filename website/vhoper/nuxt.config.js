const { resolve } = require('path')
const { readFileSync } = require('fs')
const pkg = require('./package')
const copyRes = require('./plugins/copyRes')
module.exports = {
  mode: 'universal',
  server: {
    /*    https: {
      key: readFileSync(resolve(__dirname, '../../config/tls/cert.key')),
      cert: readFileSync(resolve(__dirname, '../../config/tls/cert.pem'))
    } */
  },
  /*
   ** Headers of the page
   */
  head: {
    title: 'hoper',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: pkg.description }
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
  },

  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },

  /*
   ** Global CSS
   */
  css: [
    // 'ant-design-vue/dist/antd.less',
    '@/assets/less/antd.less',
    'normalize.css/normalize.css',
    '@/assets/css/agent_clean.css'
  ],

  /*
   ** Plugins to load before mounting the App
   */
  plugins: [
    '@/plugins/filter/hoper_antd',
    '@/plugins/filter/hoper_utils',
    '@/plugins/filter/ctx-inject',
    '@/plugins/filter/hoper_axios',
    '@/plugins/filter/hoper_ssr',
    { src: '@/plugins/filter/hoper_nossr', ssr: false }
  ],

  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/pwa',
    '@nuxtjs/apollo'
  ],
  /*
   ** Axios module configuration
   */
  axios: {
    // See https://github.com/nuxt-community/axios-module#options
    //baseURL: 'https://hoper.xyz',
    //browserBaseURL: 'https://hoper.xyz',
    proxy: false
  },
  apollo: {
    tokenName: 'yourApolloTokenName', // optional, default: apollo-token
    tokenExpires: 10, // optional, default: 7 (days)
    includeNodeModules: true, // optional, default: false (this includes graphql-tag for node_modules folder)
    authenticationType: 'Basic', // optional, default: 'Bearer'
    // optional
    //errorHandler:'怕了你了,errorHandler should be a path to an exported error handler config.',
    // required
    clientConfigs: {
      default: {
        // required
        httpEndpoint: 'https://hoper.xyz/api/graphql',
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
        httpEndpoint: 'https://hoper.xyz/api/graphql',
        // wsEndpoint: 'ws://localhost/ws/echo',
        tokenName: 'apollo-token'
      }
      // alternative: user path to config which returns exact same config options
      // test2: '~/plugins/graphql.js'
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
/*      if (ctx.isDev && ctx.isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(ts|js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/,
          options: {
            fix: true
          }
        })
      }*/
      if(ctx.isServer){
        copyRes.init()
      }
    },
    loaders: {
      less: {
        /*       modifyVars: {
          'primary-color': '#1DA57A',
          'link-color': '#1DA57A',
          'border-radius-base': '2px'
        }, */
        javascriptEnabled: true
      }
    },
    postcss: {
      plugins: {
        // https://github.com/postcss/postcss-import
        // 'postcss-import': false
        /* {
          resolve: function (id,basedir,importOptions) {
            if(id.match(/~/) && id.match(/~/).index === 0){
              console.log(join(__dirname,'node_modules',id.substr(1))+"实验")
              return join(__dirname,'node_modules',id.substr(1))
            }
          }
        } */
      }
    }
  }
}
