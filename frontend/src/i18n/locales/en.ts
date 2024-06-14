export const en = {
  nav: {
    home: "Home",
    setting: "Settings",
    config: "Project Config",
    about: "About",
    preview: "Preview & Build",
    nav: "Navigation Settings",
    sidebar: "Sidebar Settings",
  },
  languages: {
    en: "English",
    "zh-Hans": "简体中文",
    fr: "French",
  },
  topbar: {
    minimise: "Minimize",
    quit: "Quit",
  },
  common: {
    confirm: "Confirm",
    know: "Acknowledged",
    cancel: "Cancel",
    confirmDelete: "Confirm Deletion",
    deleteSuccess: "Deleted Successfully",
    hide: "Hide",
    show: "Show",
    saveWithKey: "Save (Ctrl+S)",
    base: "Basic",
    home: "Home",
    code: "Code",
    title: "Title",
    noOpenProject: "No project open. Create or open one first.",
    create: "Create",
    name: "Name",
    description: "Description",
    open: "Open",
    saveConfig: "Save Configuration",
    terminal: "terminal",
  },
  pageIndex: {
    newProject: "New Project",
    newArticle: "New Article",
    openProject: "Open Project",
    newDirectory: "New Directory",
    rename: "Rename",
    delete: "Delete",
    copy: "Copy",
    paste: "Paste",
    copyRoute: "Copy Route",
    openExistingProject: "Open Existing Project",
    openInFileBrowser: "Open in File Browser",
    refreshDirectoryTree: "Refresh Directory Tree",
    inputFolderName: "Enter Folder Name",
    inputArticleTitle: "Enter Article Title",
    folderCreated: "Folder Created",
    articleCreated: "Article Created",
    inputArticleTitleWithoutMdSuffix: "Enter Article Title without .md suffix",
    inputNewName: "Enter New Name",
    confirmNewPath: "Confirm New Path",
    confirmOrModifyNewPath: "Confirm or Modify New Path",
    path: "Path",
    copied: "Copied",
    copyFailed: "Copy Failed",
    noCopyPath: "No Path to Copy",
    directory: "Directory",
    article: "Article",
    pathExists: "Path Exists:",
    pathNotExists: "Path Does Not Exist:",
    copiedSuccess: "Successfully Copied",
    newFileInRoot: "New File (Root)",
    newFolderInRoot: "New Folder (Root)",
    currentProject: "Current Project:",
    pageProperties: "Page Property Editor",
    inputPageSeoDescription: "Enter Page SEO Description",
    inputOutlineLevel: "Optional: 0-6, 0 hides the outline",
    outlineLevel: "Outline Display Level (0-6)",
    showNav: "Display Navigation",
    showSidebar: "Display Sidebar",
    showFooter: "Display Footer",
    showEditLink: "Display Edit Link",
    showUpdateTime: "Display Update Time",
    outlinePosition: "Outline Position",
    outlineLeft: "Outline on Left",
    outlineRight: "Outline on Right",
    addMeta: "Add Meta",
    customFormatterTip:
      "Variable data is stored in 'custom' attribute, accessed via {frontmatter}.custom.{xx}",
    addCustomFormatter: "Add Custom Front Matter",
    pageType: "Page Type",
    pageTypeDoc: "Document",
    pageTypeHome: "Homepage",
    pageTypeCustom: "Custom",
    hero: {
      name: "Name",
      text: "Introduction",
      tagline: "Slogan",
      taglineExtra: "Slogan below `text`",
      image: {
        select: "Select Homepage Image",
        tips: "Homepage images default to ./images/home/{filename}",
        src: "Image URL",
        alt: "Image Description",
      },
      actions: {
        addBtnText: "Add Homepage Button",
        valuePlaceholder: "Link Destination",
        keyPlaceholder: "Text",
      },
      features: {
        addBtnText: "Add Features",
        valuePlaceholder: "Feature Details",
        keyPlaceholder: "Feature Title",
      },
    },
    useExampleVueCode: "Use Example Vue Code",
    noSelectedArticle: "No Article Selected",
  },
  pageNav: {
    selectOperationLang: "Select Operation Language",
    copy: "Copy",
    cut: "Cut",
    paste: "Paste",
    clipboard: "to Clipboard",
    coverCurrentNav: "Paste Clipboard Nav Data to Override Current Lang Nav",
    addTopNav: "Add Top-Level Navigation",
    saveCurrentNav: "Save Current Navigation",
    toCurrentNav: "To Current Navigation",
  },
  pageSidebar: {
    emptyProject: "Empty Project",
    languageSelection: "Language Selection",
    switchThemesTip: "Click to Switch to",
    confirmSwitchSidebar:
      "Switching will [Clear] the current language's sidebar data",
    confirmSwitchButtonOk: "Understood",
    confirmSwitchButtonCancel: "Nevermind",
    chooseEditedSidebar: "Choose the Currently Edited Sidebar:",
    whenRouteInTip: "When route is in",
    whenRouteInTip2: "directory, this sidebar will display",
    reloadSidebarTip: "Reload Sidebar Directory",
    addTopSidebar: "Add Top-Level Sidebar",
    autoRecognizeSidebar: "Auto-Detect Current Sidebar",
    recognitionWarning:
      "Detected sidebar data will overwrite current sidebar data",
    saveSidebar: "Save Current Sidebar",
    noSubDirectory: "No Subdirectory",
    autoRecognize: "Auto-Detect",
    sidebarData: "Sidebar Data Under Directory",
    current: "Current",
    sidebar: "Sidebar",
    singleSidebar: "Single Sidebar",
    multiSidebar: "Multiple Sidebars",
    mode: "Mode",
    noSubDirectoryTip:
      "No subdirectories found under this path, please create first!",
    selectSidebarTip: "Please select the sidebar you are currently editing",
    recognitionWarning2: "Detected sidebar data is empty, ensure",
    existFile: "files exist in the folder",
    existFolder: "folders exist in the folder",
  },
  pageProject: {
    newVitePressProject: "New VitePress Project",
    basedOn: "Based on VitePress Version:",
    enterProjectName: "Enter Project Name",
    projectNameHint: "Project name, follow file naming rules",
    enterProjectDescription: "Enter Project Description",
    projectDescriptionHint: "Project description, preferably in Chinese",
    projectRootDirectory:
      "Project root where .vitepress special directory is sought",
    chooseRootDirectory: "Choose Root Directory",
    sourceDirectory: "Source Directory (Docs)",
    sourceDirectoryHint:
      "Source directory for Markdown files, default is ./docs, can be changed after creation",
    errorProjectNameEmpty: "Please enter a project name first",
    errorProjectRootEmpty: "Please select a project directory first",
    openProjectFromDirectory: "Open Project from Directory",
    errorProjectNotVPSimple:
      "If the project wasn't created by VPSimple, config files may fail to load",
    recentProject: "Recent Projects",
    settingBase: {
      title: "Basic Settings",
      selectLogo: "Select Site Logo",
      logoSavePath: "Logo image defaults to ./images/{logo}",
      tooltips: {
        pageTitle: "The title used in config.title",
        titleSuffix:
          "Allows customization of page title suffix or the entire title",
        outlineLevel: "Outline display level",
        htmlLang:
          "The lang attribute of the site. This will appear as the <html lang='xx'> tag in page HTML",
        siteDescription: "Site description, used in <meta> tags in page HTML",
        externalLinkIcon:
          "Whether to display an external link icon next to external links in markdown",
        docPath:
          "Markdown document directory. Ensure image resources are accessible; do not change arbitrarily after setup",
        baseUrl:
          "Base URL the site will be deployed at. If deploying to a subpath, set this",
        buildPath:
          "Project build output location, relative to the project root",
        hostname: "Used when generating a sitemap.xml file for the site",
        staticAssets:
          "Specifies the directory to place generated static assets. The path should be within the build path and resolved relative to it",
        cachePath: "Cache directory, relative to the project root",
        cleanUrls: "When true, VitePress removes the .html suffix from URLs",
      },
      labels: {
        pageTitle: "Page Title",
        titleSuffix: "Title Suffix",
        outlineLevel: "Outline Display Level",
        htmlLang: "HTML Language",
        siteDescription: "Site Description",
        externalLinkIcon: "External Link Icon",
        docPath: "Docs Path",
        baseUrl: "Base URL",
        buildPath: "Build Path",
        hostname: "Hostname",
        staticAssets: "Generated Static Directory",
        cachePath: "Cache Path",
        cleanUrls: "Clean URLs",
      },
      placeholders: {
        logoUrl: "Homepage image URL",
        pageTitle: "Enter site title",
        titleSuffix: "Enter project title suffix",
        outlineLevel: "Enter outline display level",
        htmlLang: "Site's lang attribute",
        siteDescription: "Enter site description",
        baseUrl: "Base URL",
        buildPath: "Enter build path",
        hostname: "Enter hostname",
      },
    },

    settingLang: {
      title: "Multi-language Settings",
      tooltips: {
        multiLanguageEnabled:
          "[Enabled when language list is not empty], in a multi-language environment, each language has independent navigation, sidebar, and other configurations. It's recommended to decide whether the project needs multi-language support before enabling. Avoid casually deleting languages, as doing so will also delete that language's navigation, sidebar, and multi-language configurations.",
        i18nRouting:
          "When enabled, routes will automatically include the directory name, e.g., changing the local language to zh will change the URL from /foo (or /en/foo/) to /zh/foo.",
        siteTitleInfo:
          "The site title displayed after the logo, equivalent to themeConfig.siteTitle.",
        footerMessageInfo:
          "Due to design, the footer is only shown when the page does not contain a sidebar.",
        copyrightInfo:
          "Shown in the footer due to design constraints, visible only when the sidebar is absent.",
        langMenuLabelInfo:
          "Custom aria-label for the language switch button in the navbar, used only when i18n is enabled.",
        returnToTopLabelInfo:
          "Custom label for the back-to-top button, visible only in mobile views.",
        sidebarMenuLabelInfo:
          "Custom label for the sidebar menu, visible only in mobile views.",
        darkModeSwitchLabelInfo:
          "Custom label for the dark mode switch, visible only in mobile views.",
        lightModeSwitchTitleInfo:
          "Custom title for the light mode switch when hovered.",
        darkModeSwitchTitleInfo:
          "Custom title for the dark mode switch when hovered.",
        editLinkTextInfo:
          "Edit link enables displaying a link to edit the page on Git management services like GitHub or GitLab.",
        editLinkPatternInfo:
          "Pattern for the edit link to pages on Git management services.",
        prevNextPageInfo:
          "Globally enable or disable previous/next page links and customize their text.",
        nextButtonTextInfo: "Text to appear above the next page link.",
        preButtonTextInfo: "Text to appear above the previous page link.",
      },
      placeholders: {
        langDirectory: "Language Directory",
        langLabel: "Language Label",
        siteTitle: "Enter site title",
        footerMessage: "Enter footer message",
        copyright: "Enter footer copyright",
        langSwitchLabel: "Enter language switch label",
        returnToTop: "Back to top button label",
        sidebarMenu: "Custom sidebar menu label",
        darkModeSwitch: "Dark mode switch label",
        lightModeSwitchTitle: "Light mode switch hover title",
        darkModeSwitchTitle: "Dark mode switch hover title",
        editLinkText: "Edit link text",
        editLinkPattern: "Edit link pattern",
      },
      labels: {
        addLangType: "Add Language Type",
        removeConfirm:
          "All data including configuration, navigation, sidebar, etc., will be deleted.",
        prevButtonText: "Previous Page Text",
        nextButtonText: "Next Page Text",
        prevButton: "Enable Previous Page Link",
        nextButton: "Enable Next Page Link",
        i18nRouting: "Enable i18n Routing",
        siteTitle: "Site Title",
        footerMessage: "Footer Message",
        copyright: "Copyright",
        langSwitchLabel: "Language Switch Label",
        returnToTop: "Back to Top Label",
        sidebarMenu: "Custom Sidebar Menu Label",
        darkModeSwitch: "Dark Mode Switch Label",
        lightModeSwitchTitle: "Hover Light Mode Title",
        darkModeSwitchTitle: "Hover Dark Mode Title",
        editLinkText: "Edit Link Text",
        editLinkPattern: "Edit Link Pattern",
      },
    },
    settingSearch: {
      title: "Search Settings",
      searchProvider: "Search Provider",
      tooltips: {
        tooltips:
          "VitePress supports fuzzy full-text search using browser indexing (local).",
        AlgoliaAppId:
          "Algolia Search AppId, found in settings -> Team and Access -> API keys -> Application ID.",
        AlgoliaSearchKey:
          "Algolia Search apiKey, found in settings -> Team and Access -> API keys -> Search-Only API Key.",
        AlgoliaIndexName:
          "Algolia Search indexName, the name of the created index.",
      },
    },
    settingSocial: {
      title: "Social Accounts",
      tooltipsIcon: "Supported icons (click to copy):",
      addSocialAccount: "Add Social Account",
      icon: "Icon Identifier",
      link: "Redirect Link",
    },
    settingRewrite: {
      title: "Route Rewrite",
      addRewriteRule: "Add Route Rewrite Rule",
      filePath: "File Path",
      routePath: "Route Path",
    },
  },
  pageSetting: {
    settingBase: {
      title: "Basic Settings",
      autoSaveArticle: {
        label: "Toggle Auto-save on Article Switch:",
        tooltip: "Automatically saves articles when switching.",
      },
      startup: {
        label: "Start with System:",
        tooltip: "Takes effect after restarting the program.",
      },
      frontMatterSaveType: {
        label: "Front Matter Save Method:",
        tooltip: "Method for saving article front matter.",
      },
      editorCdn: {
        label: "Editor CDN:",
        tooltip:
          "If the markdown editor fails to load, manually switch the editor CDN and then 【Restart】 the program.",
      },
      sysUpdateSource: {
        label: "Update Source:",
        tooltip:
          "Chinese users are recommended to use Gitee source for updates, while non-Chinese users should prefer GitHub.",
      },
      checkUpdate: "Check for Updates",
      openDataDir: "Open Program Data Directory",
    },
  },

  pageAbout: {
    // 关于
    title: "about software",
    // 隐私协议
    privacy: "Privacy Policy",
    // 关于作者
    author: "About the Author",
    // 作者博客：
    authorBlog: "Author's Blog:",
    // 作者GitHub
    authorGithub: "Author's GitHub",
    // 作者Gitee
    authorGitee: "Author's Gitee",
    // 作者微信
    authorWx: "Author's WeChat",
    privacy1:
      "Welcome to use VitePressSimple (hereinafter referred to as 'VPSimple' or 'this software'). We understand the importance of personal privacy and are committed to protecting your personal information security. This Privacy Policy clearly explains how we collect, use, store, share, and protect your personal information while you use VPSimple. Please read and understand this agreement before using the software. Your continued use will be deemed as acceptance of this Privacy Policy.",
    privacy2: "Necessary Usage Information",
    privacy3:
      ": To provide and optimize VPSimple's services, we may collect basic usage data generated during your use of the software, such as device type, operating system version, software version, etc. This information is necessary to ensure software operation, troubleshoot issues, provide technical support, and improve product performance.",
    privacy4: "Website Analytics",
    privacy5:
      ": VPSimple uses third-party service provider 51.la for website analytics to help us understand user behavior, optimize user experience, and enhance service quality. 51.la might collect anonymous statistical data such as visit logs, IP addresses, browser types, access times, and page interactions. Please note that 51.la's privacy policy is independent of this agreement, and their data collection and processing activities are governed by their own privacy policy. Refer to the\n" +
      "51.la Privacy Policy for more details.",
    privacy6:
      "VPSimple strictly adheres to legal requirements and takes appropriate technical and organizational measures to protect the security of your personal information, preventing loss, misuse, unauthorized access, disclosure, or modification. We retain your personal information for the shortest time necessary to achieve the intended purposes and promptly delete or anonymize it when no longer needed.",
    privacy7:
      "VPSimple is released under the Apache License\n" +
      "2.0 open-source license. While the software follows open-source principles, the handling of user personal information still adheres to this Privacy Policy. The open-source license does not affect or alter our commitment to protecting your personal information.",
    privacy8:
      "This Privacy Policy takes effect from the date of publication. Thank you for choosing VitePressSimple; we will continue to respect and protect your personal information rights.",
  },
};
