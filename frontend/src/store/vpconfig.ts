import { defineStore } from "pinia";
import { GetVpConfigData } from "../../wailsjs/go/vpsimpler/VpConfig";
import { parseJsObject } from "@/utils/utils";
import { ConfigGet, PathJoin } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把vpconfig全局替换成你的store名字
export interface vpconfigStore {
  srcDir: string;
  fullSrcDir: string;
  baseDir: string;
  isInstall: boolean;
}

export const useVpconfigStore = defineStore("vpconfig", {
  state: (): vpconfigStore => ({
    srcDir: "./",
    fullSrcDir: "", //完整的doc文档原目录（绝对路径）
    baseDir: "", //绝对路径
    isInstall: false,
  }),
  actions: {
    async initConfig() {
      await this.getConfigFileContent();
    },
    //获取config.mts文件内容
    async getConfigFileContent() {
      //获取项目根目录(绝对路径)
      this.baseDir = await PathJoin([await ConfigGet(ConfigKeyProjectDir)]);

      const content = await GetVpConfigData();
      const configData = parseJsObject(content);
      if (configData == null) return;
      this.srcDir = configData["srcDir"];
      this.fullSrcDir = await PathJoin([this.baseDir, this.srcDir]);
    },
  },
  getters: {
    //相对路径
    SrcDir: (state) => state.srcDir,
  },
});
