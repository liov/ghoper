const withTypescript = require('@zeit/next-typescript');
const path = require('path');
const withCSS = require('@zeit/next-css')
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

const nextConfig = (nextConfig = []) => {
    return {
        webpack(config, options) {
        if (!options.defaultLoaders) {
            throw new Error(
                'This plugin is not compatible with Next.js versions below 5.0.0 https://err.sh/next-plugins/upgrade'
            )
        }

        const Config = function(confNum) {
                if(confNum===-1) return config;
                config = nextConfig[confNum].webpack(config, options)
                return Config(confNum-1);
            }

        return Config(nextConfig.length-1)
    },
        webpackDevMiddleware: config => {
            // Perform customizations to webpack dev middleware config
            // Important: return the modified config
            return config;
        }
    }
}

module.exports =nextConfig([
    withCSS({
    cssModules: true,
    cssLoaderOptions: {
        camelCase: true,
        namedExport: true
    }}),
    withLess({
        lessLoaderOptions: {
            javascriptEnabled: true,
            modifyVars: themeVariables, // make your antd custom effective
        },
    }),
    withTypescript({
        useFileSystemPublicRoutes: false,
        webpack: function (config, { buildId, dev }) {
            const originalEntry = config.entry;
            config.resolve = {
                ...config.resolve,
                ...{
                    alias: {
                        ...config.resolve.alias,
                        '@src': path.resolve(__dirname, 'client'),
                    }
                },
            };

            return config
        }
    })
])
