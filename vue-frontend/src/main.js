import App from "./App.vue";
import { createSSRApp, createApp } from "vue";
import { createRouter } from "./router";
import { createPinia } from "pinia";
import { createMetaManager } from "vue-meta";
import "@/assets/css/tailwind.css";

function buildApp(isSSR) {
  var appInstance = (isSSR ? createSSRApp : createApp)(App),
    routerInstance = createRouter(),
    piniaInstance = createPinia(),
    metaManagerInstance = createMetaManager(isSSR);

  appInstance.use(routerInstance).use(metaManagerInstance).use(piniaInstance);

  return {
    app: appInstance,
    router: routerInstance,
    pinia: piniaInstance,
    metaManager: metaManagerInstance,
  };
}

export { buildApp };
