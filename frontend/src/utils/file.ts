export function getDirectoryPath(filePath: string) {
  if (!filePath.endsWith(".md")) {
    return filePath;
  }
  // 先将所有反斜杠替换为正斜杠，以便统一处理
  const normalizedPath = filePath.replace("\\", "/");
  // 找到最后一个正斜杠的位置
  const lastSlashIndex = normalizedPath.lastIndexOf("/");

  // 截取到该位置的前缀字符串作为目录路径
  return lastSlashIndex !== -1
    ? normalizedPath.substring(0, lastSlashIndex)
    : "";
}
