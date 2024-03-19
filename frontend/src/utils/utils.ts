import { ToastError } from "@/utils/Toast";
import { UserConfig } from "vite";

export function parseJsObject(content: string): any {
  // 现在尝试解析为JSON
  content = content.trim();
  try {
    return JSON.parse(content);
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

export function getFileNameFromPath(absPath: string) {
  absPath = absPath.replaceAll("\\", "/");
  // 根据斜杠分割路径
  const pathArr = absPath.split("/");
  // 获取最后一个元素作为文件名
  const fileName = pathArr[pathArr.length - 1];
  return fileName;

  // 示例用法
  // const absPath = '/Users/username/Documents/file.txt';
  // const fileName = getFileNameFromPath(absPath);
  // console.log(fileName); // 输出 'file.txt'
}

export const IsEmptyString = (str: string) => {
  return str === "" || str === undefined || str === null;
};
export const IsEmptyValue = (value: any) => {
  return value === "" || value === undefined || value === null;
};

export function getNestedValue(obj: any, key: string) {
  return key.split(".").reduce((acc, curr) => {
    if (acc && typeof acc === "object" && curr in acc) {
      return acc[curr];
    } else {
      throw new Error("Invalid property path");
    }
  }, obj);
}

export function setNestedValue(obj: any, key: string, value: any) {
  const keys = key.split(".");
  let current = obj;

  for (let i = 0; i < keys.length - 1; i++) {
    if (
      !Object.prototype.hasOwnProperty.call(current, keys[i]) ||
      typeof current[keys[i]] !== "object"
    ) {
      current[keys[i]] = {};
    }
    current = current[keys[i]];
  }

  current[keys[keys.length - 1]] = value;
}
