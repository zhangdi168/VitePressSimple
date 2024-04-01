export const defaultThemeConfig: any = {
  langMenuLabel: "多语言",
  returnToTopLabel: "回到顶部",
  sidebarMenuLabel: "菜单",
  darkModeSwitchLabel: "主题",
  lightModeSwitchTitle: "切换到浅色模式",
  darkModeSwitchTitle: "切换到深色模式",
  externalLinkIcon: true, //外部连接显示图标

  //编辑链接
  editLink: {
    text: "Edit this page on GitHub",
    pattern:
      "https://github.com/zhangdi168/VitePressSimple/edit/main/docs/:path",
  },
  //页脚
  docFooter: {
    prev: "Previous page",
    next: "Next Page",
  },
  nav: [], //导航
  //侧边栏
  sidebar: {},
  i18nRouting: true,
  //站点标题
  siteTitle: "VitePressSimple",
  //页脚配置
  footer: {
    message: "",
    copyright: "",
  },
};

export const defaultLangConfig: any = {
  themeConfig: defaultThemeConfig,
};
