export interface Feature {
  // 在每个 feature 框中显示图标
  icon?: FeatureIcon;

  // feature 的标题
  title: string;

  // feature 的详情
  details: string;

  // 点击 feature 组件时的链接，可以是内部链接，也可以是外部链接。
  //
  //
  // 例如 `guide/reference/default-theme-home-page` 或 `https://example.com`
  link?: string;

  // feature 组件内显示的链接文本，最好与 `link` 选项一起使用
  //
  //
  // 例如 `Learn more`, `Visit page` 等
  linkText?: string;

  // `link` 选项的链接 rel 属性
  //
  // 例如 `external`
  rel?: string;

  // `link` 选项的链接 target 属性
  target?: string;
}

export type FeatureIcon =
  | string
  | { src: string; alt?: string; width?: string; height: string }
  | {
      light: string;
      dark: string;
      alt?: string;
      width?: string;
      height: string;
    };
