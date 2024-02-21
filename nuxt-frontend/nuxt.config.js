// https://nuxt.com/docs/api/configuration/nuxt-config
import { defineNuxtConfig } from "nuxt/config";
import config from "./config";

export default defineNuxtConfig({
  app: {
    head: {
      htmlAttrs: {
        lang: "en",
      },
    },
    cdnURL: config.CLOUDFRONT_URL,
  },
  devtools: { enabled: true },
  modules: [
    "@pinia/nuxt",
    "@nuxtjs/tailwindcss",
    "nuxt-icon",
    "@canopassoftware/blog-components",
    "nuxt-lazy-hydrate",
    "@nuxtjs/sitemap",
  ],
  site: {
    url: config.BASE_URL,
  },
  sitemap: {
    defaults: {
      lastmod: new Date(),
    },
    sources: [config.API_BASE + "/api/sitemap"],
    xsl: false,
    xslTips: false,
  },
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
  routeRules: {
    "/blog": { isr: true },
    "/resources": { isr: true },
    "/tag/**": { isr: true },
    "/author/**": { isr: true },
    "/jobs/**": { isr: true },
    "/": { prerender: true },
    "/services": { prerender: true },
    "/portfolio": { prerender: true },
    "/portfolio/**": { prerender: true },
    "/contributions": { prerender: true },
    "/about": { prerender: true },
    "/contact": { prerender: true },
    "/android-app-development": { prerender: true },
    "/ios-app-development": { prerender: true },
    "/mobile-app-development": { prerender: true },
    "/backend-development": { prerender: true, sitemap: !config.IS_PROD },
    "/flutter-app-development": { prerender: true, sitemap: !config.IS_PROD },
    "/frontend-development": { prerender: true, sitemap: !config.IS_PROD },
    "/thank-you": { prerender: true },
    "/unsubscribe": { prerender: true },
  },
  nitro: {
    compressPublicAssets: true,
    preset: "aws-lambda",
  },
});
