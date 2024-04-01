import { defineConfig } from "vitepress";

export default defineConfig({
  "title": "VitePressSimple",
  "description": "VitePressSimple文档",
  "srcDir": "./",
  "themeConfig": {
    "nav": [{ "text": "Home", "link": "/" }, { "text": "About", "link": "/about" }, {
      "text": "Contact",
      "link": "/contact"
    }],
    "sidebar": [{
      "text": "Examples",
      "items": [{ "text": "Markdown Examples", "link": "/markdown-examples" }, {
        "text": "Runtime API Examples",
        "link": "/api-examples"
      }]
    }],
    "socialLinks": [{ "icon": "github", "link": "https://github.com/vuejs/vitepress" }],
    "langMenuLabel": "多语言",
    "returnToTopLabel": "回到顶部",
    "sidebarMenuLabel": "菜单",
    "darkModeSwitchLabel": "主题",
    "lightModeSwitchTitle": "切换到浅色模式",
    "darkModeSwitchTitle": "切换到深色模式",
    "externalLinkIcon": true,
    "editLink": {
      "text": "Edit this page on GitHub",
      "pattern": "https://github.com/zhangdi168/VitePressSimple/edit/main/docs/:path"
    },
    "docFooter": { "prev": "Previous page", "next": "Next Page" },
    "i18nRouting": true,
    "siteTitle": "VitePressSimple",
    "footer": { "message": "", "copyright": "" },
    "logo": "\\images\\logo.png"
  },
  "assetsDir": "static",
  "locales": {
    "zh": {
      "label": "简体中文",
      "lang": "zh",
      "themeConfig": {
        "langMenuLabel": "多语言",
        "returnToTopLabel": "回到顶部",
        "sidebarMenuLabel": "菜单",
        "darkModeSwitchLabel": "主题",
        "lightModeSwitchTitle": "切换到浅色模式",
        "darkModeSwitchTitle": "切换到深色模式",
        "externalLinkIcon": true,
        "editLink": {
          "text": "Edit this page on GitHub",
          "pattern": "https://github.com/zhangdi168/VitePressSimple/edit/main/docs/:path"
        },
        "docFooter": { "prev": "Previous page", "next": "Next Page" },
        "nav": [{ "text": "教程", "link": "" }, { "text": "更新记录", "link": "/zh/更新记录" }],
        "sidebar": [],
        "i18nRouting": true,
        "siteTitle": "VitePressSimple",
        "footer": { "message": "VitePressSimple", "copyright": "2024" }
      }
    },
    "en": {
      "label": "English",
      "lang": "en",
      "themeConfig": {
        "langMenuLabel": "多语言",
        "returnToTopLabel": "回到顶部",
        "sidebarMenuLabel": "菜单",
        "darkModeSwitchLabel": "主题",
        "lightModeSwitchTitle": "切换到浅色模式",
        "darkModeSwitchTitle": "切换到深色模式",
        "externalLinkIcon": true,
        "editLink": {
          "text": "Edit this page on GitHub",
          "pattern": "https://github.com/zhangdi168/VitePressSimple/edit/main/docs/:path"
        },
        "docFooter": { "prev": "Previous page", "next": "Next Page" },
        "nav": [{ "text": "教程", "link": "" }, { "text": "更新记录", "link": "" }],
        "sidebar": {},
        "i18nRouting": true,
        "siteTitle": "VitePressSimple",
        "footer": { "message": "", "copyright": "" }
      }
    }
  },
  "cacheDir": "./.vitepress/cache",
  "titleTemplate": "| simple config vitepress",
  "lang": "en-US",
  "base": "/",
  "outDir": "./.vitepress/dist",
  "cleanUrls": false,
  "rewrites": {
    "packages/pkg-a/src/pkg-a-docs.md": "pkg-a/index.md",
    "packages/pkg-b/src/pkg-b-docs.md": "pkg-b/index.md"
  }
});
