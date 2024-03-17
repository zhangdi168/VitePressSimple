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
import { useVpconfigStore } from "@/store/vpconfig";
import { getFileNameFromPath, IsEmptyValue } from "@/utils/utils";
import { useLayoutStore } from "@/store/layout";

//定义首页的数据类型
export interface indexStore {
  articleTreeData: dto.TreeNode[] | null;
  vditor: Vditor | null;
  expandKeys: string[];
  selectKeys: string[];
  currArticlePath: string; //当前文章路径
  currCutPath: string;
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
      await cfg.initConfig();
      this.articleTreeData = await ParseTreeData(cfg.SrcDir);
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
    //设置当前script
    async setCurrScriptContent(content: string) {
      this.currScriptContent = content;
      this.currVueCode = this.currScriptContent + "\n" + this.currStyleContent;
    },
    //设置当前style
    async setCurrStyleContent(content: string) {
      this.currStyleContent = content;
      this.currVueCode = this.currScriptContent + "\n" + this.currStyleContent;
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
        const storeLayout = useLayoutStore();
        const content = this.vditor.getValue();
        this.currArticleFrontMatter["custom"] =
          storeLayout.componentDyAddCustomFrontMatter.getObject();
        this.currArticleFrontMatter["head"]["meta"] =
          storeLayout.componentDyAddHeader.getArray();
        const fontMatterStr = JSON.stringify(this.currArticleFrontMatter);
        const fullContent = `---\n${fontMatterStr}\n---\n${this.currVueCode}\n${content}`;
        //获取动态新增的数据

        // metas
        // console.log(
        //
        //   "storeLayout.refDyAddHeader.getArray() -- console.log",
        // );
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
      this.checkFrontMatterKey("title", this.CurrArticleTitle);
      this.checkFrontMatterKey("description", "");
      this.checkFrontMatterKey("navbar", true); //是否显示导航
      this.checkFrontMatterKey("sideBar", true); //是否显示侧栏
      this.checkFrontMatterKey("footer", false); //是否显示页脚
      this.checkFrontMatterKey("outline", 2); //是否显示页脚
      this.checkFrontMatterKey("editLink", false); //是否显示编辑链接
      this.checkFrontMatterKey("lastUpdated", true); //是否显示页脚更新时间
      this.checkFrontMatterKey("aside", "left"); //侧边栏位置
      this.checkFrontMatterKey("layout", "doc"); //页面类型
      this.checkFrontMatterKey("custom", {}); //用户自定义变量
      this.checkFrontMatterKey2("head", "meta", []); //页面类型
      // console.log(frontMatter, "frontMatter -- console.log");
    },
    //判断front matter的key是否存在,不存在则设置默认值
    checkFrontMatterKey(key: string, defaultValue: any) {
      if (IsEmptyValue(this.currArticleFrontMatter[key])) {
        this.currArticleFrontMatter[key] = defaultValue;
      }
    },
    //判断front matter的key是否存在,不存在则设置默认值
    checkFrontMatterKey2(key1: string, key2: string, defaultValue: any) {
      if (IsEmptyValue(this.currArticleFrontMatter[key1])) {
        this.currArticleFrontMatter[key1] = {};
        this.currArticleFrontMatter[key1][key2] = defaultValue;
        return;
      }
      if (IsEmptyValue(this.currArticleFrontMatter[key1][key2])) {
        this.currArticleFrontMatter[key1][key2] = defaultValue;
      }
    },
  },
  getters: {
    //大驼峰命名法
    ArticleTreeData: (state) => state.articleTreeData,
    ExpandKeys: (state) => state.expandKeys,
    SelectKeys: (state) => state.selectKeys,
    CurrCutPath: (state) => state.currCutPath,
    getCurrVueCode: (state) => state.currVueCode,
    CurrArticlePath: (state) => state.currArticlePath,
    CurrArticleFrontMatter: (state) => state.currArticleFrontMatter,
    Vditor: (state) => state.vditor,
    CurrProjectDir: (state) => state.currProjectDir,
    CurrDocDir: (state) => state.currDocDir,
    CurrArticleTitle: (state) =>
      getFileNameFromPath(state.currArticlePath).replaceAll(".md", ""),
    GetArticleFrontMatter: (state) => state.currArticleFrontMatter,
  },
});
