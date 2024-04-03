//匹配script标签正则表达式
export const regexScript: RegExp = /<script\b[^>]*>([\s\S]*?)<\/script>/g;
//匹配style标签正则表达式
export const regexStyle: RegExp = /<style\b[^>]*>([\s\S]*?)<\/style>/g;
export const parseTagContent = (markdownText: string, regex: RegExp) => {
  // 首先替换被 ``` 包围的多行代码块
  markdownText = markdownText.replace(/```[^\n]*\n([\s\S]*?)\n```/gs, "");
  // 然后替换被 ` 包围的单行代码块
  markdownText = markdownText.replace(/`[^`]*`/g, "");
  // 使用正则表达式匹配非代码块内的script标签
  // const regex = /<script\b[^>]*>([\s\S]*?)<\/script>/g;
  const results = [];
  let match;
  while ((match = regex.exec(markdownText)) !== null) {
    // console.log(match, "-----match value");
    //match[0]是包含标签的字符串 match[1]是标签内容（不包含标签）
    results.push(match[0]); // 输出：alert("Hello, world!");
  }
  return results;
};
