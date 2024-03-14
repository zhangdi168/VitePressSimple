import { defineStore } from "pinia";

//这是关于布局控制的状态管理文件
export interface layoutStore {
  showEditorView: boolean; //是否显示编辑器右侧的编辑区
  editorToolIconSize: number; //编辑器工具栏的图标大小
}

export const useLayoutStore = defineStore("layout", {
  state: (): layoutStore => ({
    showEditorView: true,
    editorToolIconSize: 26,
  }),
  actions: {
    //切换显示编辑区
    setEditorViewShow(isShow: boolean) {
      console.log(isShow, "isShow -- console.log");
      this.showEditorView = isShow;
    },
  },
  getters: {
    IsShowEditorView: (state) => state.showEditorView,
    //编辑器左编辑区的宽度(百分比)
    EditorLeftWriterWidth: (state) => {
      if (state.showEditorView) {
        //如果显示编辑区，则编辑器占比为72%
        return 72;
      } else {
        return 100;
      }
    },
    // EditorToolIconSize: (state) => state.editorToolIconSize,
  },
});
