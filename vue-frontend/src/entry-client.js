import { createApp } from "./main";
import VueGtag from "vue-gtag";
import config from "@/config.js";
import "animate.css";

const { app, router, pinia } = createApp(false);

// wait until router is ready before mounting to ensure hydration match
router.isReady().then(() => {
  const piniaInitialState = window.INITIAL_DATA;
  if (piniaInitialState) {
    pinia.state.value = piniaInitialState;
  }

  app
    .use(VueGtag, {
      config: { id: config.GAP_ID_ROOT },
    })
    .mount("#app");
});
