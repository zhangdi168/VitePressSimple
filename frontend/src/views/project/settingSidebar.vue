<template>
  <div class="flex justify-center items-center">
    <div class="p-1 mr-2 border">
      <select-setting-lang></select-setting-lang>
    </div>

    <!--    多侧边栏选择-->
    <div
      class="flex justify-center border items-center mr-4"
    >
      <div class=" ml-6">
        <q-toggle
          left-label
          @update:model-value="toggleUseManySidebars"
          label="是否使用多侧边栏:"
          v-model="isUseManySidebars"
        ></q-toggle>
      </div>
      <div v-if="isUseManySidebars">
        选择当前操作的侧栏：
        <a-radio-group
          @change="subSidebarChange"
          v-model:value="currSelectSubDir"
        >
          <a-radio
            v-for="(item, index) in subSidebarList"
            :key="index"
            :value="item"
          >{{ item }}
          </a-radio>
        </a-radio-group>
      </div>
    </div>
    <a-popconfirm

      title="识别到的侧栏数据将会全部覆盖当前侧栏数据"
      ok-text="我已知晓"
      cancel-text="算了"
      :ok-button-props="{ class: 'bg-blue-6 00 text-white' }"
      @confirm="recognitionSidebar()"
    >
      <a-button
        class="bg-blue-500 hover:bg-blue-600 text-white flex justify-center items-center"
        dense
      >
        自动识别侧栏
      </a-button>
    </a-popconfirm>
  </div>
  <hr class="my-2" />
  <!--  单侧边栏-->
  <div>

    <dy-add-sidebar
      show-add-top-nav
      :level="1"
      v-model:sidebar-array="sidebarList"
    ></dy-add-sidebar>
  </div>


  <div class="mt-4 flex justify-center">
    <a-button
      @click="saveNav()"
      class="bg-blue-600 hover:bg-blue-500 text-white flex justify-center items-center"
    >
      <icon-park
        class="mr-1"
        strokeLinejoin="bevel"
        theme="outline"
        type="save"
      />
      保存当前导航
    </a-button>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { useVpconfigStore } from "@/store/vpconfig";
import DyAddNav from "@/components/dyAddNav.vue";
import { VpNav } from "@/utils/tree";
import { IconPark } from "@icon-park/vue-next/es/all";
import DyAddSidebar from "@/components/dyAddSidebar.vue";
import { ParseTreeData } from "../../../wailsjs/go/services/ArticleTreeData";
import { StringGlobalLang } from "@/configs/cnts";
import { GetPathDir, GetPathFileName, PathJoin } from "../../../wailsjs/go/system/SystemService";
import SelectSettingLang from "@/components/selectSettingLang.vue";
import { ToastError } from "@/utils/Toast";
import { dto } from "../../../wailsjs/go/models";
import TreeNode = dto.TreeNode;
import { useIndexStore } from "@/store";

const storeConfig = useVpconfigStore();
const storeIndex = useIndexStore();
const sidebarList = ref<VpNav[]>([]);
const subSidebarList = ref<string[]>();//多侧栏子目录
const currSelectSubDir = ref("");
const isUseManySidebars = ref(false);
onMounted(() => {
  isUseManySidebars.value = !Array.isArray(storeConfig.currLangConfig["themeConfig"]["sidebar"]);
  if (isUseManySidebars) {
    getSubSidebarDirList();
    let keys = Object.keys(storeConfig.currLangConfig["themeConfig"]["sidebar"]);
    if (keys.length > 0) {
      sidebarList.value = storeConfig.currLangConfig["themeConfig"]["sidebar"][keys[0]];
    }
  } else {
    sidebarList.value = storeConfig.currLangConfig["themeConfig"]["sidebar"];
  }
});
//识别侧栏
const recognitionSidebar = async () => {
  let zhLangDir = await PathJoin([
    storeConfig.srcDir,
    storeConfig.currSettingLang
  ]);
  let baseDir =
    storeConfig.currSettingLang == StringGlobalLang
      ? storeConfig.srcDir
      : zhLangDir;
  let treeData = await ParseTreeData(baseDir); //[{key:"路径",title:""}]
  console.log(treeData, "treeData -- console.log");
  if (!treeData) {
    ToastError(`识别到的侧栏数据为空，请确保${baseDir}文件夹下存在文件`);
  } else {
    //开始转换数据,将属性转换成侧栏结构
    let existingLinks: string[] = sidebarList.value.flatMap(nav => nav.items ? [nav.link, ...nav.items.map(item => item.link)] : [nav.link]);
    existingLinks = existingLinks.filter(link => link !== undefined && link.endsWith(".md"));

    let list = convertTreeNodesToVpNavList(treeData, existingLinks);
    sidebarList.value = list;
    console.log(list, "list -- console.log");
  }
};

//侧栏改变
const subSidebarChange = () => {
  if (currSelectSubDir.value !== "") {
    sidebarList.value = storeConfig.currLangConfig["themeConfig"]["sidebar"][`/${currSelectSubDir.value}/`];
  }
  //console.log(currSelectSubDir.value, "currSelectSubDir.value -- console.log");
};

//切换是否使用多侧栏
const toggleUseManySidebars = (flag: any, evet: any) => {
  if (flag) {//获取子目录
    getSubSidebarDirList();
  }
};
//获取多侧栏的子目录
const getSubSidebarDirList = async () => {
  let list = await ParseTreeData(storeConfig.SrcLangDir);
  subSidebarList.value = list.map(item => parseText(item.key));
  currSelectSubDir.value = subSidebarList.value[0] ?? "";
};
const convertTreeNodesToVpNavList = (treeNodes: TreeNode[], existingLinks: string[]): VpNav[] => {
  let vpNavList: VpNav[] = [];
  treeNodes.forEach((treeNode) => {
    let link = parseLink(treeNode.key);
    let text = parseText(treeNode.title);
    const vpNavItem: VpNav = {
      link: link,
      text: text
    };

    if (!existingLinks.includes(vpNavItem.link)) {
      existingLinks.push(vpNavItem.link);
      vpNavList.push(vpNavItem);
      if (treeNode.children) {
        vpNavItem.items = convertTreeNodesToVpNavList(treeNode.children, existingLinks).filter(Boolean);
      }
    }
  });
  return vpNavList;
};

//解析link
const parseLink = (pathKey: string): string => {
  let path = pathKey.replaceAll(storeConfig.fullSrcDir, "").replaceAll("\\", "/");
  return `${path}`;
};
//解析Text
const parseText = (pathKey: string): string => {
  let items = pathKey.replaceAll("\\", "/").split("/");
  let text = items[items.length - 1] ?? "";
  return text;
};

const formatNavData = (data: VpNav[]) => {
  return data.map((item) => {
    const formattedItem: any = { text: item.text };
    if (item.items && item.items.length > 0) {
      // If 'items' exists and has a length greater than 0
      formattedItem.items = item.items.map((subItem) => ({
        text: subItem.text,
        link: subItem.items?.length ? undefined : subItem.link,
        // Recursive call to handle nested items
        items: subItem.items?.length ? formatNavData(subItem.items) : undefined
      }));
    } else {
      // If 'items' doesn't exist or its length is 0
      formattedItem.link = item.link;
      formattedItem.text = item.text;
    }
    return formattedItem;
  });
};
const saveNav = () => {
  const formatData: VpNav[] = formatNavData(sidebarList.value);
  console.log(formatData, "formatData -- console.log");
  storeConfig.currLangConfig["themeConfig"]["sidebar"] = formatData;
  storeConfig.saveConfig();
};
</script>
<style scoped></style>
