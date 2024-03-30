<template>
  <!--  使用界面模板1-->
  <nav-template1 v-if="useTemplateIndex === 1"></nav-template1>
  <!--  使用界面模板2-->
  <nav-template2 v-if="useTemplateIndex === 2"></nav-template2>
</template>
<script lang="ts" setup>
import NavTemplate1 from "@/layout/nav/navTemplate1.vue";
import NavTemplate2 from "@/layout/nav/navTemplate2.vue";
import { onMounted, ref } from "vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { ConfigGet } from "../wailsjs/go/system/SystemService";
import { ConfigKeyProjectDir } from "@/constant/keys/config";
import { useIndexStore } from "@/store";
import { HistoryProject } from "@/utils/historyProject";
import { useHistoryStore } from "@/store/history";
//在这里可以设置默认的模板
const useTemplateIndex = ref(2);
const storeIndex = useIndexStore();
const vpConfig = useVpconfigStore();
onMounted(async () => {
  // await HistoryProject.initList(); //初始化历史数据
  // await vpConfig.initConfig();
  await useHistoryStore().initList();
});
</script>
<style lang="scss"></style>
