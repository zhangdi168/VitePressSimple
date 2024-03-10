import { lang } from "@/utils/language";

export interface navItem {
  title: string;
  name: string;
  path: string;
  icon?: string;
  borderTop?: boolean;
}
export const NavList = (): navItem[] => {
  return [
    {
      title: lang("nav.home"),
      name: "index",
      path: "/",
      icon: "home", //icon park
      borderTop: false,
    },

    {
      title: lang("nav.setting"),
      name: "setting",
      path: "/setting",
      icon: "setting",
      borderTop: false,
    },
    {
      title: lang("nav.about"),
      name: "a",
      path: "/about",
      icon: "other",
      borderTop: false,
    },
  ];
};
