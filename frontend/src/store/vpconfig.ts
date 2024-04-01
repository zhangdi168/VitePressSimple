import { defineStore } from "pinia";
import {
  GetVpConfigData,
  SaveConfig,
} from "../../wailsjs/go/vpsimpler/VpConfig";
import { IsEmptyValue, parseJsObject } from "@/utils/utils";
import {
  ConfigGet,
  CopyPath,
  PathExists,
  PathJoin,
} from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import { ToastCheck, ToastInfo } from "@/utils/Toast";
import { defaultShareConfigValue } from "@/configs/defaultShareConfig";
import {
  defaultLangConfig,
  defaultThemeConfig,
} from "@/configs/defaultLangConfig";
import { StringGlobalLang, StringRootLang } from "@/configs/cnts";
import { CreateDir } from "../../wailsjs/go/services/ArticleTreeData";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把vpconfig全局替换成你的store名字
export interface vpconfigStore {
  srcDir: string; //相对路径
  fullSrcDir: string;
  baseDir: string;
  isInstall: boolean;
  configData: any;
  currLangConfig: Record<string, any>;
  currSettingLang: string; //当前正在设置的语言
  currLangConfigIsUseRootConfig: boolean; //当前语言配置是否指向根目录
}

export const useVpconfigStore = defineStore("vpconfig", {
  state: (): vpconfigStore => ({
    srcDir: "./", //doc目录（相对路径），取自配置文件
    baseDir: "", //根目录（绝对路径，已经过join）
    fullSrcDir: "", //doc目录（绝对路径,已经过join）
    isInstall: false,
    configData: {}, //所有语言的公共配置
    currLangConfig: {}, //当前编辑的语言的主题配置
    currSettingLang: "",
    currLangConfigIsUseRootConfig: false, //当前语言配置是否指向根目录
  }),
  actions: {
    async formatPath() {
      this.baseDir = await PathJoin([await ConfigGet(ConfigKeyProjectDir)]);
    },
    //获取config.mts文件内容
    async readVpConfig() {
      //获取项目根目录(绝对路径)
      await this.formatPath();
      await this.backupConfigFile(); //如果是首次则备份文件夹，放在 this.baseDir后面
      const content = await GetVpConfigData(); //获取config.mts文件内容
      let configData: any = {};
      if (content == "") {
        ToastInfo("读取配置文件内容为空");
      } else {
        configData = parseJsObject(content); //解析config.mts文件内容
      }
      this.configData = configData ?? {};
      this.setDefaultValue(); //检测key不存在则使用默认值
      this.srcDir = configData["srcDir"];
      this.fullSrcDir = await PathJoin([this.baseDir, this.srcDir]);
      //判断如果原目录不存在则自动创建
      if (!(await PathExists(this.fullSrcDir))) {
        await CreateDir(this.fullSrcDir);
        ToastInfo(`检测到源目录不存在，已自动创建源目录:${this.fullSrcDir}`);
      }
    },
    async ExistsProjectDir() {
      await this.formatPath();
      const isExists = await PathExists(this.baseDir);
      if (isExists) {
        return true;
      } else {
        // ToastError("项目目录不存在:" + this.baseDir);
        return false;
      }
    },
    //备份原来的配置文件，（仅仅备份一次）如果已经存在则不再进行备份
    async backupConfigFile() {
      const originDir = await PathJoin([this.baseDir, ".vitepress"]);
      const backupDir = await PathJoin([this.baseDir, ".vitepressBak"]);
      const isExistsOriginDir = await PathExists(originDir);
      const isExistsBackupDir = await PathExists(backupDir);
      if (isExistsOriginDir && !isExistsBackupDir) {
        await CopyPath(originDir, backupDir, true);
        // ToastCheck(res, "备份配置文件成功");
      }
    },
    changeCurrLangConfig(lang: string) {
      if (!this.IsUseI18n || lang === StringGlobalLang) {
        this.currLangConfigUseRootConfig(); //当前语言配置指向根目录
        return;
      }
      this.currSettingLang = lang == "" ? this.getFirstLang() : lang;
      if (this.currSettingLang == "") {
        this.currLangConfigUseRootConfig(); //当前语言配置指向根目录
        return;
      }
      const config = this.configData["locales"][this.currSettingLang];
      this.currLangConfig = IsEmptyValue(config) ? defaultLangConfig : config;
      this.checkCurrLangConfig(); //默认值检测
      this.currLangConfigIsUseRootConfig = false;
    },
    //设置当前语言配置指向根目录
    currLangConfigUseRootConfig() {
      this.currSettingLang = StringGlobalLang;
      this.currLangConfig = this.configData;
      this.currLangConfigIsUseRootConfig = true;
    },
    getFirstLang() {
      const keys = Object.keys(this.configData["locales"] ?? []);
      if (keys.length == 0) {
        //没有配置多语言
        return "";
      } else {
        return keys[0];
      }
    },
    setDefaultValue() {
      this.checkRootConfig();
      //初始化当前的语言 使用第一个语言作为当前语言，如果没使用多语言，则currLangConfig指向默认配置
      if (this.IsUseI18n) {
        this.changeCurrLangConfig(this.getFirstLang());
      } else {
        this.changeCurrLangConfig(StringGlobalLang);
      }

      this.checkCurrLangConfig();
    },
    checkRootConfig() {
      //设置 共享配置的默认值
      for (const key in defaultShareConfigValue) {
        if (IsEmptyValue(this.configData[key])) {
          this.configData[key] = defaultShareConfigValue[key];
        }
        if (key === "themeConfig") {
          //主题配置
          for (const key2 in defaultShareConfigValue.themeConfig) {
            if (IsEmptyValue(this.configData[key][key2])) {
              this.configData[key][key2] =
                defaultShareConfigValue.themeConfig[key2];
            }
          }
        }
      }
    },
    //检查当前语言配置,如果不存在key则使用默认值
    checkCurrLangConfig() {
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
      const res = await SaveConfig(JSON.stringify(this.configData));
      ToastCheck(res);
    },
  },
  getters: {
    //相对路径
    // SrcDir: (state) => state.srcDir,
    //基于语言的相对路径
    SrcLangDir: (state) => {
      if (
        state.currSettingLang == StringGlobalLang ||
        state.currSettingLang == StringRootLang
      ) {
        return state.srcDir;
      } else {
        return state.srcDir + "/" + state.currSettingLang;
      }
    },
    FullSrcLangDir: (state) => {
      if (
        state.currSettingLang == StringGlobalLang ||
        state.currSettingLang == StringRootLang
      ) {
        return state.fullSrcDir;
      } else {
        return state.fullSrcDir + "/" + state.currSettingLang;
      }
    },
    //是否使用多语言
    IsUseI18n: (state) => {
      return (
        !IsEmptyValue(state.configData.locales) &&
        Object.keys(state.configData.locales).length > 0
      );
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
