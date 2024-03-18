export interface Hero {
  // `text` 上方的字符，带有品牌颜色
  // 预计简短，例如产品名称
  name?: string;

  // hero 部分的主要文字，
  // 被定义为 `h1` 标签
  text: string;

  // `text` 下方的标语
  tagline?: string;

  // text 和 tagline 区域旁的图片
  image?: ThemeableImage;

  // 主页 hero 部分的操作按钮
  actions?: HeroAction[];
}

export type ThemeableImage =
  | string
  | { src: string; alt?: string }
  | { light: string; dark: string; alt?: string };

export interface HeroAction {
  // 按钮的颜色主题，默认为 `brand`
  theme?: "brand" | "alt";

  // 按钮的标签
  text: string;

  // 按钮的目标链接
  link: string;

  // 链接的 target 属性
  target?: string;

  // 链接的 rel 属性
  rel?: string;
}
