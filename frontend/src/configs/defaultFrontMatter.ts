export const defaultFrontMatter: any = {
  title: "", //页面标题
  description: "", //页面描述
  navbar: true, //是否显示导航
  sideBar: true, //是否显示侧栏
  footer: false, //是否显示页脚
  outline: [1, 3], //是否显示页脚
  editLink: false, //是否显示编辑链接
  lastUpdated: true, //是否显示页脚更新时间
  aside: "right", //大纲显示位置--默认右边
  layout: "doc", //页面类型
  custom: {}, //用户自定义变量
  //主页
  hero: {
    image: {
      src: "",
      alt: "",
      width: "",
      height: "",
    },
    name: "VitePressSimple",
    text: "quick to config vitePress",
    description: "",
    tagline: "", //下方标语
    actions: [],
    features: [], //功能列表
    head: [], //页面类型
    // prev: "上一页",
    // next: "",
  }, //用户自定义变量
};
