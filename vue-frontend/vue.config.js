var HtmlWebpackPlugin = require('html-webpack-plugin')
const PrerenderSPAPlugin = require('prerender-spa-plugin')
const Renderer = PrerenderSPAPlugin.PuppeteerRenderer
var webpack = require('webpack')
var path = require('path')

module.exports = {
  css: {
    loaderOptions: {
      sass: {
        additionalData: `
          @import "@/assets/scss/custom.scss";
        `,
      },
    },
  }, 
  configureWebpack: {
    devtool: 'source-map',
    plugins: [new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: '"production"'
      }
    }),
    new HtmlWebpackPlugin({
      title: 'PRODUCTION prerender-spa-plugin',
      template: 'public/index.html',
      filename: path.resolve(__dirname, 'dist/[name].html'),
      favicon: 'public/favicon-16x16.webp'
    }),
    new PrerenderSPAPlugin({
      staticDir: path.join(__dirname, 'dist'),
      routes: ['/', '/contact'],
      renderer: new Renderer()
    })]
  }
};

// if (process.env.NODE_ENV === 'development') {
//   module.exports.devtool = '#source-map'
//   module.exports.plugins = (module.exports.plugins || []).concat([
//     new webpack.DefinePlugin({
//       'process.env': {
//         NODE_ENV: '"production"'
//       }
//     }),
//     new HtmlWebpackPlugin({
//       title: 'PRODUCTION prerender-spa-plugin',
//       template: 'index.html',
//       filename: path.resolve(__dirname, 'dist/index.html'),
//       favicon: 'favicon.ico'
//     }),
//     new PrerenderSPAPlugin({
//       staticDir: path.join(__dirname, 'dist'),
//       routes: ['/', '/jobs', '/contact'],

//       renderer: new Renderer({
//         headless: true,
//         renderAfterDocumentEvent: 'render-event'
//       })
//     })
//   ])
// } else {
//   // NODE_ENV === 'development'
//   module.exports.configureWebpack.plugins = (module.exports.plugins || []).concat([
//     new webpack.DefinePlugin({
//       'process.env': {
//         NODE_ENV: '"development"'
//       }
//     }),
//     new HtmlWebpackPlugin({
//       title: 'DEVELOPMENT prerender-spa-plugin',
//       template: 'index.html',
//       filename: 'index.html',
//       favicon: 'favicon.ico'
//     }),
//   ])
// }
