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

module.exports =() => {
   return  Object.assign(
       withTypescript(
        withCSS({
            cssModules: true,
            cssLoaderOptions: {
                camelCase: true,
                namedExport: true
            }
        }),
        withLess({
            lessLoaderOptions: {
                javascriptEnabled: true,
                modifyVars: themeVariables, // make your antd custom effective
            },
        }),
        {
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
        }));
}
