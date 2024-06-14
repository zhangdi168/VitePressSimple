<template>
  <!--  使用界面模板1-->
  <nav-template1 v-if="useTemplateIndex === 1"></nav-template1>
  <!--  使用界面模板2-->
  <nav-template2 v-if="useTemplateIndex === 2"></nav-template2>
</template>
<script lang="ts" setup>
import NavTemplate1 from "@/layout/nav/navTemplate1.vue";
import NavTemplate2 from "@/layout/nav/navTemplate2.vue";
import { onBeforeUnmount, onMounted, ref } from "vue";
import { ConfigGet } from "../wailsjs/go/system/SystemService";
import { ConfigKeyLang, ConfigKeyProjectDir } from "@/constant/keys/config";
import { useIndexStore } from "@/store";
import {
  HasNewVersion,
  UpdateNewVersion,
} from "../wailsjs/go/services/UpdateService";
import { StartStaticServer } from "../wailsjs/go/services/StaticServer";
import { SystemMac } from "@/constant/enums/system";
import { useI18n } from "vue-i18n";
import { EventsOff, EventsOffAll, EventsOn } from "../wailsjs/runtime";
import { services } from "../wailsjs/go/models";
import NotifyShellData = services.NotifyShellData;
import { useShellStore } from "@/store/shell";
//在这里可以设置默认的模板
const useTemplateIndex = ref(2);
const storeIndex = useIndexStore();
const { locale } = useI18n();
// const vpConfig = useVpconfigStore();
onMounted(async () => {
  EventsOn("shell", (data: NotifyShellData) => {
    useShellStore().handlerShellNotify(data);
  });

  // await HistoryProject.initList(); //初始化历史数据
  // await vpConfig.initConfig();
  await storeIndex.getSystemType(); //获取当前系统
  await storeIndex.getVersion(); //获取当前系统版本
  await storeIndex.getStaticDir();
  await storeIndex.getStaticPort();
  ConfigGet(ConfigKeyLang).then((res) => {
    locale.value = res == "" ? "en" : res;
  }); //设置语言
  let hasNewVersion = await HasNewVersion();
  if (hasNewVersion) {
    UpdateNewVersion();
  }
  //设定初始项目
  let dir = await ConfigGet(ConfigKeyProjectDir);
  if (dir !== "") {
    await storeIndex.changeProject(dir);
  }
  //监听快捷键
  window.addEventListener("keydown", handleKeyDown);
});
onBeforeUnmount(() => {
  // 不要忘记在组件卸载时移除事件监听器，防止内存泄漏
  window.removeEventListener("keydown", handleKeyDown);

  //取消监听所有的事件
  EventsOff("shell");
});

function handleKeyDown(event: KeyboardEvent) {
  const isMacOS = storeIndex.systemType == SystemMac;

  // 根据操作系统判断使用哪个键
  const controlKey = isMacOS ? event.metaKey : event.ctrlKey;

  // 监听 Ctrl (或 Mac 的 Meta) + S 快捷键
  if (controlKey && event.key === "s") {
    console.log("快捷键被按下 -- console.log");
    storeIndex.saveCurrArticle();
  }
}
</script>
<style lang="scss"></style>
