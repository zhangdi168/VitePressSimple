//公共配置的默认值
export const defaultShareConfigValue: any = {
  srcDir: "./",
  // logo: "",
  assetsDir: "static",
  //生成sitemap的配置
  sitemap: {
    hostname: "https://example.com",
    lastmodDateOnly: false,
  },
  // darkModeSwitchLabel: "Appearance",
  // lightModeSwitchTitle: "Switch to light theme",
  // darkModeSwitchTitle: "Switch to dark theme",
  // sidebarMenuLabel: "Menu",
  // returnToTopLabel: "Return to top",
  // langMenuLabel: "Change language",
  themeConfig: {
    externalLinkIcon: false,
    i18nRouting: true,
    logo: "",
    socialLinks: [
      { icon: "github", link: "https://github.com/vuejs/vitepress" },
    ],
    outline: {
      level: 2,
      label: "On this page",
    },
    search: {
      provider: "local", //搜索类型
      options: {
        appId: "",
        apiKey: "",
        indexName: "",
        locales: {
          zh: {
            placeholder: "搜索文档",
            translations: {
              button: {
                buttonText: "搜索文档",
                buttonAriaLabel: "搜索文档",
              },
              modal: {
                searchBox: {
                  resetButtonTitle: "清除查询条件",
                  resetButtonAriaLabel: "清除查询条件",
                  cancelButtonText: "取消",
                  cancelButtonAriaLabel: "取消",
                },
                startScreen: {
                  recentSearchesTitle: "搜索历史",
                  noRecentSearchesText: "没有搜索历史",
                  saveRecentSearchButtonTitle: "保存至搜索历史",
                  removeRecentSearchButtonTitle: "从搜索历史中移除",
                  favoriteSearchesTitle: "收藏",
                  removeFavoriteSearchButtonTitle: "从收藏中移除",
                },
                errorScreen: {
                  titleText: "无法获取结果",
                  helpText: "你可能需要检查你的网络连接",
                },
                footer: {
                  selectText: "选择",
                  navigateText: "切换",
                  closeText: "关闭",
                  searchByText: "搜索提供者",
                },
                noResultsScreen: {
                  noResultsText: "无法找到相关结果",
                  suggestedQueryText: "你可以尝试查询",
                  reportMissingResultsText: "你认为该查询应该有结果？",
                  reportMissingResultsLinkText: "点击反馈",
                },
              },
            },
          },
          en: {
            //英文
            placeholder: "Search docs",
            translations: {
              button: {
                buttonText: "Search docs",
                buttonAriaLabel: "Search docs",
              },
              modal: {
                searchBox: {
                  resetButtonTitle: "Clear query",
                  resetButtonAriaLabel: "Clear query",
                  cancelButtonText: "Cancel",
                  cancelButtonAriaLabel: "Cancel",
                },
                startScreen: {
                  recentSearchesTitle: "Recent searches",
                  noRecentSearchesText: "No recent searches",
                  saveRecentSearchButtonTitle: "Save search",
                  removeRecentSearchButtonTitle: "Remove search",
                  favoriteSearchesTitle: "Favorites",
                  removeFavoriteSearchButtonTitle: "Remove from favorites",
                },
                errorScreen: {
                  titleText: "Unable to fetch results",
                  helpText: "You may want to check your network connection",
                },
                footer: {
                  selectText: "select",
                  navigateText: "navigate",
                  closeText: "close",
                  searchByText: "Search by",
                },
                noResultsScreen: {
                  noResultsText: "No results for",
                  suggestedQueryText: "You can try searching for",
                  reportMissingResultsText: "You think it should have results?",
                  reportMissingResultsLinkText: "Click here to report it",
                },
              },
            },
          },
        },
      },
    },
  },
  locales: {
    // zh: {
    //   label: "简体中文",
    //   lang: "zh",
    // },
    // en: {
    //   label: "English",
    //   lang: "en",
    // },
  },
  cacheDir: "./.vitepress/cache",
  title: "VitePressSimple",
  titleTemplate: "| simple config vitepress",
  description: "",
  lang: "en-US",
  base: "/",
  outDir: "./.vitepress/dist",
  cleanUrls: false,
  rewrites: {
    "packages/pkg-a/src/pkg-a-docs.md": "pkg-a/index.md",
    "packages/pkg-b/src/pkg-b-docs.md": "pkg-b/index.md",
  },
};
