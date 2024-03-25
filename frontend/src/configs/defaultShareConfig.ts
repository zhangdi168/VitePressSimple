//公共配置的默认值
export const defaultShareConfigValue: any = {
  srcDir: "./docs",
  logo: "",
  assetsDir: "static",
  darkModeSwitchLabel: "Appearance",
  lightModeSwitchTitle: "Switch to light theme",
  darkModeSwitchTitle: "Switch to dark theme",
  sidebarMenuLabel: "Menu",
  returnToTopLabel: "Return to top",
  langMenuLabel: "Change language",
  externalLinkIcon: false,
  locales: {
    zh: {
      label: "简体中文",
      lang: "zh",
    },
    en: {
      label: "English",
      lang: "en",
    },
  },
  cacheDir: "./.vitepress/cache",
  title: "VitePressSimple",
  titleTemplate: "| simple config vitepress",
  description: "",
  lang: "en-US",
  base: "/",
  outDir: "./.vitepress/dist",
  map: false,
  cleanUrls: false,
  socialLinks: [{ icon: "github", link: "https://github.com/vuejs/vitepress" }],
  rewrites: {
    "packages/pkg-a/src/pkg-a-docs.md": "pkg-a/index.md",
    "packages/pkg-b/src/pkg-b-docs.md": "pkg-b/index.md",
  },
};
