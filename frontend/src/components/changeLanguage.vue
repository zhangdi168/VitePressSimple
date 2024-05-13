<template>
  <!--    切换语言-->
  <div class="flex items-center">
    language：
    <div class="language border items-center flex justify-center">
      <divW
        v-for="item in languages"
        :key="item"
        :class="{ active: item === locale }"
        class="lang-item"
        @click="onclickLanguageHandle(item)"
      >
        {{ lang("languages." + item) }}
      </divW>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useI18n } from "vue-i18n";
import { lang } from "../utils/language";

const { t, availableLocales: languages, locale } = useI18n();

//设置语言
const onclickLanguageHandle = (item: string) => {
  item !== locale.value ? (locale.value = item) : false;
};
</script>
<style lang="scss" scoped>
.language {
  margin-right: 18px;
  border-radius: 8px;
  overflow: hidden;

  .lang-item {
    display: inline-block;
    min-width: 50px;
    height: 30px;
    line-height: 30px;
    padding: 0 5px;
    background-color: transparent;
    text-align: center;
    text-decoration: none;
    color: #000000;
    font-size: 14px;

    &:hover {
      background-color: #ff050542;
      cursor: pointer;
    }

    &.active {
      background-color: #6888d5;
      color: #ffffff;
      cursor: not-allowed;
    }
  }
}

.bar {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: flex-end;
  min-width: 150px;

  .bar-btn {
    display: inline-block;
    min-width: 80px;
    height: 30px;
    line-height: 30px;
    padding: 0 5px;
    margin-left: 8px;
    background-color: #ab7edc;
    border-radius: 2px;
    text-align: center;
    text-decoration: none;
    color: #000000;
    font-size: 14px;

    &:hover {
      background-color: #d7a8d8;
      color: #ffffff;
      cursor: pointer;
    }
  }
}
</style>
