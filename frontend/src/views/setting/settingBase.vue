<template>
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

  <div class="ml-5 mt-5">
    <config-radio
      :config-key="ConfigKeyVditorCdn"
      tooltip="如果出现markdown编辑器加载不出来的情况，可以手动切换编辑器cdn后【重启】程序"
      label="编辑器cdn"
      :items="[
        { label: 'unpkg', value: VditorCdnUnpkg },
        { label: 'zstatic', value: VditorCdnZstatic },
        { label: 'jsdeliver', value: VditorCdnJsdelivr },
      ]"
    ></config-radio>
  </div>
  <div class="ml-5 mt-5">
    <config-radio
      :config-key="ConfigKeySysUpdateSource"
      tooltip="中国用户推荐使用gitee源更新，非中国用户推荐使用github源更新"
      label="在线更新源"
      :items="[
        { label: 'github', value: 'github' },
        { label: 'gitee', value: 'gitee' },
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
<script setup lang="ts">
import {
  ConfigKeyChangeAutoSave,
  ConfigKeyFrontMatterSaveType,
  ConfigKeyIsStartup,
  ConfigKeySysUpdateSource,
  ConfigKeyVditorCdn,
} from "@/constant/keys/config";
import { useLayoutStore } from "@/store/layout";
import {
  GetSystemUserHomeDir,
  OpenFileBrowser,
  PathJoin,
} from "../../../wailsjs/go/system/SystemService";
import IconBtn from "@/components/iconBtn.vue";
import ConfigRadio from "@/components/configRadio.vue";
import ConfigSwitch from "@/components/configSwitch.vue";
import { UpdateNewVersion } from "../../../wailsjs/go/services/UpdateService";
import {
  VditorCdnJsdelivr,
  VditorCdnUnpkg,
  VditorCdnZstatic,
} from "@/constant/enums/cdn";

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
const checkUpdate = () => {
  UpdateNewVersion();
};
</script>
<style scoped></style>
