<template>
  <div style="margin-top: 28vh" v-show="storeIndex.IsEmptyProject">
    <a-empty description="没有正在打开项目，先创建或者打开一个项目吧">
      <a-button @click="createProject" class="bg-green-300" type="default"
        >新建项目
      </a-button>
      <a-button @click="openProject" class="ml-4 bg-green-200" type="default"
        >打开项目
      </a-button>
    </a-empty>
  </div>
  <div>
    <create-project-popup ref="refCreateNewProject"></create-project-popup>
  </div>
  <div>
    <open-project-popup ref="refOpenNewProject"></open-project-popup>
  </div>
</template>
<script setup lang="ts">
import { isEmptyArray } from "@/utils/array";
import { useIndexStore } from "@/store";
import CreateProjectPopup from "@/views/project/createProjectPopup.vue";
import { ref } from "vue";
import OpenProjectPopup from "@/views/project/openProjectPopup.vue";
const storeIndex = useIndexStore();
//创建新项目
const refCreateNewProject = ref();
const createProject = () => {
  refCreateNewProject.value.showModal();
};

//打开项目
const refOpenNewProject = ref();
const openProject = () => {
  refOpenNewProject.value.showModal();
};
</script>
<style scoped></style>
