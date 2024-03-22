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

    <!--    切换日/夜间文字- -->

    <sim-input
      v-model="storeConfig.configData['lightModeSwitchTitle']"
      tooltip="自定义[悬停时]显示的浅色模式开关标题"
      placeholder=""
      label="悬停时浅色模式标题"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['darkModeSwitchTitle']"
      tooltip="自定义[悬停时]显示的深色模式开关标题"
      placeholder=""
      label="悬停时深色模式标题"
    ></sim-input>
    <!--    仅手机端生效-->
    <sim-input
      v-model="storeConfig.configData['darkModeSwitchLabel']"
      tooltip="自定义深色模式开关标签，该标签仅在移动端视图中显示"
      placeholder=""
      label="深色模式开关标签"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['sidebarMenuLabel']"
      tooltip="自定义侧边栏菜单标签，该标签仅在移动端视图中显示"
      placeholder=""
      label="自定义侧边栏菜单标签"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['returnToTopLabel']"
      tooltip="自定义返回顶部按钮的标签，该标签仅在移动端视图中显示"
      placeholder=""
      label="返回顶部按钮的标签"
    ></sim-input>

    <!--    多语言-->
    <sim-input
      v-model="storeConfig.configData['lang']"
      tooltip="站点的 lang 属性。这将呈现为页面 HTML 中的 <html lang='en-US'> 标签"
      placeholder="请输入项目标题后缀"
      label="html语言"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['langMenuLabel']"
      tooltip="自定义导航栏中语言切换按钮的 aria-label，仅当使用 i18n 时才使用此选项签"
      placeholder="请输入项目标题后缀"
      label="语言切换label"
    ></sim-input>

    <sim-switch
      v-model="storeConfig.configData['externalLinkIcon']"
      tooltip="是否在 markdown 中的外部链接旁显示外部链接图标"
      placeholder="请输入项目标题后缀"
      label="外部链接旁显示图标"
    ></sim-switch>
    <sim-input
      v-model="storeConfig.configData['themeConfig']['footer']['message']"
      tooltip="由于设计原因，仅当页面不包含侧边栏时才会显示页脚。"
      placeholder="请输入页脚提示信息"
      label="页脚信息"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['themeConfig']['footer']['copyright']"
      tooltip="由于设计原因，仅当页面不包含侧边栏时才会显示页脚。"
      placeholder="请输入页脚版权信息"
      label="页脚版权信息"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['themeConfig']['editLink']['text']"
      tooltip="编辑链接可让显示链接以编辑 Git 管理服务 (例如 GitHub 或 GitLab) 上的页面"
      placeholder=""
      label="编辑链接文本"
    ></sim-input>
    <sim-input
      v-model="storeConfig.configData['themeConfig']['editLink']['pattern']"
      tooltip="编辑链接可让显示链接以编辑 Git 管理服务 (例如 GitHub 或 GitLab) 上的页面"
      placeholder=""
      label="编辑链接地址"
    ></sim-input>

    <hr class="my-1 w-full" />
    <sim-bool-input
      class="w-full"
      v-model="storeConfig.configData['themeConfig']['docFooter']['prev']"
      input-label="上一页文本"
      bool-tooltip="用于全局启用或禁用上一页链接"
      bool-label="是否启用上一页连接"
      input-tooltip="用于自定义出现在上一页链接上方的文本"
    ></sim-bool-input>
    <sim-bool-input
      class="w-full"
      v-model="storeConfig.configData['themeConfig']['docFooter']['next']"
      input-label="下一页文本"
      bool-tooltip="用于全局启用或禁用下一页链接"
      bool-label="是否启用下一页连接"
      input-tooltip="用于自定义出现在下一页链接上方的文本"
    ></sim-bool-input>
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
      tooltip="指定放置生成的静态资源的目录。该路径应位于构建路径内，并相对于它进行解析。"
      label="生成静态目录"
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
      保存配置
    </a-button>
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
import { IconPark } from "@icon-park/vue-next/es/all";
import { onMounted, ref } from "vue";
import SimBoolInput from "@/components/simBoolInput.vue";

const storeConfig = useVpconfigStore();
const storeIndex = useIndexStore();
const isUsePrev = ref(false);
const isUseNext = ref(false);
const prevText = ref("");
const nextText = ref("");
onMounted(() => {});
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
