import { getDirectoryPath } from "@/utils/file";
import { Move } from "../../wailsjs/go/services/ArticleTreeData";
import { ToastCheck } from "@/utils/Toast";

//移动文件
export const moveTo = async (path: string, target: string) => {
  const targetDir = getDirectoryPath(target);
  const result = await Move(path, targetDir);
  ToastCheck(result);
};
