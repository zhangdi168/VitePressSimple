<template>
  <!--  <div class="my-1 mx-2">-->
  <!--    <a-textarea-->
  <!--        v-model:value="title"-->
  <!--        placeholder="请输入标题"-->
  <!--        auto-size-->
  <!--    />-->
  <!--  </div>-->

  <div
    v-if="
      !isEmptyArray(storeIndex.articleTreeData) &&
      storeIndex.currArticlePath == ''
    "
  >
    <a-empty description="未选择文章" />
  </div>
  <div
    style="margin-top: 28vh"
    v-show="isEmptyArray(storeIndex.articleTreeData)"
  >
    <a-empty description="先创建或者打开一个项目吧">
      <!--      <template #description>-->
      <!--        <span>-->
      <!--          <a href="#api">Description</a>-->
      <!--        </span>-->
      <!--      </template>-->
      <a-button @click="createProject" class="bg-green-400" type="default"
        >新建项目
      </a-button>
      <a-button
        @click="storeIndex.selectProjectDir()"
        class="ml-4 bg-green-200"
        type="default"
        >打开项目
      </a-button>
    </a-empty>
  </div>
  <div
    v-show="storeIndex.currArticlePath != ''"
    id="vditor"
    ref="vditor"
    class="mx-2"
    style="height: 100vh; overflow-y: hidden"
  ></div>

  <div>
    <create-project-popup ref="refCreateNewProject"></create-project-popup>
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

const title = ref("");
// interface Props {
//   // 如果有props在这里定义
// }
//
// const props = defineProps<Props>();
const storeIndex = useIndexStore();
const contentEditor = ref<Vditor>();

//创建新项目
const refCreateNewProject = ref();
const createProject = () => {
  refCreateNewProject.value.showModal();
};

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
