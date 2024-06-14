import { defineStore } from "pinia";
import { ConfigGet, ConfigSet } from "../../wailsjs/go/system/SystemService";
import { ConfigKeyHistoryProject } from "@/constant/keys/config";
import { ToastError } from "@/utils/Toast";
//这是一个简单的推荐store案例，可以在这里定义你的状态
//新建pinia时把historyProjects全局替换成你的store名字
export interface historyProjectsStore {
  currentList: string[];
}

export const useHistoryStore = defineStore("historyProjects", {
  state: (): historyProjectsStore => ({
    currentList: [],
  }),
  actions: {
    async add(dir: string) {
      await this.initList();
      if (dir === "") return;
      // 从当前列表中移除指定目录，确保后续添加不会重复
      this.currentList = this.currentList.filter(
        (item) => item !== dir && item !== "",
      );

      // 将指定目录添加到列表的开头
      this.currentList.unshift(dir);

      // 限制列表长度不超过10，超出部分会被移除
      this.currentList = this.currentList.slice(0, 10);

      // 保存更新后的列表到配置文件
      this.saveToConfig();
    },
    async remove(dir: string) {
      await this.initList();
      this.currentList = this.currentList.filter((item) => item !== dir);
      this.saveToConfig();
    },
    //初始化列表
    async initList() {
      this.currentList = [];
      const currListString = await ConfigGet(ConfigKeyHistoryProject);
      try {
        this.currentList = JSON.parse(currListString);
      } catch (e: any) {
        ToastError("读取历史项目失败" + e.message);
        // this.currentList = [];
      }
      // console.log(this.currentList, "this.currentList -- console.log");
    },
    saveToConfig() {
      ConfigSet(ConfigKeyHistoryProject, JSON.stringify(this.currentList));
    },
  },
  getters: {},
});
