export function addQuotesToUnquotedKeys(jsonString: string): string {
  // 匹配无引号的key：以字母、下划线或$开头，后面跟着0个或多个字母、数字、下划线或$
  return jsonString.replace(/(\w+):/g, function (match, p1) {
    // 将匹配到的无引号key加上双引号
    return `"${p1}":`;
  });
}
