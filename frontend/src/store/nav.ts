import { defineStore } from "pinia";
import { useIndexStore } from "@/store/index";

export interface navStore {
  activeName: string; //当前选中的导航
  activePath: string; //当前选中的导航路径
}
export const useNavStore = defineStore("nav", {
  state: (): navStore => ({
    activeName: "index",
    activePath: "/",
  }),
  actions: {
    setActiveNav(path: string, name: string) {
      this.activeName = name;
      this.activePath = path;
      const indexStore = useIndexStore();
      indexStore.clearCurrData();
    },
    isActiveNav(name: string) {
      return this.activeName === name;
    },
  },
  getters: {
    ActiveName: (state) => state.activeName,
    ActivePath: (state) => state.activePath,
  },
});
