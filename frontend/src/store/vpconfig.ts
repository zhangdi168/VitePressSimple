import { defineStore } from "pinia";
import { GetVpConfigData } from "../../wailsjs/go/vpsimpler/VpConfig";
import { UserConfig } from "vite";
import { ToastError } from "@/utils/Toast";
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
    getConfigFileContent() {
      GetVpConfigData().then((content) => {
        // 使用正则将非标准键名转换为标准JSON格式
        // const standardJson = content.replace(/(?<!https?):(?!\/\/)/g, '":');
        // const data = JSON.parse(standardJson);
        // console.log(JSON.parse(content), "content -- console.log");
        const objectLiteralPattern = /^\s*\{.*\}\s*$/; // 简单的对象字面量检查
        if (objectLiteralPattern.test(content)) {
          // 极其不推荐的做法，无奈之举，不然没办法正确解析配置的内容
          const data = eval(`(${content})`) as UserConfig; // 用括号包裹以确保解析为对象而非代码块
        } else {
          ToastError("配置文件内容校验不通过请检查");
        }
      });
    },
  },
  getters: {},
});
