const withTypescript = require('@zeit/next-typescript');
const path = require('path');

module.exports = withTypescript({
  useFileSystemPublicRoutes: false,
  webpack: function (config, { buildId, dev }) {
    const originalEntry = config.entry;
    config.module.rules.push({
      test: /\.css$/,
      loader:'style-loader!css-loader'
    },{
      test: /\.less$/,
      use: [
        'style-loader',
        'css-loader',
        'less-loader'
      ],
      include: /node_modules/,
    })
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
});
