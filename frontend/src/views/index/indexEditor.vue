<script lang="ts" setup>
import { nextTick, onMounted, ref } from "vue";
import IndexEditorTool from "./indexEditorTool.vue";
import IndexEditorWriter from "./indexEditorWriter.vue";
import IndexEditorReader from "./indexEditorView.vue";
import { useLayoutStore } from "@/store/layout";

const storeLayout = useLayoutStore();
// const refSplitterEditor = ref();
// onMounted(() => {
//   nextTick(() => {
//   });
// });
</script>

<template>
  <div style="height: 100vh; overflow-y: hidden">
    <!--    顶部工具栏-->
    <index-editor-tool
      class="sticky z-10 top-0 py-1.5 border-b w-full"
    ></index-editor-tool>

    <!--    编辑器写作和预览区-->
    <q-splitter
      ref="refSplitterEditor"
      :unit="'%'"
      style="height: 100vh"
      :model-value="storeLayout.EditorLeftWriterWidth"
    >
      <template v-slot:before>
        <!--编辑器写作区，确保内容区域高度自适应-->
        <index-editor-writer></index-editor-writer>
      </template>

      <template v-slot:after>
        <!--编辑器预览区，确保内容区域高度自适应-->
        <index-editor-reader></index-editor-reader>
      </template>
    </q-splitter>

    <!--    编辑器底部统计区-->
    <!--    <index-editor-footer style="height: 33px"></index-editor-footer>-->
  </div>
</template>

<style scoped></style>
