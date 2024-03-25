<template>
  <div class="flex justify-center items-center">
    <q-toggle
      left-label
      label="是否使用多侧边栏:"
      v-model="useManeSidebars"
    ></q-toggle>
    <a-button
      class="bg-blue-500 hover:bg-blue-600 text-white flex justify-center items-center"
      dense
    >
      识别侧栏
    </a-button>
  </div>
  <hr class="my-2" />
  <!--  单侧边栏-->
  <div v-if="!useManeSidebars">
    <dy-add-sidebar
      show-add-top-nav
      :level="1"
      v-model:nav-array="sidebarList"
    ></dy-add-sidebar>
  </div>
  <!--  多侧边栏-->
  <div v-else>
    <dy-add-sidebar
      show-add-top-nav
      :level="1"
      v-model:nav-array="sidebarList"
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

const useManeSidebars = ref(false); //是否使用多侧边栏侧边栏

const storeConfig = useVpconfigStore();
const sidebarList = ref<VpNav[]>([]);

onMounted(() => {
  sidebarList.value = storeConfig.configData["themeConfig"]["nav"];
  console.log(sidebarList.value, "navList -- console.log");
});

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
const saveNav = () => {
  const formatData: VpNav[] = formatNavData(sidebarList.value);
  console.log(formatData, "formatData -- console.log");
  storeConfig.configData["themeConfig"]["nav"] = formatData;
  storeConfig.saveConfig();
};
</script>
<style scoped></style>
