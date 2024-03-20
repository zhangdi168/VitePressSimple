<template>
  <div>
    <dy-add-nav
      show-add-top-nav
      :level="1"
      v-model:nav-array="navList"
    ></dy-add-nav>
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

const storeConfig = useVpconfigStore();
const navList = ref<VpNav[]>([]);

onMounted(() => {
  navList.value = storeConfig.configData["themeConfig"]["nav"];
  console.log(navList.value, "navList -- console.log");
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
  const formatData: VpNav[] = formatNavData(navList.value);
  console.log(formatData, "formatData -- console.log");
  storeConfig.configData["themeConfig"]["nav"] = formatData;
  storeConfig.saveConfig();
};
</script>
<style scoped></style>
