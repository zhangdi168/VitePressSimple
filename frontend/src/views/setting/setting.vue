<template>
  <!--  <div class="ml-5 mt-5">-->
  <!--    <config-color-->
  <!--      :config-key="ConfigKeyLayoutNavBgColor"-->
  <!--      label="左侧导航背景颜色"-->
  <!--    ></config-color>-->
  <!--  </div>-->
  <div class="ml-5 mt-5 flex justify-start">
    <config-switch
      label="切换文章自动保存:"
      tooltip="文章切换时自动保存"
      :config-key="ConfigKeyChangeAutoSave"
    ></config-switch>
    <config-switch
      class="mx-3"
      label="是否开机自启:"
      tooltip="重启程序后生效"
      :config-key="ConfigKeyIsStartup"
    ></config-switch>
  </div>
  <div class="ml-5 mt-5">
    <config-radio
      :config-key="ConfigKeyFrontMatterSaveType"
      tooltip="文章的frontMatter 保存方式"
      label="frontMatter保存方式"
      :items="[
        { label: 'json', value: 'json' },
        { label: 'yaml', value: 'yaml' },
      ]"
    ></config-radio>
  </div>
  <hr class="my-3" />
  <div class="flex items-center justify-center">
    <icon-btn
      class="mx-2"
      @click="checkUpdate()"
      icon="update-rotation"
      text="检查更新"
    ></icon-btn>
    <icon-btn
      @click="openDataDir()"
      btn-class="bg-orange"
      icon="seo-folder"
      text="打开程序数据目录"
    ></icon-btn>
  </div>
</template>
<script lang="ts" setup>
import ConfigSwitch from "@/components/configSwitch.vue";
import {
  ConfigKeyChangeAutoSave,
  ConfigKeyFrontMatterSaveType,
  ConfigKeyIsStartup,
  ConfigKeyLayoutNavBgColor,
} from "@/constant/keys/config";

import { useLayoutStore } from "@/store/layout";
import ConfigRadio from "@/components/configRadio.vue";
import IconBtn from "@/components/iconBtn.vue";
import {
  GetSystemUserHomeDir,
  OpenFileBrowser,
  PathJoin,
} from "../../../wailsjs/go/system/SystemService";

const storeLayout = useLayoutStore();
const saveSetting = () => {
  storeLayout.loadUserSetting();
};

//打开数据保存目录
const openDataDir = async () => {
  let homeDir = await GetSystemUserHomeDir();
  let dir = await PathJoin([homeDir, "VitePressSimple"]);
  await OpenFileBrowser(dir);
};

//检查更新
const checkUpdate = () => {};
</script>

<style scoped></style>
