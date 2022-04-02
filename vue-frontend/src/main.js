import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import VueGtag from "vue-gtag";
import config from "@/config.js";
import "@/assets/css/tailwind.css";

import "animate.css";

createApp(App)
  .use(router)
  .use(
    VueGtag,
    {
      config: {
        id: config.GAP_ID_ROOT,
      },
    },
    router
  )
  .mount("#app");
