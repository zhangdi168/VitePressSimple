import { defineConfig } from "vitepress";
import { vpSimpleNav } from "./vpsimple/nav";
import { vpSimpleSidebar } from "./vpsimple/sidebar";
import { vpSimpleSocialLinks } from "./vpsimple/socialLinks";

export default defineConfig({
  title: "vitepress-demo",
  description: "A VitePress Site",
  themeConfig: {
    srcDir: './docs',
    nav: vpSimpleNav,//导航
    sidebar: vpSimpleSidebar,//侧栏
    // @ts-ignore 社交账号
    socialLinks: vpSimpleSocialLinks,
  }
})
