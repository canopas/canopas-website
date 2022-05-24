// @ts-check
const fs = require("fs");
const path = require("path");
const express = require("express");
const vite = require("vite");
const serveStatic = require("serve-static");
const compression = require("compression");

const isTest = process.env.NODE_ENV === "test" || !!process.env.VITE_TEST_BUILD;

async function createServer(
  root = process.cwd(),
  isProd = process.env.NODE_ENV === "production"
) {
  const resolve = (p) => path.resolve(__dirname, p);

  const indexProd = isProd
    ? fs.readFileSync(resolve("dist/client/index.html"), "utf-8")
    : "";

  const manifest = isProd
    ? // @ts-ignore
      JSON.parse(fs.readFileSync(resolve("dist/client/ssr-manifest.json")))
    : {};

  const app = express();

  /**
   * @type {import('vite').ViteDevServer}
   */
  let viteServer;
  if (!isProd) {
    viteServer = await vite.createServer({
      root,
      logLevel: isTest ? "error" : "info",
      server: {
        middlewareMode: "ssr",
        watch: {
          // During tests we edit the files too fast and sometimes chokidar
          // misses change events, so enforce polling for consistency
          usePolling: true,
          interval: 100,
        },
      },
    });
    // use vite's connect instance as middleware
    app.use(viteServer.middlewares);
  } else {
    app.use(compression());
    app.use(
      serveStatic(resolve("dist/client"), {
        index: false,
      })
    );
  }

  app.use("*", async (req, res) => {
    try {
      const url = req.originalUrl;

      let template, render;
      if (!isProd) {
        // always read fresh template in dev
        template = fs.readFileSync(resolve("index.html"), "utf-8");
        template = await viteServer.transformIndexHtml(url, template);
        render = (await viteServer.ssrLoadModule("/src/entry-server.js"))
          .render;
      } else {
        template = indexProd;
        render = (await import("./dist/server/entry-server.js")).render;
      }

      var [appHtml, preloadLinks, teleports, renderState] = await render(
        url,
        manifest
      );

      const html = template
        .replace(`<!--tele-ports-->`, teleports.head)
        .replace(`<!--preload-links-->`, preloadLinks)
        .replace(`<!--app-store-->`, renderState)
        .replace(`<!--app-html-->`, appHtml);

      res.status(200).set({ "Content-Type": "text/html" }).end(html);
    } catch (e) {
      viteServer && viteServer.ssrFixStacktrace(e);
      console.log(e.stack);
      res.status(500).end(e.stack);
    }
  });

  return { app, viteServer };
}

if (!isTest) {
  createServer().then(({ app }) =>
    app.listen(3080, () => {
      console.log("http://localhost:3080");
    })
  );
}
