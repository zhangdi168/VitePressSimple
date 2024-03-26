//递归深拷贝 支持含义function的对象
export function DeepClone(obj: Record<string, any>): Record<string, any> {
  if (typeof obj !== "object" || obj === null) {
    return obj;
  }

  let cloneObj;

  // 判断对象是数组还是普通对象，创建相应的新对象
  if (Array.isArray(obj)) {
    cloneObj = [];
  } else {
    cloneObj = {};
  }

  for (const key in obj) {
    // 使用 hasOwnProperty 方法判断 key 是否是自身属性
    if (Object.prototype.hasOwnProperty.call(obj, key)) {
      // @ts-ignore
      cloneObj[key] = DeepClone(obj[key]); // 递归调用深拷贝函数
    }
  }

  return cloneObj;
}
