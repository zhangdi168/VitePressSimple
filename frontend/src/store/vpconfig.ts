import { defineStore } from "pinia";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把vpconfig全局替换成你的store名字
export interface vpconfigStore {
  currDocDir: string;
}

export const useVpconfigStore = defineStore("vpconfig", {
  state: (): vpconfigStore => ({
    currDocDir: "",
  }),
  actions: {
    //获取config.mts文件内容
    getConfigFileContent() {},
  },
  getters: {},
});
