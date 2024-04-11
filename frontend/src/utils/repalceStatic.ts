// 定义辅助工具函数
import { useIndexStore } from "@/store";

export function escapeRegExp(str: string): string {
  return str.replace(/[.*+?^${}()|[\]\\]/g, "\\$&"); // 转义正则特殊字符
}

// 第一个函数：将域名转成字符串 "./vpstatic/"
export function replaceImageUrlToLocalStatic(markdownText: string): string {
  // const newBaseUrl = "./vpstatic/";
  const storeIndex = useIndexStore();
  const port = storeIndex.staticServerPort;
  const vpstatic = "/" + storeIndex.staticBaseDir + "/";
  // const port = await ConfigGet(ConfigKeySysStaticServerPort);
  // const newBaseUrl = await ConfigGet(ConfigKeySysProjectStaticDirName);
  const originalBaseUrl = "http://localhost:" + port + "/";
  const regex = new RegExp(escapeRegExp(originalBaseUrl) + "(images/.*)", "g");

  return markdownText.replaceAll(regex, vpstatic + "$1");
  // return markdownText.replaceAll(originalBaseUrl, vpstatic);
}

// 第二个函数：将字符串 "./vpstatic/" 替换成域名 "http://localhost:9874/"
export function replaceLocalStaticToImageUrl(markdownText: string): string {
  const storeIndex = useIndexStore();
  const port = storeIndex.staticServerPort;
  const vpstatic = "/" + storeIndex.staticBaseDir + "/";
  // const localStaticPath = "./vpstatic/";
  const originalBaseUrl = "http://localhost:" + port + "/";

  const regex = new RegExp(escapeRegExp(vpstatic) + "(images/.*)", "g");

  return markdownText.replaceAll(regex, originalBaseUrl + "$1");
  // return markdownText.replaceAll(vpstatic, originalBaseUrl);
}
