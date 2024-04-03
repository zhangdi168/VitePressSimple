import { defineStore } from "pinia";
import { ConfigGet } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyLayoutNavBgColor } from "@/constant/keys/config";
import { IsEmptyValue } from "@/utils/utils";

//这是关于布局控制的状态管理文件
export interface layoutStore {
  showEditorView: boolean; //是否显示编辑器右侧的编辑区
  editorToolIconSize: number; //编辑器工具栏的图标大小
  componentDyAddHeader: any;
  componentDyAddCustomFrontMatter: any;
  colorBgNav: string;
}

export const useLayoutStore = defineStore("layout", {
  state: (): layoutStore => ({
    showEditorView: true,
    editorToolIconSize: 24,
    componentDyAddHeader: null,
    componentDyAddCustomFrontMatter: null,
    colorBgNav: "#ebebeb",
  }),
  actions: {
    //切换显示编辑区
    setEditorViewShow(isShow: boolean) {
      console.log(isShow, "isShow -- console.log");
      this.showEditorView = isShow;
    },
    async loadUserSetting() {
      //左侧导航栏背景颜色
      const leftNavBgColor = await ConfigGet(ConfigKeyLayoutNavBgColor);
      console.log(leftNavBgColor, "leftNavBgColor -- console.log");
      if (!IsEmptyValue(leftNavBgColor)) this.colorBgNav = leftNavBgColor;
    },
  },
  getters: {
    IsShowEditorView: (state) => state.showEditorView,
    //编辑器左编辑区的宽度(百分比)
    // EditorSplitterUnit: (state) => (state.showEditorView ? "px" : "%"),
    EditorLeftWriterWidth: (state) => {
      if (state.showEditorView) {
        //如果显示编辑区
        return 70; //70%
      } else {
        //如果隐藏编辑区
        return 100; //100%
      }
    },
    // EditorToolIconSize: (state) => state.editorToolIconSize,
  },
});
