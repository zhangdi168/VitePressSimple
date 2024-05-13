<template>
  <div class="ml-5 mt-5">
    <change-language></change-language>
  </div>
  <div class="ml-5 mt-5 flex justify-start">
    <config-switch
      :config-key="ConfigKeyChangeAutoSave"
      :label="lang('pageSetting.settingBase.autoSaveArticle.label')"
      :tooltip="lang('pageSetting.settingBase.autoSaveArticle.tooltip')"
    ></config-switch>
    <config-switch
      class="mx-3"
      :label="lang('pageSetting.settingBase.startup.label')"
      :tooltip="lang('pageSetting.settingBase.startup.tooltip')"
      :config-key="ConfigKeyIsStartup"
    ></config-switch>
  </div>
  <div class="ml-5 mt-5">
    <config-radio
      :config-key="ConfigKeyFrontMatterSaveType"
      :label="lang('pageSetting.settingBase.frontMatterSaveType.label')"
      :tooltip="lang('pageSetting.settingBase.frontMatterSaveType.tooltip')"
      :items="[
        { label: 'json', value: 'json' },
        { label: 'yaml', value: 'yaml' },
      ]"
    ></config-radio>
  </div>

  <div class="ml-5 mt-5">
    <config-radio
      :config-key="ConfigKeyVditorCdn"
      :label="lang('pageSetting.settingBase.editorCdn.label')"
      :tooltip="lang('pageSetting.settingBase.editorCdn.tooltip')"
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
      :label="lang('pageSetting.settingBase.sysUpdateSource.label')"
      :tooltip="lang('pageSetting.settingBase.sysUpdateSource.tooltip')"
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
      :text="lang('pageSetting.settingBase.checkUpdate')"
    ></icon-btn>
    <icon-btn
      @click="openDataDir()"
      btn-class="bg-orange"
      icon="seo-folder"
      :text="lang('pageSetting.settingBase.openDataDir')"
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
import { lang } from "@/utils/language";
import ChangeLanguage from "@/components/changeLanguage.vue";

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
