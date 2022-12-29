import App from "./App.vue";
import { createSSRApp, createApp } from "vue";
import { createRouter } from "./router";
import { createPinia } from "pinia";
import { createMetaManager } from "vue-meta";
import "@/assets/css/tailwind.css";

export function buildApp(isSSR) {
  const app = isSSR ? createSSRApp(App) : createApp(App);
  const router = createRouter();
  const pinia = createPinia();
  const metaManager = createMetaManager(isSSR);

  app.use(router).use(metaManager).use(pinia);

  return { app, router, pinia, metaManager };
}
