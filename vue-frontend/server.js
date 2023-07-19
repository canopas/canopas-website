// @ts-nocheck
import fs from "fs";
import path from "path";
import express from "express";
import { createServer as createViteServer } from "vite";
import serveStatic from "serve-static";
import compression from "compression";
import Cache from "./utils/cache.js";

const __dirname = path.resolve();

const cache = new Cache();

const isTest = process.env.NODE_ENV === "test" || !!process.env.VITE_TEST_BUILD;

async function createServer(
  root = process.cwd(),
  isProd = process.env.NODE_ENV === "production",
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
    viteServer = await createViteServer({
      root,
      logLevel: isTest ? "error" : "info",
      server: {
        middlewareMode: "true",
        appType: "custom",
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
        immutable: true,
        maxAge: "365d",
      }),
    );
  }

  app.use("*", async (req, res) => {
    try {
      const url = req.originalUrl;

      var html = cache.get(url);
      var statusCode = 200;

      if (html == null || !isProd) {
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

        var [appHtml, preloadLinks, teleports, renderState, currentRouteName] =
          await render(url, manifest);

        html = template
          .replace(`<!--tele-ports-->`, teleports.head)
          .replace(`<!--preload-links-->`, preloadLinks)
          .replace(`<!--app-store-->`, renderState)
          .replace(`<!--app-html-->`, appHtml);

        statusCode = currentRouteName.includes("404") ? 404 : statusCode;
        cache.put(url, html);
      }

      res.status(statusCode).set({ "Content-Type": "text/html" }).end(html);
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
    }),
  );
}

// for test use
export default createServer;
