import { defineStore } from "pinia";
import {
  GetVpConfigData,
  SaveConfig,
} from "../../wailsjs/go/vpsimpler/VpConfig";
import { IsEmptyValue, parseJsObject } from "@/utils/utils";
import { ConfigGet, PathJoin } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import { ToastCheck } from "@/utils/Toast";
import { defaultShareConfigValue } from "@/configs/defaultShareConfig";
import {
  defaultLangConfig,
  defaultThemeConfig,
} from "@/configs/defaultLangConfig";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把vpconfig全局替换成你的store名字
export interface vpconfigStore {
  srcDir: string;
  fullSrcDir: string;
  baseDir: string;
  isInstall: boolean;
  configData: any;
  currLangConfig: Record<string, any>;
  currSettingLang: string; //当前正在设置的语言
}

export const useVpconfigStore = defineStore("vpconfig", {
  state: (): vpconfigStore => ({
    srcDir: "./",
    fullSrcDir: "", //完整的doc文档原目录（绝对路径）
    baseDir: "", //绝对路径
    isInstall: false,
    configData: {}, //所有语言的公共配置
    currLangConfig: {}, //当前编辑的语言的主题配置
    currSettingLang: "",
  }),
  actions: {
    async initConfig() {
      await this.readVpConfig();
    },
    //获取config.mts文件内容
    async readVpConfig() {
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
    changeCurrLangConfig(lang: string) {
      if (!this.IsI18nRouting) {
        //没有启用多语言
        this.currSettingLang = "";
        this.currLangConfig = this.configData;
        return;
      }

      this.currSettingLang = lang;
      if (this.currSettingLang == "") {
        const keys = Object.keys(this.configData["locales"] ?? []);
        if (keys.length == 0) {
          //没有配置多语言
          this.currSettingLang = "";
          this.currLangConfig = this.configData;
          return;
        } else {
          this.currSettingLang = keys[0];
        }
      } else {
        const config = (this.currLangConfig =
          this.configData["locales"][this.currSettingLang]);
        if (IsEmptyValue(config)) {
          this.currLangConfig = defaultLangConfig;
        } else {
          this.currLangConfig = config;
        }
      }
    },

    setDefaultValue() {
      //设置 共享配置的默认值
      for (const key in defaultShareConfigValue) {
        if (IsEmptyValue(this.configData[key])) {
          this.configData[key] = defaultShareConfigValue[key];
        }
      }
      //初始化当前的语言 使用第一个语言作为当前语言，如果没使用多语言，则currLangConfig指向默认配置
      this.changeCurrLangConfig("");
      for (const key in defaultLangConfig) {
        if (IsEmptyValue(this.currLangConfig[key])) {
          this.currLangConfig[key] = defaultLangConfig[key];
        }
        if (key === "themeConfig") {
          //主题配置
          for (const key2 in defaultThemeConfig) {
            if (IsEmptyValue(this.currLangConfig[key][key2])) {
              this.currLangConfig[key][key2] = defaultThemeConfig[key2];
            }
          }
        }
      }
    },

    //保存配置
    async saveConfig() {
      if (this.currSettingLang == "") {
        const res = await SaveConfig(JSON.stringify(this.configData));
        ToastCheck(res);
      }
    },
  },
  getters: {
    //相对路径
    SrcDir: (state) => state.srcDir,
    //是否使用多语言
    IsI18nRouting: (state) => {
      return !!state.configData.themeConfig.i18nRouting;
    },
    GetLangArray(state) {
      const langArray = [];
      for (const key in state.configData.locales) {
        langArray.push(key);
      }
      return langArray;
    },
  },
});
