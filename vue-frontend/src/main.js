import App from "./App.vue";
import { createSSRApp } from "vue";
import { createRouter } from "./router";
import store from "./store";
import { createMetaManager } from "vue-meta";

export function createApp(isSSR) {
  const app = createSSRApp(App);
  const router = createRouter();
  const metaManager = createMetaManager(isSSR);
  app.use(router).use(store).use(metaManager);
  return { app, router, store, metaManager };
}
