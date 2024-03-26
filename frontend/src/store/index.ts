import { defineStore } from "pinia";
import Vditor from "vditor";
import { ToastError, ToastInfo } from "@/utils/Toast";
import {
  ParseTreeData,
  WriteFileContent,
} from "../../wailsjs/go/services/ArticleTreeData";
import { dto } from "../../wailsjs/go/models";
import { getDirectoryPath } from "@/utils/file";
import { ConfigSet, SelectDir } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import { useVpconfigStore } from "@/store/vpconfig";
import { getFileNameFromPath, IsEmptyValue } from "@/utils/utils";
import { useLayoutStore } from "@/store/layout";
import { VitePressHome } from "@/types/home";
import { moveTo } from "@/utils/system";
import { defaultFrontMatter } from "@/configs/defaultFrontMatter";

//定义首页的数据类型
export interface indexStore {
  articleTreeData: dto.TreeNode[] | null;
  vditor: Vditor | null;
  expandKeys: string[];
  selectKeys: string[];
  currArticlePath: string; //当前文章路径
  currCutPath: string;
  currHomeConfig?: VitePressHome; //当前主页配置
  currScriptContent: string; //当前js代码
  currStyleContent: string; //当前css代码
  currVueCode: string; //当前vue代码 css+js
  currProjectDir: string; //当前项目的根目录
  currDocDir: string; //文档所在目录
  currArticleFrontMatter: Record<string, any>; //当前文章front matter
}

export const useIndexStore = defineStore("index", {
  state: (): indexStore => ({
    articleTreeData: null,
    vditor: null,
    currProjectDir: "",
    currDocDir: "",
    expandKeys: [],
    selectKeys: [],
    currArticlePath: "", //当前文章路径
    currScriptContent: "", //当前js代码
    currStyleContent: "", //当前css代码
    currVueCode: "", //当前vue代码 js+css
    currArticleFrontMatter: {}, //当前文章front matter
    currCutPath: "", //当前剪切路径
  }),
  //定义actions
  actions: {
    async loadTreeData() {
      const cfg = useVpconfigStore();
      await cfg.readVpConfig();
      //这里只需要传入相对路径即可
      this.articleTreeData = await ParseTreeData(cfg.srcDir);
    },

    //设置编辑器的实例
    async setVditorInstance(vditor_: Vditor | null) {
      this.vditor = vditor_;
    },
    moveTo(path: string, target: string) {
      moveTo(path, getDirectoryPath(target)).then(() => {
        this.loadTreeData();
      });
    },
    //设置当前script
    async setCurrScriptContent(content: string) {
      this.currScriptContent = content;
      this.currVueCode = this.currScriptContent + "\n" + this.currStyleContent;
      this.currVueCode = this.currVueCode.trim();
    },
    //设置当前style
    async setCurrStyleContent(content: string) {
      this.currStyleContent = content;
      this.currVueCode = this.currScriptContent + "\n" + this.currStyleContent;
      this.currVueCode = this.currVueCode.trim();
    },
    clearCurrData() {
      this.currArticlePath = "";
      this.currCutPath = "";
      this.currVueCode = "";
      this.currArticleFrontMatter = {};
    },
    selectProjectDir() {
      SelectDir("请选择项目目录").then((dir) => {
        if (dir === "") {
          return ToastError("取消选择目录");
        }
        this.currProjectDir = dir;
        ConfigSet(ConfigKeyProjectDir, dir);
        this.loadTreeData();
      });
    },
    //保存文章
    async saveCurrArticle() {
      if (this.vditor) {
        const storeLayout = useLayoutStore();
        const content = this.vditor.getValue();
        const fontMatterStr = JSON.stringify(this.currArticleFrontMatter);
        const fullContent = `---\n${fontMatterStr}\n---\n${this.currVueCode}\n${content}`;
        //获取动态新增的数据
        WriteFileContent(this.currArticlePath, fullContent).then(() => {
          ToastInfo("已提交保存");
        });
      }
    },
    //设置当前文章路径
    setCurrArticlePath(path: string) {
      this.currArticlePath = path;
    },
    setCurrCutPath(path: string) {
      this.currCutPath = path;
    },
    setCurrVueCode(code: string) {
      this.currVueCode = code;
    },
    //设置当前文章front matter
    setCurrArticleFrontMatter(frontMatter: any) {
      for (const key in defaultFrontMatter) {
        if (IsEmptyValue(frontMatter[key])) {
          frontMatter[key] = defaultFrontMatter[key];
        }
      }
      if (frontMatter["title"] == "")
        frontMatter["title"] = this.CurrArticleTitle;
      this.currArticleFrontMatter = frontMatter;

      //当前页面是主页
      // if (this.currArticleFrontMatter["layout"] == "home") {
      //   this.currHomeConfig = this.currArticleFrontMatter[
      //     "hero"
      //   ] as VitePressHome;
      // }
    },
  },
  getters: {
    //大驼峰命名法
    ArticleTreeData: (state) => state.articleTreeData,
    ExpandKeys: (state) => state.expandKeys,
    SelectKeys: (state) => state.selectKeys,
    CurrCutPath: (state) => state.currCutPath,
    // getCurrVueCode: (state) => state.currVueCode,
    CurrArticlePath: (state) => state.currArticlePath,
    CurrArticleFrontMatter: (state) => state.currArticleFrontMatter,
    Vditor: (state) => state.vditor,
    CurrProjectDir: (state) => state.currProjectDir,
    GetCurrDocDir: (state) => {
      const cfg = useVpconfigStore();
      const resultDir = state.currDocDir + cfg.srcDir;
      if (cfg.IsUseI18n) {
        //如果设置了多语言 则文档的起始目录为 {源目录}/{lang}
        return resultDir + "/" + cfg.currSettingLang;
      }
      return resultDir;
    },
    CurrArticleTitle: (state) =>
      getFileNameFromPath(state.currArticlePath).replaceAll(".md", ""),
    GetArticleFrontMatter: (state) => state.currArticleFrontMatter,
  },
});
