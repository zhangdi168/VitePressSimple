<template>
  <div
    v-show="storeIndex.currArticlePath != ''"
    id="vditor"
    ref="vditor"
    class="mx-2"
    style="height: 100vh; overflow-y: hidden"
  ></div>
  <empty-project></empty-project>
  <div
    v-if="
      !isEmptyArray(storeIndex.articleTreeData) &&
      storeIndex.currArticlePath == ''
    "
  >
    <a-empty description="未选择文章" />
  </div>
</template>
<script lang="ts" setup>
import { defineProps, ref, onMounted, nextTick } from "vue";
import Vditor from "vditor";
import "vditor/dist/index.css";
import { useIndexStore } from "../../store";
import { DefaultVtitorOptions } from "@/configs/vditor";
import { isEmptyArray } from "@/utils/array";
import CreateProjectPopup from "@/views/project/createProjectPopup.vue";
import EmptyProject from "@/components/emptyProject.vue";

const title = ref("");

const storeIndex = useIndexStore();
const contentEditor = ref<Vditor>();

onMounted(() => {
  nextTick(() => {
    contentEditor.value = new Vditor("vditor", DefaultVtitorOptions);
    storeIndex.setVditorInstance(contentEditor.value);
  });
});
</script>
<style scoped>
a {
  color: #42b983;
}
</style>
