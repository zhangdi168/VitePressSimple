<template>
  <div class="flex flex-wrap mx-8 mt-3">
    {{ lang("pageProject.settingSocial.tooltipsIcon") }}
    <div @click="CopyText(item)" class="cursor-pointer" v-for="item in soList">
      <a-tag color="cyan">{{ item }}</a-tag>
    </div>
  </div>
  <hr class="my-3" />
  <dy-add-k-v
    class="mt-2 mx-6"
    :add-btn-text="lang('pageProject.settingSocial.addSocialAccount')"
    :key-placeholder="lang('pageProject.settingSocial.icon')"
    value-placeholder="lang('pageProject.settingSocial.link')"
    key-name="icon"
    value-name="link"
    ref="refDyAdd"
    v-model:objs="storeConfig.configData['themeConfig']['socialLinks']"
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
      {{ lang("common.saveConfig") }}
    </a-button>
  </div>
</template>
<script setup lang="ts">
import DyAddKV from "@/components/dyAddKV.vue";
import { ToastError, ToastSuccess } from "@/utils/Toast";
import { useVpconfigStore } from "@/store/vpconfig";
import { IconPark } from "@icon-park/vue-next/es/all";
import { onMounted, ref } from "vue";
import { lang } from "../../utils/language";

onMounted(() => {
  console.log(
    storeConfig.configData["themeConfig"]["socialLinks"],
    " -- console.log",
  );
});
const storeConfig = useVpconfigStore();
const soList = [
  "github",
  "discord",
  "facebook",
  "youtube",
  "instagram",
  "linkedin",
  "mastodon",
  "npm",
  "slack",
  "twitter",
  "x",
];
const CopyText = async (text: string) => {
  try {
    await navigator.clipboard.writeText(text);
    ToastSuccess("复制成功");
  } catch (err: any) {
    ToastError("复制失败" + err.message);
  }
};
const refDyAdd = ref();
const saveSocialConfig = () => {
  storeConfig.saveConfig();
};
</script>

<style scoped></style>
