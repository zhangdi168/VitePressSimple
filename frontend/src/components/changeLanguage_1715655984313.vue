template
<template>
  <!-- 使用useI18n钩子函数获取国际化相关功能 -->
  const { t, availableLocales: languages, locale } = useI18n(); /** *
  处理点击语言项的事件 * @param {string} item - 被点击的语言项 * @description
  根据点击的语言项来切换当前的语言 */ const onclickLanguageHandle = (item:
  string) => { // 如果点击的语言项不等于当前语言，则更新为点击的语言项 item !==
  locale.value ? (locale.value = item) : false; };
</template>
