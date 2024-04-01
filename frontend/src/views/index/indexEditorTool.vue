<template>
  <div
    class="flex items-center justify-between pr-5"
    v-show="!storeIndex.IsEmptyProject"
  >
    <div class="text-gray-400 ml-1.5">当前项目:{{ storeConfig.baseDir }}</div>
    <div class="flex items-center justify-end">
      <a-tooltip class="mx-1 cursor-pointer">
        <template #title
          >{{ storeLayout.showEditorView ? "隐藏" : "显示" }}页面属性编辑栏
        </template>
        <icon-park
          class="select-none"
          @click="storeLayout.setEditorViewShow(false)"
          v-show="storeLayout.showEditorView"
          :size="storeLayout.editorToolIconSize"
          fill="#493c3c"
          strokeLinejoin="bevel"
          theme="outline"
          type="file-display"
        />
        <icon-park
          class="select-none"
          @click="storeLayout.setEditorViewShow(true)"
          v-show="!storeLayout.showEditorView"
          :size="storeLayout.editorToolIconSize"
          fill="#493c3c"
          strokeLinejoin="bevel"
          theme="outline"
          type="file-hiding"
        />
      </a-tooltip>

      <a-tooltip
        class="mx-1 cursor-pointer"
        @click="storeIndex.saveCurrArticle()"
      >
        <template #title>保存 ctrl+s</template>
        <icon-park
          :size="storeLayout.editorToolIconSize"
          fill="#493c3c"
          strokeLinejoin="bevel"
          theme="outline"
          type="save"
        />
      </a-tooltip>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";
import { useIndexStore } from "@/store";
import { useLayoutStore } from "@/store/layout";
import { useVpconfigStore } from "@/store/vpconfig";

const storeIndex = useIndexStore();
const storeConfig = useVpconfigStore();
const storeLayout = useLayoutStore();
</script>
<style scoped></style>
