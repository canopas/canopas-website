var path = require("path");
var webpack = require("webpack");
var HtmlWebpackPlugin = require("html-webpack-plugin");
const PrerenderSPAPlugin = require("@dreysolano/prerender-spa-plugin");
const { VueLoaderPlugin } = require("vue-loader");

module.exports = {
  entry: path.resolve(__dirname, "./src/main.js"),
  output: {
    path: path.resolve(__dirname, "./dist"),
    publicPath: "/",
    filename: "build.js",
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: "vue-loader",
      },
      {
        test: /\.js$/,
        loader: "babel-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.(jpe?g|png|gif|svg|webp)$/i,
        loader: "file-loader",
        options: {
          outputPath: "assets/images",
          name: "[name].[ext]?[hash]",
        },
      },
      {
        test: /\.(sa|sc|c)ss$/,
        use: [
          "vue-style-loader",
          "css-loader",
          {
            loader: "sass-loader",
            options: {
              implementation: require.resolve("sass"),
              additionalData: `
              @import "bootstrap/scss/functions";
              @import "bootstrap/scss/variables";
              @import "bootstrap/scss/mixins";
            `,
            },
          },
        ],
      },
    ],
  },
  resolve: {
    alias: {
      vue$: "vue/dist/vue.esm-bundler.js",
    },
    modules: [
      path.join(__dirname, "./src"),
      path.join(__dirname, "./node_modules"),
    ],
    extensions: [".js", ".jsx", ".vue", ".json", ".css", ".scss", ".svg"],
  },
  devServer: {
    historyApiFallback: true,
    noInfo: false,
  },
  devtool: "source-map",
  plugins: [
    new VueLoaderPlugin(),
    new webpack.DefinePlugin({
      "process.env": {
        NODE_ENV: '"production"',
      },
    }),
    new HtmlWebpackPlugin({
      title: "PRODUCTION prerender-spa-plugin",
      template: "public/index.html",
      filename: path.resolve(__dirname, "dist/index.html"),
      favicon: "public/favicon-16x16.webp",
    }),
    new PrerenderSPAPlugin({
      staticDir: path.join(__dirname, "dist"),
      routes: ["/", "/contact", "/jobs"],
      renderer: new PrerenderSPAPlugin.PuppeteerRenderer({
        headless: true,
      }),
    }),
  ],
};
