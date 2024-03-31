import { defineStore } from "pinia";
import Vditor from "vditor";
import { ToastCheck, ToastError, ToastInfo, ToastSuccess } from "@/utils/Toast";
import {
  ParseTreeData,
  WriteFileContent,
} from "../../wailsjs/go/services/ArticleTreeData";
import { dto } from "../../wailsjs/go/models";
import { getDirectoryPath } from "@/utils/file";
import {
  ConfigSet,
  PathExists,
  PathJoin,
  SelectDir,
} from "../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import { useVpconfigStore } from "@/store/vpconfig";
import { getFileNameFromPath, IsEmptyValue } from "@/utils/utils";
import { useLayoutStore } from "@/store/layout";
import { VitePressHome } from "@/types/home";
import { moveTo } from "@/utils/system";
import { defaultFrontMatter } from "@/configs/defaultFrontMatter";
import { HistoryProject } from "@/utils/historyProject";
import { useHistoryStore } from "@/store/history";
import { isEmptyArray } from "@/utils/array";

//定义首页的数据类型
export interface indexStore {
  articleTreeData: dto.TreeNode[] | null;
  vditor: Vditor | null;
  expandKeys: string[];
  selectKeys: string[];
  currArticlePath: string; //当前文章路径
  currCopyPath: string;
  searchValue: string; //搜索值
  currHomeConfig?: VitePressHome; //当前主页配置
  currScriptContent: string; //当前js代码
  currStyleContent: string; //当前css代码
  currVueCode: string; //当前vue代码 css+js
  currProjectDir: string; //当前项目的根目录
  currDocDir: string; //文档所在目录
  IsEmptyProject: boolean; //是否为空项目
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
    IsEmptyProject: true, //是否为空项目
    searchValue: "", //搜索值
    currArticlePath: "", //当前文章路径
    currScriptContent: "", //当前js代码
    currStyleContent: "", //当前css代码
    currVueCode: "", //当前vue代码 js+css
    currArticleFrontMatter: {}, //当前文章front matter
    currCopyPath: "", //当前剪切路径
  }),
  //定义actions
  actions: {
    async loadTreeData() {
      const cfg = useVpconfigStore();
      if (await cfg.ExistsProjectDir()) {
        await cfg.readVpConfig(); //读取配置 以及组装路径
        this.articleTreeData = await ParseTreeData(cfg.srcDir);
      }
    },

    async checkProjectDir(dir: string): Promise<string> {
      //判空
      if (dir == "") {
        return "不能切换为空路径项目";
      }
      //判断兖是否存在
      const isExists = await PathExists(dir);
      if (!isExists) {
        useHistoryStore().remove(dir);
        return "项目路径不存在:" + dir;
      }
      //判断是否是.vitepress目录
      const vitePressDir = await PathJoin([dir, ".vitepress"]);
      const isExistsVpDir = await PathExists(vitePressDir);
      if (!isExistsVpDir) {
        return `切换失败，所选路径${dir}不存在“.vitepress”文件夹`;
      }
      return "";
    },
    //切换项目
    async changeProject(dir: string) {
      //先检查传入的项目目录是否合法
      const checkString = await this.checkProjectDir(dir);
      if (checkString != "") {
        ToastError(checkString);
        return;
      }
      this.IsEmptyProject = false; //设置为非空项目
      useHistoryStore().add(dir); //加入到历史项目数据中
      //设置当前的项目路径
      await ConfigSet(ConfigKeyProjectDir, dir);
      this.clearCurrData();
      await this.loadTreeData();
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
      this.currCopyPath = "";
      this.currVueCode = "";
      this.currArticleFrontMatter = {};
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
      this.currCopyPath = path;
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
    CurrCutPath: (state) => state.currCopyPath,
    // getCurrVueCode: (state) => state.currVueCode,
    CurrArticlePath: (state) => state.currArticlePath,
    CurrArticleFrontMatter: (state) => state.currArticleFrontMatter,
    Vditor: (state) => state.vditor,
    CurrProjectDir: (state) => state.currProjectDir,
    // IsEmptyProject: (state) => isEmptyArray(state.articleTreeData),
    IsEmptyTreeData: (state) => isEmptyArray(state.articleTreeData),
    CurrArticleTitle: (state) =>
      getFileNameFromPath(state.currArticlePath).replaceAll(".md", ""),
    GetArticleFrontMatter: (state) => state.currArticleFrontMatter,
  },
});
