const PrerenderSPAPlugin = require("prerender-spa-plugin");
var path = require("path");
var HtmlWebpackPlugin = require("html-webpack-plugin");

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
  configureWebpack: (config) => {
    return {
      devtool: "source-map",
      plugins: [
        new HtmlWebpackPlugin({
          title: "PRODUCTION prerender-spa-plugin",
          template: "public/index.html",
          filename: path.resolve(__dirname, "dist/[name].html"),
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
  },
};
