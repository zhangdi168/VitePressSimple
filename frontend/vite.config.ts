import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";
import { quasar, transformAssetUrls } from "@quasar/vite-plugin";
import { resolve } from "path";
import { nodePolyfills } from "vite-plugin-node-polyfills";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue({ template: { transformAssetUrls } }),
    vueJsx(),
    nodePolyfills(),
    quasar({
      sassVariables: "src/quasar-variables.sass",
    }),
  ],
  resolve: {
    alias: {
      "@": resolve(__dirname, "src"),
      "~": resolve(__dirname, "wails"),
    },
  },
  build: {
    rollupOptions: {
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`,
      },
    },
  },
});
