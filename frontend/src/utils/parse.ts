export function extractJsContent(markdownText: string) {
  // 首先替换被 ``` 或 ` 包围的部分
  markdownText = markdownText.replace(/(```.*?```|`.+?`)/gs, "");

  // 使用正则表达式匹配 <script setup>{{jsContent}}</script>
  const regex = /<script\ssetup>(.*)<\/script>/g;

  const results = [];
  const match = regex.exec(markdownText);
  console.log(match, "-----match value");
  while (match !== null) {
    // 把匹配到的jsContent部分添加到结果数组中
    results.push(match[1]);
  }

  return results;
}
