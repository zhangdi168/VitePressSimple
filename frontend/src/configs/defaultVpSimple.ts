//存储项目级的vpsimple软件配置
import { VPSimpleConfig } from "@/types/vpsimpleConfig";

export const defaultVpSimple: VPSimpleConfig = {
  shellBaseDir: "", //shell运行目录
  gitBaseDir: "", //git运行目录
  cmdDocsDev: "npm run docs:dev", //启动文档开发服务
  cmdDocsBuild: "npm run docs:build", //打包文档
  cmdNpmInstall: "npm install", //启动文档服务
  cmdGitInit: "git init", //git init
  cmdGitPull: "git pull", //git pull
  cmdGitAdd: "git add .",
  cmdGitCommit: "git commit -m 'autoupdate'",
  cmdGitPush: "git push", //git add . + git commit + git push
};
