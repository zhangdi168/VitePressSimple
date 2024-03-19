<template>
  <div>
    <div class="flex justify-start items-center">
      <div class="ml-1">页面类型:</div>

      <q-radio
        v-model="storeIndex.currArticleFrontMatter['layout']"
        val="doc"
        label="文档"
      />
      <q-radio
        v-model="storeIndex.currArticleFrontMatter['layout']"
        val="home"
        label="主页"
      />
      <q-radio
        v-model="storeIndex.currArticleFrontMatter['layout']"
        disable
        val="page"
        label="自定义"
      />
    </div>

    <div
      class="mt-2"
      v-if="storeIndex.currArticleFrontMatter['layout'] == 'home'"
    >
      <div class="px-1 mb-2">
        <a-input
          v-model:value="storeIndex.currArticleFrontMatter['hero']['name']"
          placeholder="请输入名称"
          prefix=""
          suffix="名称"
        />
      </div>
      <div class="px-1 mb-2">
        <a-input
          v-model:value="storeIndex.currArticleFrontMatter['hero']['text']"
          placeholder="请输入简介"
          prefix=""
          suffix="简介"
        />
      </div>
      <div class="px-1 mb-2">
        <a-input
          v-model:value="storeIndex.currArticleFrontMatter['hero']['tagline']"
          placeholder="`text` 下方的标语"
          prefix=""
          suffix="标语"
        />
      </div>
      <div class="my-3 flex justify-between">
        <div class="flex-1">
          <a-input
            disabled
            v-model:value="
              storeIndex.currArticleFrontMatter['hero']['image']['src']
            "
            placeholder="主页图片的 alt 属性"
            class="w-full"
          >
          </a-input>
        </div>
        <div class="mx-2">
          <a-button class="bg-blue-200" @click="selectLogo">
            <q-tooltip anchor="bottom left" self="bottom right"
              >主页图片默认存放于./images/home/{文件名}
            </q-tooltip>
            选择主页图片
          </a-button>
        </div>
      </div>
      <hr class="my-3" />
      <div>
        <dy-add-k-v
          v-model:arr="storeIndex.currArticleFrontMatter['hero']['actions']"
          add-btn-text="添加主页按钮"
          value-placeholder="跳转链接"
          add-btn-class="bg-orange-200"
          key-placeholder="文本"
          key-name="text"
          value-name="link"
        ></dy-add-k-v>
      </div>
      <hr class="my-3" />
      <div>
        <dy-add-k-v
          v-model:arr="storeIndex.currArticleFrontMatter['features']"
          add-btn-text="添加features"
          value-placeholder="功能详情"
          add-btn-class="bg-orange-200"
          key-placeholder="功能标题"
          key-name="title"
          value-name="details"
        ></dy-add-k-v>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { useIndexStore } from "@/store";
import IndexPopAddFeatures from "@/views/index/indexPopAddFeatures.vue";
import { ref } from "vue";
import DyAddKV from "@/components/dyAddKV.vue";
import {
  CopyPath,
  GetPathExt,
  GetPathFileName,
  PathJoin,
  SelectFile,
} from "../../../wailsjs/go/system/SystemService";
import { ToastCheck, ToastError } from "@/utils/Toast";
import { useVpconfigStore } from "@/store/vpconfig";

const storeIndex = useIndexStore();
const storeVpConfig = useVpconfigStore();

const selectLogo = async () => {
  let oriImagePath = await SelectFile("选择主页图片", "");
  console.log(oriImagePath, "filePath -- console.log");
  let ext = await GetPathExt(oriImagePath);
  let allowExt = [".png", ".jpg", ".jpeg", ".bmp"];
  if (!allowExt.includes(ext)) {
    ToastError("请选则图片格式文件" + allowExt);
    return;
  }
  //当前文章的文件名
  let currFileName = await GetPathFileName(storeIndex.currArticlePath);
  currFileName = currFileName.replaceAll(".md", "");
  //组装新路径
  let publicDir = await PathJoin([storeVpConfig.fullSrcDir, "public"]);
  let newImagePath = await PathJoin([
    publicDir,
    "images",
    "home",
    currFileName + "_home" + ext,
  ]);
  let copyResult = await CopyPath(oriImagePath, newImagePath, false);
  ToastCheck(copyResult);
  storeIndex.currArticleFrontMatter["hero"]["image"]["src"] =
    newImagePath.replaceAll(publicDir, "");
};

const refPopShowFeatures = ref();
const setFeatures = () => {
  refPopShowFeatures.value.showModal();
};
</script>
<style scoped></style>
