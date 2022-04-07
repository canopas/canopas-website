const virtualFile = "@virtual-file";
const virtualId = "\0" + virtualFile;
const nestedVirtualFile = "@nested-virtual-file";
const nestedVirtualId = "\0" + nestedVirtualFile;
import Pages from "vite-plugin-pages";
import { defineConfig } from "vite";
import { promises as fs } from "fs";
import vuePlugin from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
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
  plugins: [
    vuePlugin(),
    Pages(),
    vueJsx(),
    {
      name: "fix-swipper-css",
      enforce: "pre",
      resolveId(id) {
        if (id === "swiper.css") return "fix-swiper.css";
      },
      async load(id) {
        if (id === "fix-swiper.css") {
          return await fs.readFile(
            "./node_modules/swiper/swiper.min.css",
            "utf-8"
          );
        }
      },
    },
    {
      name: "fix-swipper-pagination-css",
      enforce: "pre",
      resolveId(id) {
        if (id === "swiper.css.pagination") return "fix-swiper-pagination.css";
      },
      async load(id) {
        if (id === "fix-swiper-pagination.css") {
          return await fs.readFile(
            "./node_modules/swiper/modules/pagination/pagination.min.css",
            "utf-8"
          );
        }
      },
    },
    {
      name: "virtual",
      resolveId(id) {
        if (id === "@foo") {
          return id;
        }
      },
      load(id) {
        if (id === "@foo") {
          return `export default { msg: 'hi' }`;
        }
      },
    },
    {
      name: "virtual-module",
      resolveId(id) {
        if (id === virtualFile) {
          return virtualId;
        } else if (id === nestedVirtualFile) {
          return nestedVirtualId;
        }
      },
      load(id) {
        if (id === virtualId) {
          return `export { msg } from "@nested-virtual-file";`;
        } else if (id === nestedVirtualId) {
          return `export const msg = "[success] from conventional virtual file"`;
        }
      },
    },
  ],
  build: {
    minify: false,
    rollupOptions: {
      output: {
        format: "es", // Transpile to ESM instead of CJS
      },
    },
  },
});
