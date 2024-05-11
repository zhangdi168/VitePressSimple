<template>
  <div class="flex justify-start items-center my-2 mx-8">
    <!--    选择logo图片-->
    <div class="my-3 w-1/3 flex justify-between">
      <div class="flex-1">
        <a-input
          disabled
          v-model:value="storeConfig.configData['themeConfig']['logo']"
          :placeholder="lang('pageProject.settingBase.placeholders.logoUrl')"
          class="w-full"
        >
        </a-input>
      </div>

      <div class="mx-2">
        <a-button class="bg-blue-200" @click="selectLogo">
          <q-tooltip anchor="bottom left" self="bottom right">
            {{ lang("pageProject.settingBase.logoSavePath") }}
          </q-tooltip>
          {{ lang("pageProject.settingBase.selectLogo") }}
        </a-button>
      </div>
    </div>

    <sim-input
      v-model="storeConfig.configData['title']"
      :tooltip="lang('pageProject.settingBase.tooltips.pageTitle')"
      :placeholder="lang('pageProject.settingBase.placeholders.pageTitle')"
      :label="lang('pageProject.settingBase.labels.pageTitle')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['titleTemplate']"
      :tooltip="lang('pageProject.settingBase.tooltips.titleSuffix')"
      :placeholder="lang('pageProject.settingBase.placeholders.titleSuffix')"
      :label="lang('pageProject.settingBase.labels.titleSuffix')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['themeConfig']['outline']['level']"
      :tooltip="lang('pageProject.settingBase.tooltips.outlineLevel')"
      :placeholder="lang('pageProject.settingBase.placeholders.outlineLevel')"
      :label="lang('pageProject.settingBase.labels.outlineLevel')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['lang']"
      :tooltip="lang('pageProject.settingBase.tooltips.htmlLang')"
      :placeholder="lang('pageProject.settingBase.placeholders.htmlLang')"
      :label="lang('pageProject.settingBase.labels.htmlLang')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['description']"
      :tooltip="lang('pageProject.settingBase.tooltips.siteDescription')"
      :placeholder="
        lang('pageProject.settingBase.placeholders.siteDescription')
      "
      :label="lang('pageProject.settingBase.labels.siteDescription')"
    ></sim-input>

    <!--    切换日/夜间文字- -->

    <!--    仅手机端生效-->

    <!--    多语言-->

    <sim-switch
      v-model="storeConfig.configData['themeConfig']['externalLinkIcon']"
      :tooltip="lang('pageProject.settingBase.tooltips.externalLinkIcon')"
      :placeholder="lang('pageProject.settingBase.placeholders.titleSuffix')"
      :label="lang('pageProject.settingBase.labels.externalLinkIcon')"
    ></sim-switch>

    <hr class="my-1 w-full" />

    <sim-input
      v-model="storeConfig.configData['srcDir']"
      :tooltip="lang('pageProject.settingBase.tooltips.docPath')"
      :label="lang('pageProject.settingBase.labels.docPath')"
    ></sim-input>

    <sim-input
      v-model="storeConfig.configData['base']"
      :tooltip="lang('pageProject.settingBase.tooltips.baseUrl')"
      :placeholder="lang('pageProject.settingBase.placeholders.baseUrl')"
      :label="lang('pageProject.settingBase.labels.baseUrl')"
    ></sim-input>

    <sim-input
      v-model="storeConfig.configData['outDir']"
      :tooltip="lang('pageProject.settingBase.tooltips.buildPath')"
      :label="lang('pageProject.settingBase.labels.buildPath')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['sitemap']['hostname']"
      :tooltip="lang('pageProject.settingBase.tooltips.hostname')"
      :label="lang('pageProject.settingBase.labels.hostname')"
    ></sim-input>

    <sim-input
      v-model="storeConfig.configData['assetsDir']"
      :tooltip="lang('pageProject.settingBase.tooltips.staticAssets')"
      :label="lang('pageProject.settingBase.labels.staticAssets')"
    ></sim-input>

    <sim-input
      v-model="storeConfig.configData['cacheDir']"
      :tooltip="lang('pageProject.settingBase.tooltips.cachePath')"
      :label="lang('pageProject.settingBase.labels.staticAssets')"
    ></sim-input>
    <sim-switch
      v-model="storeConfig.configData['cachePath']"
      :tooltip="lang('pageProject.settingBase.tooltips.cleanUrls')"
      label="cleanUrls"
    ></sim-switch>
    <!--    <sim-switch-->
    <!--      v-model="storeConfig.configData['map']"-->
    <!--      tooltip="设置为 true 时，生产应用程序将在 MPA 模式下构建。MPA 模式默认提供 零 JavaScript 支持，代价是禁用客户端导航，并且需要明确选择加入才能进行交互。"-->
    <!--      label="map模式"-->
    <!--    ></sim-switch>-->
  </div>
  <hr class="my-2" />
  <div class="flex justify-center">
    <a-button
      @click="saveBaseConfig"
      class="bg-blue-600 hover:bg-blue-500 text-white flex justify-center items-center"
    >
      <icon-park
        class="mr-1"
        strokeLinejoin="bevel"
        theme="outline"
        type="save"
      />
      {{ lang("common.saveConfig") }}
    </a-button>
  </div>
</template>
<script setup lang="ts">
import { useVpconfigStore } from "@/store/vpconfig";
import SimInput from "@/components/simInput.vue";
import SimSwitch from "@/components/simSwitch.vue";
import {
  CopyPath,
  GetPathExt,
  PathJoin,
  SelectFile,
} from "../../../wailsjs/go/system/SystemService";
import { ToastCheck, ToastError } from "@/utils/Toast";
import { useIndexStore } from "@/store";
import { IconPark } from "@icon-park/vue-next/es/all";
import { onBeforeMount, onMounted, ref } from "vue";
import { lang } from "@/utils/language";

const storeConfig = useVpconfigStore();
const storeIndex = useIndexStore();
onBeforeMount(() => {});
const saveBaseConfig = () => {
  storeConfig.saveConfig();
};
const selectLogo = async () => {
  let oriImagePath = await SelectFile("选择主页图片", "");
  console.log(oriImagePath, "filePath -- console.log");
  let ext = await GetPathExt(oriImagePath);
  let allowExt = [".png", ".jpg", ".jpeg", ".bmp"];
  if (!allowExt.includes(ext)) {
    ToastError("请选则图片格式文件" + allowExt);
    return;
  }
  //组装新路径
  let publicDir = await PathJoin([storeConfig.fullSrcDir, "public"]);
  let newImagePath = await PathJoin([publicDir, "images", "logo" + ext]);
  let copyResult = await CopyPath(oriImagePath, newImagePath, false);
  ToastCheck(copyResult);
  storeConfig.configData["themeConfig"]["logo"] = newImagePath.replaceAll(
    publicDir,
    "",
  );
};
</script>
<style scoped></style>
