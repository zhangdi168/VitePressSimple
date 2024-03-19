import { defineStore } from "pinia";
import {
  GetVpConfigData,
  SaveConfig,
} from "../../wailsjs/go/vpsimpler/VpConfig";
import { IsEmptyValue, parseJsObject } from "@/utils/utils";
import { ConfigGet, PathJoin } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import { ToastCheck } from "@/utils/Toast";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把vpconfig全局替换成你的store名字
export interface vpconfigStore {
  srcDir: string;
  fullSrcDir: string;
  baseDir: string;
  isInstall: boolean;
  configData: any;
}

export const useVpconfigStore = defineStore("vpconfig", {
  state: (): vpconfigStore => ({
    srcDir: "./",
    fullSrcDir: "", //完整的doc文档原目录（绝对路径）
    baseDir: "", //绝对路径
    isInstall: false,
    configData: {},
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
      this.configData = configData ?? {};
      this.setDefaultValue(); //设置默认值
      this.srcDir = configData["srcDir"];
      this.fullSrcDir = await PathJoin([this.baseDir, this.srcDir]);
    },
    setDefaultValue() {
      this.checkConfigKey("srcDir", "./");
      this.checkConfigKey("cacheDir", "./.vitepress/cache");
      this.checkConfigKey("title", "VitePressSimple");
      this.checkConfigKey("titleTemplate", "| simple config vitepress"); //标题后缀      this.checkConfigKey("description", ""); //标题后缀
      this.checkConfigKey("lang", "en-US"); //标题后缀      this.checkConfigKey("description", ""); //标题后缀
      this.checkConfigKey("base", "/"); //站点将部署到的 base URL。如果计划在子路径例如 GitHub 页面）下部署站点，则需要设置此项
      this.checkConfigKey("outDir", "./.vitepress/dist"); //项目的构建输出位置，相对于项目根目录
      this.checkConfigKey("map", false); //项目的构建输出位置，相对于项目根目录
      this.checkConfigKey("cleanUrls", false); //项目的构建输出位置，相对于项目根目录
      this.checkConfigKey2("themeConfig", "logo", ""); //站点图标
      this.checkConfigKey2("themeConfig", "siteTitle", "VitePressSimple"); //站点名称
      this.checkConfigKey("rewrites", {
        "packages/pkg-a/src/pkg-a-docs.md": "pkg-a/index.md",
        "packages/pkg-b/src/pkg-b-docs.md": "pkg-b/index.md",
      }); //路由重写
    },
    //检查配置key是否存在，不存在则设置默认值
    checkConfigKey(key: string, defaultValue: any) {
      if (IsEmptyValue(this.configData[key])) {
        this.configData[key] = defaultValue;
      }
    },
    //判断front matter的key是否存在,不存在则设置默认值
    checkConfigKey2(key1: string, key2: string, defaultValue: any) {
      if (IsEmptyValue(this.configData[key1])) {
        this.configData[key1] = {};
        this.configData[key1][key2] = defaultValue;
        return;
      }
      if (IsEmptyValue(this.configData[key1][key2])) {
        this.configData[key1][key2] = defaultValue;
      }
    },
    //保存配置
    async saveConfig() {
      const res = await SaveConfig(JSON.stringify(this.configData));
      ToastCheck(res);
    },
  },
  getters: {
    //相对路径
    SrcDir: (state) => state.srcDir,
  },
});
