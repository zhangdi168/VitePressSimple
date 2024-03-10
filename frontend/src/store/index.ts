import { defineStore } from "pinia";
import Vditor from "vditor";
import { ToastInfo } from "@/utils/Toast";
import {
  ParseTreeData,
  WriteFileContent,
} from "../../wailsjs/go/services/ArticleTreeData";
import { dto } from "../../wailsjs/go/models";

//定义首页的数据类型
export interface indexStore {
  articleTreeData: dto.TreeNode[];
  vditor: Vditor | null;
  expandKeys: string[];
  selectKeys: string[];
  currArticlePath: string; //当前文章路径
  currCutPath: string;
  test: string;
  currVueCode: string; //当前vue代码
  currArticleFrontMatter: Record<string, any>; //当前文章front matter
}

export const useIndexStore = defineStore("index", {
  state: (): indexStore => ({
    articleTreeData: [],
    vditor: null,
    test: "123",
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
      this.articleTreeData = await ParseTreeData();
    },
    //设置编辑器的实例
    async setVditorInstance(vditor_: Vditor | null) {
      this.vditor = vditor_;
    },
    clearCurrData() {
      this.currArticlePath = "";
      this.currCutPath = "";
      this.currVueCode = "";
      this.currArticleFrontMatter = {};
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
    Test: (state) => state.test,
  },
});
