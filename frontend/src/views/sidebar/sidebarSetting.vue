<template>
  <empty-project></empty-project>
  <div v-if="!storeIndex.IsEmptyProject" class="mx-6 mt-4 text-xl">
    <!--        语言选择-->
    <div class="flex justify-start items-center min-h-11 pb-1 ml-2 pr-2 border">
      <select-setting-lang
        @change-lang="langChange"
        class="mt-1"
      ></select-setting-lang>
      <a-popconfirm
        class="ml-3 pt-1"
        title="切换后将[清空]当前语言的侧栏数据"
        ok-text="我已知晓"
        cancel-text="算了"
        :ok-button-props="{ class: 'bg-blue-6 00 text-white' }"
        @confirm="changeSidebarConfirmModal()"
      >
        <q-tooltip class="text-2xl">
          <span class="text-cyan-500"> {{ storeConfig.currSettingLang }}</span
          >当前为
          <span class="text-cyan-500"
            >{{ isUseManySidebars ? "多" : "单" }}侧栏</span
          >
          模式，
          <span class="text-red-500"
            >点击切换成{{ isUseManySidebars ? "单" : "多" }}侧栏</span
          >
        </q-tooltip>
        <a-button class="cursor-pointer bg-green-100">
          <icon-park
            strokeLinejoin="bevel"
            theme="outline"
            class="pr-1"
            type="switch-themes"
          />
        </a-button>
      </a-popconfirm>
    </div>
    <!--    多侧边栏选择-->
    <div
      class="flex pl-3 justify-between ml-2 items-center min-h-11 pb-2 border"
      v-if="isUseManySidebars"
    >
      <div>
        选择当前编辑的侧栏：
        <a-radio-group
          class="mt-2"
          @change="subSidebarChange"
          v-model:value="currSelectSidebarKey"
        >
          <a-radio-button
            v-for="(item, index) in currSidebarSubDirList"
            :key="index"
            :value="item"
          >
            <q-tooltip
              >当路由位于 ‘{{ item }}’ 目录时，会显示此侧边栏</q-tooltip
            >
            {{ item }}
          </a-radio-button>
        </a-radio-group>
      </div>
      <div class="text-red-500" v-if="!hasSubSidebarDir">
        {{ emptySubDirText }}
      </div>
      <div class="mr-4 cursor-pointer" @click="setCurrTreeData">
        <q-tooltip>重新加载侧栏目录</q-tooltip>
        <icon-park strokeLinejoin="bevel" theme="outline" type="refresh" />
      </div>
    </div>
    <!--        顶部操作按钮栏目-->
    <div class="mt-4 flex justify-start">
      <!--      添加顶级侧栏-->
      <a-button
        class="bg-blue-200 mx-2 flex justify-center items-center hover:bg-blue-100"
        @click="addTopSidebar()"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="add-one"
        />
        新增顶级侧栏
      </a-button>
      <!--    自动识别侧栏-->
      <div class="mx-1">
        <a-popconfirm
          title="识别到的侧栏数据将会全部覆盖当前侧栏数据"
          ok-text="我已知晓"
          cancel-text="算了"
          :ok-button-props="{ class: 'bg-blue-500 text-white' }"
          @confirm="recognitionSidebar()"
        >
          <q-tooltip
            >自动识别{{ storeConfig.SrcLangDir }}目录下的侧栏数据
          </q-tooltip>
          <a-button
            class="bg-blue-500 hover:bg-blue-600 text-white flex justify-center items-center"
            dense
          >
            <icon-park
              class="mr-1"
              strokeLinejoin="bevel"
              theme="outline"
              type="scanning"
            />
            自动识别当前侧栏
          </a-button>
        </a-popconfirm>
      </div>
      <!--      保存到配置文件-->
      <a-button
        @click="saveSidebar()"
        class="bg-blue-600 hover:bg-blue-500 text-white flex justify-center items-center"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="save"
        />
        保存当前侧栏
      </a-button>
    </div>
    <!--        动态添加侧栏-->
    <dy-add-sidebar
      class="ml-4"
      ref="refSidebar"
      :level="1"
      v-model:sidebar-array="sidebarTree"
    ></dy-add-sidebar>
  </div>
</template>
<script setup lang="ts">
import { computed, createVNode, nextTick, onMounted, ref } from "vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { VpNav } from "@/utils/tree";
import { IconPark } from "@icon-park/vue-next/es/all";
import DyAddSidebar from "@/components/dyAddSidebar.vue";
import { ParseTreeData } from "../../../wailsjs/go/services/ArticleTreeData";
import { StringGlobalLang } from "@/configs/cnts";
import { PathJoin } from "../../../wailsjs/go/system/SystemService";
import SelectSettingLang from "@/components/selectSettingLang.vue";
import { ToastError, ToastSuccess } from "@/utils/Toast";
import { dto } from "../../../wailsjs/go/models";
import { useIndexStore } from "@/store";
import DyAddKV from "@/components/dyAddKV.vue";
import TreeNode = dto.TreeNode;
import DyAddNav from "@/components/dyAddNav.vue";
import { IsEmptyValue } from "@/utils/utils";
import EmptyProject from "@/components/emptyProject.vue";

const splitterModel = ref(36);
const storeConfig = useVpconfigStore();
const storeIndex = useIndexStore();
const sidebarTree = ref<VpNav[]>([]); //当前指向的侧栏
const currSidebarSubDirList = ref<string[]>([]); //识别到的多侧栏子目录
// const currSidebarList = ref<string[]>([]);//当前设置的多侧栏子目录
const currSelectSidebarKey = ref("");
const isUseManySidebars = ref(false);
onMounted(() => {
  nextTick(() => {
    checkCurrIsUseManySidebars();
    setCurrTreeData();
  });
});

const checkCurrIsUseManySidebars = () => {
  isUseManySidebars.value = !Array.isArray(
    storeConfig.currLangConfig["themeConfig"]["sidebar"],
  );
};
const hasSubSidebarDir = computed(() => {
  return currSidebarSubDirList.value && currSidebarSubDirList.value.length > 0;
});
const emptySubDirText = computed(() => {
  if (isUseManySidebars.value && !hasSubSidebarDir.value) {
    return storeConfig.SrcLangDir + " 路径下不存在子目录，请先前往创建！";
  }
  return "";
});

//设置当前侧栏指向的数据
const setCurrTreeData = async () => {
  checkCurrIsUseManySidebars(); //根据配置文件检测当前是否为多侧栏模式
  if (isUseManySidebars.value) {
    //多侧栏模式
    await getSubSidebarDirList(); //获取本地侧栏目录列表
    sidebarTree.value =
      storeConfig.currLangConfig["themeConfig"]["sidebar"][
        currSelectSidebarKey.value
      ];
  } else {
    //单侧栏模式
    sidebarTree.value = storeConfig.currLangConfig["themeConfig"]["sidebar"];
  }

  if (IsEmptyValue(sidebarTree.value)) {
    sidebarTree.value = [];
  }
};

//识别侧栏
const recognitionSidebar = async () => {
  if (emptySubDirText.value !== "") {
    ToastError(emptySubDirText.value);
    return;
  }

  let zhLangDir = await PathJoin([
    storeConfig.srcDir,
    storeConfig.currSettingLang,
  ]);
  let baseDir = storeConfig.srcDir;
  if (isUseManySidebars.value) {
    if (currSelectSidebarKey.value == "") {
      ToastError("请选择当前编辑的侧栏");
      return;
    }
    //多侧栏模式，再加上选中的侧栏
    baseDir = baseDir + currSelectSidebarKey.value;
  }
  let treeData = await ParseTreeData(baseDir); //[{key:"路径",title:""}]
  console.log(treeData, "treeData -- console.log");
  if (!treeData) {
    ToastError(`识别到的侧栏数据为空，请确保${baseDir}文件夹下存在文件`);
  } else {
    //开始转换数据,将属性转换成侧栏结构
    let existingLinks: string[] = sidebarTree.value.flatMap((nav) =>
      nav.items
        ? [nav.link, ...nav.items.map((item) => item.link)]
        : [nav.link],
    );
    existingLinks = existingLinks.filter(
      (link) => link !== undefined && link.endsWith(".md"),
    );

    let list = convertTreeNodesToVpNavList(treeData, existingLinks);
    sidebarTree.value = list;
    console.log(list, "list -- console.log");
  }
};

//侧栏改变
const subSidebarChange = () => {
  setCurrTreeData();
};

//语言改变
const langChange = () => {
  currSelectSidebarKey.value = "";
  // console.log(storeConfig.SrcLangDir, "storeConfig.SrcLangDir -- console.log");
  setCurrTreeData();
};

//切换侧栏确认弹窗
const changeSidebarConfirmModal = () => {
  if (isUseManySidebars.value) {
    console.log("当前为多侧栏，即将切换成单侧栏 -- console.log");
    storeConfig.currLangConfig["themeConfig"]["sidebar"] = [];
    isUseManySidebars.value = false;
  } else {
    console.log("当前为单侧栏，即将切换成多侧栏 -- console.log");
    storeConfig.currLangConfig["themeConfig"]["sidebar"] = {};
    isUseManySidebars.value = true;
  }
  setCurrTreeData();
};

//获取多侧栏的子目录
const getSubSidebarDirList = async () => {
  if (!isUseManySidebars.value) {
    return; //非多侧栏模式下，此函数用户不到
  }
  let list = await ParseTreeData(storeConfig.SrcLangDir);
  if (!list) {
    // ToastError(
    //   `识别到的子侧栏数据为空，请确保${storeConfig.SrcLangDir}下存在子文件夹`,
    // );
    currSidebarSubDirList.value = [];
    return;
  }

  currSidebarSubDirList.value = list
    .filter((item) => item && !item.title.endsWith(".md"))
    .map((item) => {
      if (storeConfig.IsUseI18n) {
        return `/${storeConfig.currSettingLang}/${item.title}/`;
      } else {
        return `/${item.title}/`;
      }
    });

  if (currSelectSidebarKey.value == "") {
    currSelectSidebarKey.value = currSidebarSubDirList.value[0];
  }
};
const convertTreeNodesToVpNavList = (
  treeNodes: TreeNode[],
  existingLinks: string[],
): VpNav[] => {
  let vpNavList: VpNav[] = [];
  treeNodes.forEach((treeNode) => {
    let link = parseLink(treeNode.key);
    let text = parseText(treeNode.title);
    const vpNavItem: VpNav = {
      link: link,
      text: text,
    };

    if (!existingLinks.includes(vpNavItem.link)) {
      existingLinks.push(vpNavItem.link);
      vpNavList.push(vpNavItem);
      if (treeNode.children) {
        vpNavItem.items = convertTreeNodesToVpNavList(
          treeNode.children,
          existingLinks,
        ).filter(Boolean);
      }
    }
  });
  return vpNavList;
};

//解析link
const parseLink = (pathKey: string): string => {
  let path = pathKey
    .replaceAll(storeConfig.fullSrcDir, "")
    .replaceAll("\\", "/");
  return `${path}`.replaceAll(".md", "");
};
//解析Text
const parseText = (pathKey: string): string => {
  let items = pathKey.replaceAll("\\", "/").split("/");
  let text = items[items.length - 1] ?? "";
  return text.replaceAll(".md", "");
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
        items: subItem.items?.length ? formatNavData(subItem.items) : undefined,
      }));
    } else {
      // If 'items' doesn't exist or its length is 0
      formattedItem.link = item.link;
      formattedItem.text = item.text;
    }
    return formattedItem;
  });
};
const saveSidebar = () => {
  if (emptySubDirText.value !== "") {
    ToastError(emptySubDirText.value);
    return;
  }
  const formatData: VpNav[] = formatNavData(sidebarTree.value);
  checkCurrIsUseManySidebars();
  if (isUseManySidebars.value) {
    if (currSelectSidebarKey.value == "") {
      ToastError("请选择操作的侧栏");
      return;
    }
    storeConfig.currLangConfig["themeConfig"]["sidebar"][
      currSelectSidebarKey.value
    ] = formatData;
  } else {
    storeConfig.currLangConfig["themeConfig"]["sidebar"] = formatData;
  }

  storeConfig.saveConfig();
};

const refSidebar = ref();

const addTopSidebar = () => {
  if (emptySubDirText.value !== "") {
    ToastError(emptySubDirText.value);
    return;
  }
  refSidebar.value.addTopNav();
};
</script>
<style scoped></style>
