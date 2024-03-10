<template>
  <div class="flex items-center justify-end pr-5">
    <a-tooltip class="mx-1 cursor-pointer">
      <template #title>front matter配置</template>
      <icon-park
        :size="28"
        fill="blue"
        strokeLinejoin="bevel"
        theme="outline"
        type="seo"
      />
    </a-tooltip>

    <a-tooltip
      class="mx-1 cursor-pointer"
      @click="storeIndex.saveCurrArticle()"
    >
      <template #title>保存 ctrl+s</template>
      <icon-park
        :size="28"
        fill="black"
        strokeLinejoin="bevel"
        theme="outline"
        type="save"
      />
    </a-tooltip>
  </div>
</template>
<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";
import { useIndexStore } from "@/store";

const storeIndex = useIndexStore();
</script>
<style scoped></style>
