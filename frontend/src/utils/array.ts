//盘地暖一个数组是否为空或者长度为0
export function isEmptyArray(arr: any[] | null) {
  if (arr == null) return true;
  return arr.length === 0 || arr.length === undefined;
}
