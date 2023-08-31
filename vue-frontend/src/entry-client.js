import { buildApp } from "./main";
import mixpanel from "mixpanel-browser";
import config from "@/config.js";
import "animate.css";
import { useContributionStore } from "./stores/contribution";
const { app, router, pinia } = buildApp(!1);
router.isReady().then(() => {
  var piniaInitialState = window.INITIAL_DATA;
  piniaInitialState && (pinia.state.value = piniaInitialState),
    mixpanel.init(config.MIX_PANEL_TOKEN),
    config.IS_PROD ||
      useContributionStore(pinia).fetchStars("fetchContributionStars"),
    app.provide("mixpanel", mixpanel).mount("#app");
});
