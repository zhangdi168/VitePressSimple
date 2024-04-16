<template>
  <!-- 头部-添加文章等 -->
  <div class="px-2 py-2 border-b" v-show="!storeIndex.IsEmptyProject">
    <div class="flex items-center justify-start" style="overflow-y: hidden">
      <div class="mx-1 cursor-pointer">
        <!--        <q-tooltip>新建文件(根目录)</q-tooltip>-->
        <a-dropdown arrow>
          <a class="ant-dropdown-link" @click.prevent>
            <icon-park
              :size="iconSize"
              fill="black"
              strokeLinejoin="bevel"
              theme="outline"
              type="add"
            />
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item @click="createProject()">
                <a href="javascript:;">新建项目</a>
              </a-menu-item>
              <a-menu-item @click="createFile()">
                <a href="javascript:;">新建文章</a>
              </a-menu-item>
              <a-menu-item @click="createDir()">
                <a href="javascript:;">新建目录</a>
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </div>
      <div class="mx-1 cursor-pointer" @click="openProject()">
        <q-tooltip>打开现有项目</q-tooltip>
        <a class="ant-dropdown-link" @click.prevent>
          <icon-park
            :size="iconSize"
            fill="#333"
            strokeLinejoin="bevel"
            theme="outline"
            type="layers"
          />
        </a>
      </div>
      <div
        class="mx-1 cursor-pointer"
        @click="OpenFileBrowser(storeVpConfig.baseDir)"
      >
        <q-tooltip>在文件浏览器中打开</q-tooltip>
        <icon-park
          :size="iconSize"
          fill="#333"
          strokeLinejoin="bevel"
          theme="outline"
          type="folder-close"
          @click="storeIndex.loadTreeData()"
        />
      </div>
      <div class="mx-1 cursor-pointer">
        <q-tooltip>刷新目录树</q-tooltip>
        <icon-park
          :size="iconSize"
          fill="#333"
          strokeLinejoin="bevel"
          theme="outline"
          type="refresh"
          @click="storeIndex.loadTreeData()"
        />
      </div>
    </div>

    <!--  弹窗-->
    <div>
      <create-project-popup ref="refCreateNewProject"></create-project-popup>
    </div>
    <div>
      <open-project-popup ref="refOpenNewProject"></open-project-popup>
    </div>

    <input-model
      ref="refModalCreateDir"
      placeholder="请输入文件夹名称"
      title="请输入文件夹名称"
      @submit-input-modal="onSubmitInputModalCreateDir"
    ></input-model>
    <input-model
      ref="refModalCreateFile"
      title="请输入文章标题"
      placeholder="请输入文章名称,不带.md后缀"
      @submit-input-modal="onSubmitInputModalCreateFile"
    ></input-model>
  </div>
</template>
<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";

import { useIndexStore } from "@/store";
import { ref } from "vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { OpenFileBrowser } from "../../../wailsjs/go/system/SystemService";
import InputModel from "@/components/inputModel.vue";
import {
  CreateDir,
  CreateFile,
} from "../../../wailsjs/go/services/ArticleTreeData";
import { ToastCheck } from "@/utils/Toast";
import CreateProjectPopup from "@/views/project/createProjectPopup.vue";
import OpenProjectPopup from "@/views/project/openProjectPopup.vue";

const iconSize = ref(24);
const storeIndex = useIndexStore();
const storeVpConfig = useVpconfigStore();

//弹出层回调 新建目录
const refModalCreateDir = ref();
const createDir = () => {
  refModalCreateDir.value.showModal("");
};
const onSubmitInputModalCreateDir = async (value: string) => {
  let fullDir = storeVpConfig.fullSrcDir + "/" + value;
  let res = await CreateDir(fullDir);
  if (ToastCheck(res, "文件夹创建完成")) await storeIndex.loadTreeData();
};

//弹出层回调 新建文件
const refModalCreateFile = ref();
const createFile = async () => {
  refModalCreateFile.value.showModal("");
};
const onSubmitInputModalCreateFile = async (value: string) => {
  let fullFile = storeVpConfig.fullSrcDir + "/" + value + ".md";
  let res = await CreateFile(fullFile);
  if (ToastCheck(res, "文件创建完成")) await storeIndex.loadTreeData();
};

//新建项目
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
