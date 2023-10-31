/* plugins/mixpanel.js */

import { defineNuxtPlugin } from "#app";
import mixpanel from "mixpanel-browser";
import config from "../config.js";

export default defineNuxtPlugin((nuxtApp) => {
  mixpanel.init(config.MIX_PANEL_TOKEN);
  nuxtApp.provide("mixpanel", mixpanel);
});
