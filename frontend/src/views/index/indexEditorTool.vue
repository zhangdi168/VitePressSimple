<template>
  <div class="flex items-center justify-end pr-5">
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
</template>
<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";
import { useIndexStore } from "@/store";
import { useLayoutStore } from "@/store/layout";

const storeIndex = useIndexStore();
const storeLayout = useLayoutStore();
</script>
<style scoped></style>
