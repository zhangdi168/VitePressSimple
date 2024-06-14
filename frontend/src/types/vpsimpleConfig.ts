export interface VPSimpleConfig {
  shellBaseDir: string;
  gitBaseDir: string; // git  仓库根目录
  cmdDocsDev: string;
  cmdDocsBuild: string;
  cmdNpmInstall: string;
  cmdGitInit: string;
  cmdGitPull: string;
  cmdGitAdd: string;
  cmdGitCommit: string;
  cmdGitPush: string;
}
