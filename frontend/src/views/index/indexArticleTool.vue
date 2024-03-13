<template>
  <!-- 头部-添加文章等 -->
  <div class="px-2 py-2 border-b">
    <div class="flex items-center justify-start" style="overflow-y: hidden">
      <div @click="createDir" class="mx-1 cursor-pointer">
        <!--        <q-tooltip>新建文件(根目录)</q-tooltip>-->
        <a-dropdown arrow>
          <a class="ant-dropdown-link" @click.prevent>
            <icon-park
              :size="iconSize"
              fill="#333"
              strokeLinejoin="bevel"
              theme="outline"
              type="add"
            />
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item @click="createFile()">
                <a href="javascript:;">新建文件</a>
              </a-menu-item>
              <a-menu-item @click="createDir()">
                <a href="javascript:;">新建文件夹</a>
              </a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </div>

      <div class="mx-1 cursor-pointer">
        <q-tooltip>刷新目录</q-tooltip>
        <icon-park
          :size="iconSize"
          fill="#333"
          strokeLinejoin="bevel"
          theme="outline"
          type="refresh"
          @click="storeIndex.loadTreeData()"
        />
      </div>

      <div class="mx-1 cursor-pointer">
        <q-tooltip>刷新目录</q-tooltip>
        <icon-park
          :size="iconSize"
          fill="#5f4364"
          strokeLinejoin="bevel"
          theme="outline"
          type="search"
          @click="storeIndex.loadTreeData()"
        />
      </div>
    </div>
    <input-model
      ref="refModalCreateDir"
      placeholder="请输入文件夹名称"
      @submit-input-modal="onSubmitInputModalCreateDir"
    ></input-model>
    <input-model
      ref="refModalCreateFile"
      placeholder="请输入文件名称"
      @submit-input-modal="onSubmitInputModalCreateFile"
    ></input-model>
  </div>
</template>
<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";

import { useIndexStore } from "@/store";
import { nextTick, onMounted, ref } from "vue";
import { FileSearchOutlined } from "@ant-design/icons-vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { ConfigGet } from "../../../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import InputModel from "@/components/inputModel.vue";
import { getParentDirectory } from "@/utils/file";
import {
  CreateDir,
  CreateFile,
  Rename,
} from "../../../wailsjs/go/services/ArticleTreeData";
import { ToastCheck } from "@/utils/Toast";

const iconSize = ref(24);
const storeIndex = useIndexStore();
const storeVpConfig = useVpconfigStore();

//弹出层回调 新建目录
const refModalCreateDir = ref();
const createDir = () => {
  refModalCreateDir.value.showModal("");
};
const onSubmitInputModalCreateDir = async (value: string) => {
  let fullDir = storeVpConfig.FullSrcDir + "/" + value;
  console.log(fullDir, "-----fullDir value");
  let res = await CreateDir(fullDir);
  if (ToastCheck(res, "文件夹创建完成")) await storeIndex.loadTreeData();
};

//弹出层回调 新建文件
const refModalCreateFile = ref();
const createFile = async () => {
  refModalCreateFile.value.showModal("");
};
const onSubmitInputModalCreateFile = async (value: string) => {
  let fullFile = storeVpConfig.FullSrcDir + "/" + value;
  console.log(fullFile, "-----fullFile value");
  let res = await CreateFile(fullFile);
  if (ToastCheck(res, "文件创建完成")) await storeIndex.loadTreeData();
};
</script>
<style scoped></style>
