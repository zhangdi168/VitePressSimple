import { defineStore } from "pinia";
import { GetVpConfigData } from "../../wailsjs/go/vpsimpler/VpConfig";
import { IsEmptyString, parseJsObject } from "@/utils/utils";
import { ConfigGet } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把vpconfig全局替换成你的store名字
export interface vpconfigStore {
  srcDir: string;
  baseDir: string;
  isInstall: boolean;
}

export const useVpconfigStore = defineStore("vpconfig", {
  state: (): vpconfigStore => ({
    srcDir: "./",
    baseDir: "",
    isInstall: false,
  }),
  actions: {
    async initConfig() {
      if (!this.isInstall) {
        await this.getConfigFileContent();
      }
      //获取项目根目录
      this.baseDir = await ConfigGet(ConfigKeyProjectDir);
    },
    //获取config.mts文件内容
    async getConfigFileContent() {
      const content = await GetVpConfigData();
      const configData = parseJsObject(content);
      if (configData == null) return;
      console.log(configData, "vp config get completed----- value");
      this.srcDir = configData["srcDir"];
    },
  },
  getters: {
    //完整的doc文档原目录（绝对路径）
    FullSrcDir: (state) => {
      if (IsEmptyString(state.srcDir) || state.srcDir == "./") {
        return state.baseDir;
      }
      return state.baseDir + "/" + state.srcDir;
    },
    //相对路径
    SrcDir: (state) => state.srcDir,
  },
});
