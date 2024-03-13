import { defineStore } from "pinia";
import Vditor from "vditor";
import { ToastError, ToastInfo } from "@/utils/Toast";
import {
  Move,
  ParseTreeData,
  WriteFileContent,
} from "../../wailsjs/go/services/ArticleTreeData";
import { dto } from "../../wailsjs/go/models";
import { getDirectoryPath } from "@/utils/file";
import {
  ConfigGet,
  ConfigSet,
  SelectDir,
} from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import { GetVpConfigData } from "../../wailsjs/go/vpsimpler/VpConfig";

//定义首页的数据类型
export interface indexStore {
  articleTreeData: dto.TreeNode[] | null;
  vditor: Vditor | null;
  expandKeys: string[];
  selectKeys: string[];
  currArticlePath: string; //当前文章路径
  currCutPath: string;
  currProjectDir: string; //当前项目的根目录
  currDocDir: string; //文档所在目录
  currVueCode: string; //当前vue代码
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
    currCutPath: "", //当前剪切路径
    currVueCode: "", //当前vue代码
    currArticleFrontMatter: {}, //当前文章front matter
  }),
  //定义actions
  actions: {
    async loadTreeData() {
      GetVpConfigData().then((content) => {
        // const data = JSON.parse(content);
        console.log(content, "content -- console.log");
      });
      return;
      this.articleTreeData = await ParseTreeData();
      console.log(this.articleTreeData, "this.articleTreeData -- console.log");
    },
    //设置编辑器的实例
    async setVditorInstance(vditor_: Vditor | null) {
      this.vditor = vditor_;
    },
    moveTo(path: string, target: string) {
      const targetDir = getDirectoryPath(target);
      console.log(target, targetDir);

      Move(path, targetDir).then(() => {
        this.loadTreeData();
      });
    },
    clearCurrData() {
      this.currArticlePath = "";
      this.currCutPath = "";
      this.currVueCode = "";
      this.currArticleFrontMatter = {};
    },
    loadCurrProjectFromConfig() {
      ConfigGet(ConfigKeyProjectDir).then((dir) => {
        if (dir != "") {
          this.currDocDir = dir;
          this.loadTreeData();
        }
      });
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
        const content = this.vditor.getValue();
        const fontMatterStr = JSON.stringify(this.currArticleFrontMatter);
        const fullContent = `---\n${fontMatterStr}\n---${content}`;
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
      this.currArticleFrontMatter = frontMatter;
    },
    //判断是否有选中文章
    hasSelectArticle() {
      return this.CurrCutPath !== "";
    },
  },
  getters: {
    //大驼峰命名法
    ArticleTreeData: (state) => state.articleTreeData,
    ExpandKeys: (state) => state.expandKeys,
    SelectKeys: (state) => state.selectKeys,
    CurrCutPath: (state) => state.currCutPath,
    CurrVueCode: (state) => state.currVueCode,
    CurrArticlePath: (state) => state.currArticlePath,
    CurrArticleFrontMatter: (state) => state.currArticleFrontMatter,
    Vditor: (state) => state.vditor,
    CurrProjectDir: (state) => state.currProjectDir,
    CurrDocDir: (state) => state.currDocDir,
  },
});
