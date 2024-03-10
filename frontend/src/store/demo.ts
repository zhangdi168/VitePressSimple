import { defineStore } from "pinia";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把demo全局替换成你的store名字
export interface demoStore {
  demoName: string;
}
export const useDemoStore = defineStore("demo", {
  state: (): demoStore => ({
    demoName: "",
  }),
  actions: {
    setDemoName(name: string) {
      this.demoName = name;
    },
  },
  getters: {
    DemoName: (state) => state.demoName,
  },
});
