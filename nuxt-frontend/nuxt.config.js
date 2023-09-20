// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["@pinia/nuxt", "@nuxtjs/tailwindcss"],
  css: [
    "~/assets/css/global.css",
    "~/assets/css/app.css",
    "@fortawesome/fontawesome-svg-core/styles.css",
  ],
  generate: { fallback: true },
});
