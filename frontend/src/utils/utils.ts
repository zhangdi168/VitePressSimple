import { ToastError } from "@/utils/Toast";
import { UserConfig } from "vite";

export function parseJsObject(content: string): any {
  // 现在尝试解析为JSON
  content = content.trim();
  try {
    const jsonObj = JSON.parse(content);
    return JSON.stringify(jsonObj);
  } catch (error) {
    //解析异常 尝试使用eval的方式进行解析
    if (!content.endsWith("}") || !content.startsWith("{")) {
      ToastError("config.mts文件内容错误，请检查后重试");
      return "";
    }
    try {
      //不推荐的做法，无奈之举，不然没办法正确解析配置的内容
      return eval(`(${content})`) as UserConfig;
    } catch (e) {
      ToastError("config.mts文件内容解析错误，请检查后重试");
      return null;
    }
  }
}

export const IsEmptyString = (str: string) => {
  return str === "" || str === undefined || str === null;
};
