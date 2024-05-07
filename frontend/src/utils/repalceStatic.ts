// 定义辅助工具函数
import { useIndexStore } from "@/store";
import { useVpconfigStore } from "@/store/vpconfig";

export function escapeRegExp(str: string): string {
  return str.replace(/[.*+?^${}()|[\]\\]/g, "\\$&"); // 转义正则特殊字符
}

// 第一个函数：将域名转成字符串 "/vpstatic/"
export function replaceImageUrlToLocalStatic(markdownText: string): string {
  const vpstatic = getReplaceVpstaticValue();
  const originalBaseUrl = getReplaceStaticBaseUrl();

  // const regex = new RegExp(escapeRegExp(originalBaseUrl) + "(images/.*)", "g");

  // return markdownText.replaceAll(regex, vpstatic + "$1");
  return markdownText.replaceAll(originalBaseUrl, vpstatic);
}

// 第二个函数：将字符串 "/vpstatic/" 替换成域名 "http://localhost:9874/"
export function replaceLocalStaticToImageUrl(markdownText: string): string {
  // return markdownText;

  const vpstatic = getReplaceVpstaticValue();
  const originalBaseUrl = getReplaceStaticBaseUrl();

  const regex = new RegExp(escapeRegExp(vpstatic) + "(images/.*)", "g");

  // return markdownText.replaceAll(regex, originalBaseUrl + "$1");
  return markdownText.replaceAll(vpstatic, originalBaseUrl);
}

const getReplaceVpstaticValue = () => {
  const storeIndex = useIndexStore();
  const storeVpConfig = useVpconfigStore();
  let vpstatic = "";
  if (storeVpConfig.srcDir == "./") {
    //原目录在项目根目录
    vpstatic = "](/" + storeIndex.staticBaseDir + "/";
  } else {
    //原目录不在项目根目录 如./docs
    vpstatic = "](/../" + storeIndex.staticBaseDir + "/";
  }

  return vpstatic;
};

const getReplaceStaticBaseUrl = () => {
  const storeIndex = useIndexStore();
  const port = storeIndex.staticServerPort;
  return "](http://localhost:" + port + "/";
};
