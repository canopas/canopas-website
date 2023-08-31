import { buildApp } from "./main";
import { renderToString } from "vue/server-renderer";
import { renderMetaToString } from "vue-meta/ssr/index.js";

async function render(url, manifest) {
  const { app, router, pinia } = buildApp(true);

  // Set the router to the desired URL before rendering
  router.push(url);
  await router.isReady();

  // Passing SSR context object which will be available via useSSRContext()
  const ctx = {};
  const html = await renderToString(app, ctx);
  await renderMetaToString(app, ctx);

  const renderState = `
  <script>
    window.INITIAL_DATA = ${JSON.stringify(pinia.state.value)}
  </script>`;

  // Generate preload links based on module IDs from the SSR context
  const preloadLinks = renderPreloadLinks(ctx.modules, manifest);

  // Return the rendered HTML, preload links, teleports, render state, and route name
  return [
    html,
    preloadLinks,
    ctx.teleports,
    renderState,
    router.currentRoute._value.name,
  ];
}

function renderPreloadLinks(modules, manifest) {
  let links = "";
  const seen = new Set();
  modules.forEach((id) => {
    const files = manifest[id];
    if (files) {
      files.forEach((file) => {
        if (!seen.has(file)) {
          seen.add(file);
          links += renderPreloadLink(file);
        }
      });
    }
  });
  return links;
}

function renderPreloadLink(file) {
  if (file.endsWith(".css")) {
    return `<link rel="preload" as="style" href="${file}" onload="this.rel='stylesheet'">`;
  } else if (file.includes("bg")) {
    return `<link rel="preload" as="image" href="${file}">`;
  } else if (file.includes(".woff2")) {
    return `<link rel="preload" as="font" href="${file}" type="font/woff2" crossorigin="anonymous">`;
  } else {
    return "";
  }
}

export { render };
