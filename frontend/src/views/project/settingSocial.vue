<template>
  <div class="flex flex-wrap mx-8 mt-3">
    当前支持的图标标识(点击可复制)：
    <div @click="CopyText(item)" class="cursor-pointer" v-for="item in soList">
      <a-tag color="cyan">{{ item }}</a-tag>
    </div>
  </div>
  <hr class="my-3" />
  <dy-add-k-v
    class="mt-2 mx-6"
    add-btn-text="添加社交账户"
    key-placeholder="图标标识"
    value-placeholder="跳转链接"
    v-model:obj="storeConfig.configData['socialLinks']"
  ></dy-add-k-v>

  <hr class="my-2" />
  <div class="flex justify-center">
    <a-button
      @click="saveSocialConfig"
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
import DyAddKV from "@/components/dyAddKV.vue";
import { ToastError, ToastSuccess } from "@/utils/Toast";
import { useVpconfigStore } from "@/store/vpconfig";
import { IconPark } from "@icon-park/vue-next/es/all";

const storeConfig = useVpconfigStore();
const soList = [
  "discord",
  "facebook",
  "github",
  "instagram",
  "linkedin",
  "mastodon",
  "npm",
  "slack",
  "twitter",
  "x",
  "youtube",
];
const CopyText = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
    ToastSuccess("复制成功");
  } catch (err: any) {
    ToastError("复制失败" + err.message);
  }
};
const saveSocialConfig = () => {
  storeConfig.saveConfig();
};
</script>

<style scoped></style>
