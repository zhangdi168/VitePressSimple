<template>
  <!--  <div class="my-1 mx-2">-->
  <!--    <a-textarea-->
  <!--        v-model:value="title"-->
  <!--        placeholder="请输入标题"-->
  <!--        auto-size-->
  <!--    />-->
  <!--  </div>-->

  <div v-if="storeIndex.currArticlePath == ''">
    <a-empty description="未选择文章" />
  </div>
  <div
    v-show="storeIndex.currArticlePath != ''"
    id="vditor"
    ref="vditor"
    class="mx-2"
    style="height: 100vh; overflow-y: hidden"
  ></div>
</template>
<script lang="ts" setup>
import { defineProps, ref, onMounted, nextTick } from "vue";
import Vditor from "vditor";
import "vditor/dist/index.css";
import { useIndexStore } from "../../store";
import { DefaultVtitorOptions } from "@/configs/vditor";

const title = ref("");
// interface Props {
//   // 如果有props在这里定义
// }
//
// const props = defineProps<Props>();
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
