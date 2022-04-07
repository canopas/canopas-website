import { createApp } from "./main";
import VueGtag from "vue-gtag";
import config from "@/config.js";
import "animate.css";

const { app, router, store } = createApp(false);

// wait until router is ready before mounting to ensure hydration match
router.isReady().then(() => {
  const storeInitialState = window.INITIAL_DATA;
  if (storeInitialState) {
    store.replaceState(storeInitialState);
  }

  app
    .use(VueGtag, {
      config: { id: config.GAP_ID_ROOT },
    })
    .mount("#app");
});
