import { defineConfig } from "vitepress";
	export default defineConfig({
  "title": "VitePressSimple",
  "description": "VitePressSimple文档",
  "srcDir": "./",
  "themeConfig": {
    "nav": [
      {
        "text": "Home",
        "link": "/"
      },
      {
        "text": "About",
        "link": "/about"
      },
      {
        "text": "Contact",
        "link": "/contact"
      }
    ],
    "sidebar": [
      {
        "text": "Examples",
        "items": [
          {
            "text": "Markdown Examples",
            "link": "/markdown-examples"
          },
          {
            "text": "Runtime API Examples",
            "link": "/api-examples"
          }
        ]
      }
    ],
    "socialLinks": [
      {
        "icon": "github",
        "link": "https://github.com/zhangdi168/VitePressSimple"
      }
    ],
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
    "docFooter": {
      "prev": "Previous page",
      "next": "Next Page"
    },
    "i18nRouting": true,
    "siteTitle": "VitePressSimple",
    "footer": {
      "message": "",
      "copyright": ""
    },
    "logo": "\\images\\logo.png",
    "outline": {
      "level": "1",
      "label": "On this page"
    },
    "search": {
      "provider": "local",
      "options": {
        "appId": "s11asa",
        "apiKey": "sasa",
        "indexName": "",
        "locales": {
          "zh": {
            "placeholder": "搜索文档",
            "translations": {
              "button": {
                "buttonText": "搜索文档",
                "buttonAriaLabel": "搜索文档"
              },
              "modal": {
                "searchBox": {
                  "resetButtonTitle": "清除查询条件",
                  "resetButtonAriaLabel": "清除查询条件",
                  "cancelButtonText": "取消",
                  "cancelButtonAriaLabel": "取消"
                },
                "startScreen": {
                  "recentSearchesTitle": "搜索历史",
                  "noRecentSearchesText": "没有搜索历史",
                  "saveRecentSearchButtonTitle": "保存至搜索历史",
                  "removeRecentSearchButtonTitle": "从搜索历史中移除",
                  "favoriteSearchesTitle": "收藏",
                  "removeFavoriteSearchButtonTitle": "从收藏中移除"
                },
                "errorScreen": {
                  "titleText": "无法获取结果",
                  "helpText": "你可能需要检查你的网络连接"
                },
                "footer": {
                  "selectText": "选择",
                  "navigateText": "切换",
                  "closeText": "关闭",
                  "searchByText": "搜索提供者"
                },
                "noResultsScreen": {
                  "noResultsText": "无法找到相关结果",
                  "suggestedQueryText": "你可以尝试查询",
                  "reportMissingResultsText": "你认为该查询应该有结果？",
                  "reportMissingResultsLinkText": "点击反馈"
                }
              }
            }
          },
          "en": {
            "placeholder": "Search docs",
            "translations": {
              "button": {
                "buttonText": "Search docs",
                "buttonAriaLabel": "Search docs"
              },
              "modal": {
                "searchBox": {
                  "resetButtonTitle": "Clear query",
                  "resetButtonAriaLabel": "Clear query",
                  "cancelButtonText": "Cancel",
                  "cancelButtonAriaLabel": "Cancel"
                },
                "startScreen": {
                  "recentSearchesTitle": "Recent searches",
                  "noRecentSearchesText": "No recent searches",
                  "saveRecentSearchButtonTitle": "Save search",
                  "removeRecentSearchButtonTitle": "Remove search",
                  "favoriteSearchesTitle": "Favorites",
                  "removeFavoriteSearchButtonTitle": "Remove from favorites"
                },
                "errorScreen": {
                  "titleText": "Unable to fetch results",
                  "helpText": "You may want to check your network connection"
                },
                "footer": {
                  "selectText": "select",
                  "navigateText": "navigate",
                  "closeText": "close",
                  "searchByText": "Search by"
                },
                "noResultsScreen": {
                  "noResultsText": "No results for",
                  "suggestedQueryText": "You can try searching for",
                  "reportMissingResultsText": "You think it should have results?",
                  "reportMissingResultsLinkText": "Click here to report it"
                }
              }
            }
          }
        }
      }
    }
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
          "pattern": "https://github.com/zhangdi168/VitePressSimple/edit/master/docs/:path"
        },
        "outline": {
          "level": "deep",
          "label": "目录"
        },
        "docFooter": {
          "prev": "Previous page",
          "next": "Next Page"
        },
        "nav": [
          {
            "text": "教程",
            "link": "/zh/使用手册/1.关于软件"
          },
          {
            "text": "更新记录",
            "link": "/zh/更新记录/v1.0.0"
          }
        ],
        "sidebar": {
          "/使用文档/": [
            {
              "text": "初体验",
              "link": "/zh//zh/使用文档/初体验"
            },
            {
              "text": "软件获取",
              "link": "/zh//zh/使用文档/软件获取"
            }
          ],
          "/zh/使用文档/": [
            {
              "text": "初体验",
              "link": "/zh/使用文档/初体验"
            },
            {
              "text": "软件获取",
              "link": "/zh/使用文档/软件获取"
            }
          ],
          "/zh/更新记录/": [
            {
              "text": "v1.0.0",
              "link": "/zh/更新记录/v1.0.0"
            }
          ],
          "/zh/使用手册/": [
            {
              "text": "1.关于软件",
              "link": "/zh/使用手册/1.关于软件"
            },
            {
              "text": "2.入门使用",
              "link": "/zh/使用手册/2.入门使用"
            },
            {
              "text": "3.项目配置",
              "link": "/zh/使用手册/3.项目配置"
            },
            {
              "text": "4.导航管理",
              "link": "/zh/使用手册/4.导航管理"
            },
            {
              "text": "5.侧栏管理",
              "link": "/zh/使用手册/5.侧栏管理"
            },
            {
              "text": "6.软件设置",
              "link": "/zh/使用手册/6.软件设置"
            }
          ]
        },
        "i18nRouting": true,
        "siteTitle": "VitePressSimple",
        "footer": {
          "message": "VitePressSimple",
          "copyright": "2024"
        }
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
        "docFooter": {
          "prev": "Previous page",
          "next": "Next Page"
        },
        "nav": [
          {
            "text": "tutorial",
            "link": "/en/使用手册/1.关于软件"
          },
          {
            "text": "versions",
            "link": "/en/更新记录/v1.0.0"
          }
        ],
        "sidebar": {
          "/en/使用手册/": [
            {
              "text": "1.about software",
              "link": "/en/使用手册/1.关于软件"
            },
            {
              "text": "2.Getting started",
              "link": "/en/使用手册/2.入门使用"
            },
            {
              "text": "3.project configuration",
              "link": "/en/使用手册/3.项目配置"
            },
            {
              "text": "4.navigation management",
              "link": "/en/使用手册/4.导航管理"
            },
            {
              "text": "5.sidebar management",
              "link": "/en/使用手册/5.侧栏管理"
            },
            {
              "text": "6.software settings",
              "link": "/en/使用手册/6.软件设置"
            }
          ],
          "/en/更新记录/": [
            {
              "text": "v1.0.0",
              "link": "/en/更新记录/v1.0.0"
            }
          ]
        },
        "i18nRouting": true,
        "siteTitle": "VitePressSimple",
        "footer": {
          "message": "VitePressSimple",
          "copyright": "2024"
        }
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
  },
  "sitemap": {
    "hostname": "http://vpsimple.xiaod.co/",
    "lastmodDateOnly": false
  }
});
