import { buildApp } from "./main";
import mixpanel from "mixpanel-browser";
import config from "@/config.js";
import "animate.css";
import { useContributionStore } from "./stores/contribution";

const { app, router, pinia } = buildApp(false);

// wait until router is ready before mounting to ensure hydration match
router.isReady().then(() => {
  const piniaInitialState = window.INITIAL_DATA;
  if (piniaInitialState) {
    pinia.state.value = piniaInitialState;
  }

  mixpanel.init(config.MIX_PANEL_TOKEN);

  if (!config.IS_PROD) {
    useContributionStore(pinia).fetchStars("fetchContributionStars");
  }

  app.provide("mixpanel", mixpanel).mount("#app");
});
