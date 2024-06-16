<template>
  <div
    class="flex items-center justify-between pr-5"
    v-show="!storeIndex.IsEmptyProject"
    :style="StyleNoDrag"
  >
    <div class="flex items-center justify-start">
      <!--    终端-->
      <index-editor-tool-shell></index-editor-tool-shell>
      <!-- 隐藏或关闭右侧栏-->
      <a-tooltip class="mx-1 cursor-pointer">
        <template #title
          >{{
            storeLayout.showEditorView
              ? lang("common.hide")
              : lang("common.show")
          }}{{ lang("pageIndex.pageProperties") }}
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

      <!--      保存-->
      <a-tooltip
        class="mx-1 cursor-pointer"
        @click="storeIndex.saveCurrArticle()"
      >
        <template #title>{{ lang("common.saveWithKey") }}</template>
        <icon-park
          :size="storeLayout.editorToolIconSize"
          fill="#493c3c"
          strokeLinejoin="bevel"
          theme="outline"
          type="save"
        />
      </a-tooltip>
    </div>

    <div class="text-gray-400 ml-1.5">
      {{ lang("pageIndex.currentProject") }}{{ storeConfig.baseDir }}
    </div>
  </div>
</template>
<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";
import { useIndexStore } from "@/store";
import { useLayoutStore } from "@/store/layout";
import { useVpconfigStore } from "@/store/vpconfig";
import { lang } from "../../utils/language";
import IndexEditorToolShell from "@/views/index/indexEditorToolShell.vue";
import { StyleNoDrag } from "@/configs/cnts";

const storeIndex = useIndexStore();
const storeConfig = useVpconfigStore();
const storeLayout = useLayoutStore();
</script>
<style scoped></style>
