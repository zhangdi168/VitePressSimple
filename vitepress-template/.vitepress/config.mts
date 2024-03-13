import { defineConfig } from "vitepress";

export default defineConfig({
  title: "vitepress-demo",
  description: "A VitePress Site",
  srcDir: "./docs",
  themeConfig: {
    //导航
    nav: [
      {
        text: "Home",
        link: "/"
      },
      {
        text: "About",
        link: "/about"
      },
      {
        text: "Contact",
        link: "/contact"
      }
    ],
    //侧栏
    sidebar: [
      {
        text: "Examples",
        items: [
          { text: "Markdown Examples", link: "/markdown-examples" },
          { text: "Runtime API Examples", link: "/api-examples" }
        ]
      }
    ],
    //社交账号
    socialLinks: [
      { icon: "github", link: "https://github.com/vuejs/vitepress" }
    ]
  }
});
