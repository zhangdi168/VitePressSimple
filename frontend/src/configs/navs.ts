import { lang } from "@/utils/language";

export interface navItem {
  title: string;
  name: string;
  path: string;
  icon?: string;
  iconColor?: string;
  borderTop?: boolean;
}

export const NavList = (): navItem[] => {
  // const color = "#8babed";
  const color = "#676a67";
  return [
    {
      title: lang("nav.home"),
      name: "index",
      path: "/",
      icon: "home", //icon park
      borderTop: false,
      iconColor: color,
    },

    {
      title: lang("nav.config"),
      name: "projectSetting",
      path: "/projectSetting",
      icon: "config",
      borderTop: false,
      iconColor: color,
    },
    {
      title: lang("nav.nav"),
      name: "navSetting",
      path: "/navSetting",
      icon: "navigation",
      borderTop: false,
      iconColor: color,
    },
    {
      title: lang("nav.sidebar"),
      name: "sidebarSetting",
      path: "/sidebarSetting",
      icon: "mindmap-list",
      borderTop: false,
      iconColor: color,
    },
    // {
    //   title: lang("nav.preview"),
    //   name: "preview",
    //   path: "/preview",
    //   icon: "web-page",
    //   borderTop: false,
    // },
    {
      title: lang("nav.setting"),
      name: "setting",
      path: "/setting",
      icon: "more-app",
      borderTop: true,
      iconColor: color,
    },
    {
      title: lang("nav.about"),
      name: "a",
      path: "/about",
      icon: "more-three",
      borderTop: false,
      iconColor: color,
    },
  ];
};
