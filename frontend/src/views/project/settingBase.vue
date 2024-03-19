<template>
  <div class="flex justify-start items-center my-2 mx-8">
    <!--    选择logo图片-->
    <div class="my-3 w-1/3 flex justify-between">
      <div class="flex-1">
        <a-input
          disabled
          v-model:value="storeConfig.configData['themeConfig']['logo']"
          placeholder="主页图片的链接"
          class="w-full"
        >
        </a-input>
      </div>
      <div class="mx-2">
        <a-button class="bg-blue-200" @click="selectLogo">
          <q-tooltip anchor="bottom left" self="bottom right"
            >logo图片默认存放于./images/{logo}
          </q-tooltip>
          选择站点LOGO
        </a-button>
      </div>
    </div>
    <sim-input
      v-model="storeConfig.configData['themeConfig']['siteTitle']"
      tooltip="即themeConfig.siteTitle 站点的标题，在logo后显示"
      placeholder="请输入站点标题"
      label="站点标题"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['title']"
      tooltip="即config.title 的标题"
      placeholder="请输入站点标题"
      label="页面标题"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['description']"
      tooltip="站点的描述。这将呈现为页面 HTML 中的 <meta> 标签"
      placeholder="请输入站点的描述"
      label="站点描述"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['titleTemplate']"
      tooltip="允许自定义每个页面的标题后缀或整个标题"
      placeholder="请输入项目标题后缀"
      label="标题后缀"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['lang']"
      tooltip="站点的 lang 属性。这将呈现为页面 HTML 中的 <html lang='en-US'> 标签"
      placeholder="请输入项目标题后缀"
      label="语言"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['srcDir']"
      tooltip="markdown 页面的目录，相对于项目根目录"
      label="文档路径"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['base']"
      tooltip="站点将部署到的base URL。如果计划在子路径页面下部署站点，则需要设置此项"
      placeholder="baseUrl"
      label="baseUrl"
    ></sim-input>

    <sim-input
      v-model="storeConfig.configData['outDir']"
      tooltip="项目的构建输出位置，相对于项目根目录"
      label="构建路径"
    ></sim-input>

    <sim-input
      v-model="storeConfig.configData['assetsDir']"
      tooltip="指定放置生成的静态资源的目录。该路径应位于 outDir 内，并相对于它进行解析。"
      label="生成静态路径"
    ></sim-input>

    <sim-input
      v-model="storeConfig.configData['cacheDir']"
      tooltip="缓存文件的目录，相对于项目根目录。"
      label="缓存路径"
    ></sim-input>
    <sim-switch
      v-model="storeConfig.configData['cleanUrls']"
      tooltip="当设置为 true 时，VitePress 将从 URL 中删除 .html 后缀"
      label="cleanUrls"
    ></sim-switch>
    <sim-switch
      v-model="storeConfig.configData['map']"
      tooltip="设置为 true 时，生产应用程序将在 MPA 模式下构建。MPA 模式默认提供 零 JavaScript 支持，代价是禁用客户端导航，并且需要明确选择加入才能进行交互。"
      label="map模式"
    ></sim-switch>
  </div>
</template>
<script setup lang="ts">
import { useVpconfigStore } from "@/store/vpconfig";
import SimInput from "@/components/simInput.vue";
import SimSwitch from "@/components/simSwitch .vue";
import {
  CopyPath,
  GetPathExt,
  GetPathFileName,
  PathJoin,
  SelectFile,
} from "../../../wailsjs/go/system/SystemService";
import { ToastCheck, ToastError } from "@/utils/Toast";
import { useIndexStore } from "@/store";

const storeConfig = useVpconfigStore();
const storeIndex = useIndexStore();

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
