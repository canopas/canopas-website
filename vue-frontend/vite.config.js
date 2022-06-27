import Pages from "vite-plugin-pages";
import { defineConfig } from "vite";
import vuePlugin from "@vitejs/plugin-vue";
import path from "path";

/**
 * @type {import('vite').UserConfig}
 */

export default defineConfig({
  ssr: {
    noExternal: [/^vue-meta*/, /^@fortawesome*/],
  },
  resolve: {
    alias: [
      {
        find: "@",
        replacement: path.resolve("./src"),
      },
    ],
  },
  settings: {
    "import/resolver": {
      alias: {
        map: [["@", "./src"]],
      },
    },
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: `@import "@/assets/scss/custom.scss";`,
      },
    },
  },
  plugins: [vuePlugin(), Pages()],
  build: {
    transpile: ["swiper"],
    rollupOptions: {
      output: {
        format: "es", // Transpile to ESM instead of CJS
      },
    },
  },
});
