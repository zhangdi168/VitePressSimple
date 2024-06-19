import { defaultVpSimple } from "@/configs/defaultVpSimple";

export const EnShell = {
  //终端配置
  config: "Terminal Config",
  //运行日志
  runLog: "Run Log",
  //默认值
  default: "Default",
  //暂无运行中的终端
  noRunTerminal: "No Running Terminal",
  //起始目录不能为空
  runPathEmpty: "Run Path Can Not Be Empty",
  //"cmd不能为空"
  cmdEmpty: "Cmd Can Not Be Empty",

  docsBuild: "Docs Build Command",
  //运行目录
  labels: {
    //默认目录
    runPath: "Default Run Path",
    //git运行的目录
    gitPath: "Git Run Path",
    //文档dev命令
    docsDev: "Docs Dev Command",
    //文档build命令
    docsBuild: "Docs Build Command",
  },
  tips: {
    //运行目录的路径，默认值是项目根目录
    runPath: "Run Path, Default Value Is Project Root Path",
    //git运行的目录，一般为vitepress项目根目录，但有时vitepres是做其它项目的子项目，所以需要指定
    gitPath:
      "Git Run Path, Default Value Is Vitepress Project Root Path, But Sometimes Vitepress Is A Sub Project Of Other Projects, So You Need To Specify",
    //文档dev命令，默认值是vitepress项目根目录下的dev命令
    docsDev: "Docs Dev Command, Default Value Is " + defaultVpSimple.cmdDocsDev,
    //文档build命令，默认值是vitepress项目根目录下的build命令
    docsBuild:
      "Docs Build Command, Default Value Is" + defaultVpSimple.cmdDocsBuild,
  },
};
