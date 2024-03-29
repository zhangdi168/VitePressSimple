import { defineStore } from "pinia";

//这是关于布局控制的状态管理文件
export interface layoutStore {
  showEditorView: boolean; //是否显示编辑器右侧的编辑区
  editorToolIconSize: number; //编辑器工具栏的图标大小
  componentDyAddHeader: any;
  componentDyAddCustomFrontMatter: any;
}

export const useLayoutStore = defineStore("layout", {
  state: (): layoutStore => ({
    showEditorView: true,
    editorToolIconSize: 24,
    componentDyAddHeader: null,
    componentDyAddCustomFrontMatter: null,
  }),
  actions: {
    //切换显示编辑区
    setEditorViewShow(isShow: boolean) {
      console.log(isShow, "isShow -- console.log");
      this.showEditorView = isShow;
    },
    //设置编辑器工具栏的图标大小
    setEditorToolIconSize(size: number) {
      this.editorToolIconSize = size;
    },
    //设置refDyAddHeader的值
    setComponentDyAddHeader(ref: any) {
      this.componentDyAddHeader = ref;
    },
    //设置refDyAddFrontMatter的值
    setComponentDyAddCustomFrontMatter(ref: any) {
      this.componentDyAddCustomFrontMatter = ref;
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
