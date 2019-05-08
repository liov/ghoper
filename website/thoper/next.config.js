const withTypescript = require('@zeit/next-typescript');
const path = require('path');
const withCss = require('@zeit/next-css')
const withLess = require('@zeit/next-less')
const lessToJS = require('less-vars-to-js')
const fs = require('fs')

// Where your antd-custom.less file lives
const themeVariables = lessToJS(
    fs.readFileSync(path.resolve(__dirname, './pages/assets/antd-custom.less'), 'utf8')
)

// fix: prevents error when .less files are required by node
if (typeof require !== 'undefined') {
    require.extensions['.less'] = file => {}
    require.extensions['.css'] = file => {}
}

module.exports = withTypescript(withCss(withLess({
    useFileSystemPublicRoutes: false,
 /*   cssModules: true,
    cssLoaderOptions: {
        camelCase: true,
        namedExport: true
    },*/
    lessLoaderOptions: {
        javascriptEnabled: true,
        modifyVars: themeVariables, // make your antd custom effective
    },
    webpack (config, options) {
        config.resolve = {
            ...config.resolve,
            ...{
                alias: {
                    ...config.resolve.alias,
                    '@src': path.resolve(__dirname, 'client'),
                }
            },
        };
 /*       if (!options.isServer) {
            /!* Using next-css *!/
            for (let entry of options.defaultLoaders.css) {
                if (entry.loader === 'css-loader') {
                    entry.loader = 'typings-for-css-modules-loader'
                    break
                }
            }
            for (let entry of options.defaultLoaders.less) {
                if (entry.loader === 'css-loader') {
                    entry.loader = 'typings-for-css-modules-loader';
                    break
                }
            }
        }*/
        return config
    },
    webpackDevMiddleware: config => {
        // Perform customizations to webpack dev middleware config
        // Important: return the modified config
        return config;
    }
})))


