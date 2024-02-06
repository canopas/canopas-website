// https://nuxt.com/docs/api/configuration/nuxt-config
import { defineNuxtConfig } from "nuxt/config";
export default defineNuxtConfig({
  app: {
    head: {
      htmlAttrs: {
        lang: "en",
      },
    },
  },
  devtools: { enabled: true },
  modules: [
    "@pinia/nuxt",
    "@nuxtjs/tailwindcss",
    "nuxt-icon",
    "@canopassoftware/blog-components",
  ],
  css: ["~/assets/css/global.css", "~/assets/css/app.css"],
  generate: { fallback: true },
  imports: {
    dirs: ["stores"],
  },
  plugins: [{ src: "~/plugins/mixpanel" }],
  hooks: {
    "pages:extend"(pages) {
      pages.push({
        name: "Error404Page",
        path: "/:pathMatch(.*)*",
        file: "~/error.vue",
      });
    },
  },
  nitro: {
    prerender: {
      failOnError: false,
    },
  },
});
