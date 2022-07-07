import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    modules: ['@nuxtjs/tailwindcss'],
    tailwindcss: {
        viewer: false,
        configPath: 'tailwind.config.cjs',
    },
    css: [
    ],
    vite: {
        css: {
            preprocessorOptions: {
              scss: {
                additionalData: `@import "@/assets/scss/custom.scss";`
              },
            },
          }
    },
})
