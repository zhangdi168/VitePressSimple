export function getDirectoryPath(path: string) {
  if (!path.endsWith(".md")) return path;
  // 验证并处理 Windows 系统下的盘符
  const hasDriveLetter = /^[a-zA-Z]:\\/i.test(path);
  let driveLetter = "";

  if (hasDriveLetter) {
    const match = path.match(/^[a-zA-Z]:/i);
    driveLetter = match![0];
    path = path.slice(driveLetter.length);
  }

  // 分离路径中的目录和文件部分
  const parts = path.split(/\/|\\/);

  // 如果路径是根目录或者仅包含一个单独的驱动器号，则返回原路径
  if (parts.length === 1 || (parts.length === 2 && !parts[1])) {
    return (
      driveLetter +
      (path.endsWith("/") || path.endsWith("\\") ? path : path + "/")
    );
  }

  // 重新构建并返回目录部分
  let directoryPath = driveLetter;
  for (let i = 0; i < parts.length - 1; i++) {
    directoryPath += parts[i] + "/";
  }
  return directoryPath;
}
export function getParentDirectory(path: string): string {
  // 验证并处理 Windows 系统下的盘符
  const hasDriveLetter = /^[a-zA-Z]:\\/i.test(path);
  let driveLetter = "";

  if (hasDriveLetter) {
    const match = path.match(/^[a-zA-Z]:/i);
    driveLetter = match![0];
    path = path.slice(driveLetter.length);
  }

  // 分离路径中的目录和文件部分
  const parts = path.split(/\/|\\/);

  // 如果路径是根目录或者仅包含一个单独的驱动器号，则返回原路径的父级（对于根目录而言，父级还是它本身）
  if (parts.length === 1 || (parts.length === 2 && !parts[1])) {
    return (
      driveLetter +
      (path.endsWith("/") || path.endsWith("\\") ? path : path + "/") +
      ".."
    );
  }

  // 对于非根目录的情况，构建并返回上级目录路径
  let parentDirectoryPath = driveLetter;
  const isFile = path.endsWith(".md");
  const reduceNum = isFile ? 1 : 0;
  for (let i = 0; i < parts.length - 1; i++) {
    // 注意这里是减1，因为我们要获取上级目录
    parentDirectoryPath += parts[i] + "/";
  }

  // 如果路径不是以斜杠或反斜杠结尾，添加相应的路径分隔符
  // if (
  //   !parentDirectoryPath.endsWith("/") &&
  //   !parentDirectoryPath.endsWith("\\")
  // ) {
  //   parentDirectoryPath += "/";
  // }

  // 添加表示上级目录的 '..'
  //parentDirectoryPath += "..";

  return parentDirectoryPath.replace(/\$/, "").replace(/\$/, "");
}

// 删除md后缀
export function removeMdExtension(str: string) {
  return str.replace(/\.md$/, "");
}
