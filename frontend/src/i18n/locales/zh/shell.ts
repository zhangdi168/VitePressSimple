import { defaultVpSimple } from "@/configs/defaultVpSimple";

export const ZhShell = {
  //终端配置
  config: "终端配置",
  //运行日志
  runLog: "运行日志",
  //默认值
  default: "默认值",
  //暂无运行中的终端
  noRunTerminal: "暂无运行中的终端",
  //起始目录不能为空
  runPathEmpty: "起始目录不能为空",
  //"cmd不能为空"
  cmdEmpty: "cmd不能为空",

  docsBuild: "文档build命令",
  //运行目录
  labels: {
    //默认目录
    runPath: "默认命令运行目录",
    //git运行的目录
    gitPath: "git运行的目录",
    //文档dev命令
    docsDev: "文档dev命令",
    //文档build命令
    docsBuild: "文档build命令",
    //文档serve命令
  },
  tips: {
    //运行目录的路径，默认值是项目根目录
    runPath: "运行目录的路径，默认值是项目根目录",
    //git运行的目录，一般为vitepress项目根目录，但有时vitepres是做其它项目的子项目，所以需要指定git运行的目录
    gitPath:
      "git运行的目录，一般为vitepress项目根目录，但有时vitepres是做其它项目的子项目，所以需要指定git运行的目录",
    //文档dev命令，默认值是vitepress项目根目录下的dev命令
    docsDev: "文档dev命令，默认值:" + defaultVpSimple.cmdDocsDev,
    //文档build命令，默认值是vitepress项目根目录下的build命令
    docsBuild: "文档build命令，默认值:" + defaultVpSimple.cmdDocsBuild,
  },
};
