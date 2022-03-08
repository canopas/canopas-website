var HtmlWebpackPlugin = require("html-webpack-plugin");
const PrerenderSPAPlugin = require("prerender-spa-plugin");
var webpack = require("webpack");
var path = require("path");

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
    devtool: "source-map",
    plugins: [
      new webpack.DefinePlugin({
        "process.env": {
          NODE_ENV: '"production"',
        },
      }),
      new HtmlWebpackPlugin({
        title: "PRODUCTION prerender-spa-plugin",
        template: "public/index.html",
        filename: path.resolve(__dirname, "dist/[name].html"),
        favicon: "public/favicon-16x16.webp",
      }),
      new PrerenderSPAPlugin({
        staticDir: path.join(__dirname, "dist"),
        routes: ["/"],
        renderer: new PrerenderSPAPlugin.PuppeteerRenderer({
          maxConcurrentRoutes: 1,
          headless: true,
          navigationOptions: {
            waitUntil: "domcontentloaded",
          },
          renderAfterElementExists: "#app",
        }),
      }),
    ],
  },
};
