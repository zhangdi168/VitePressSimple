<template>
  <!-- 头部-添加文章等 -->
  <div class="px-2 py-2 border-b">
    <div class="flex items-center justify-start" style="overflow-y: hidden">
      <div class="mx-1 cursor-pointer">
        <q-tooltip>新建文件(根目录)</q-tooltip>
        <icon-park
          :size="iconSize"
          fill="#333"
          strokeLinejoin="bevel"
          theme="outline"
          type="add"
        />
      </div>

      <div class="mx-1 cursor-pointer">
        <q-tooltip>刷新目录</q-tooltip>
        <icon-park
          :size="iconSize"
          fill="#333"
          strokeLinejoin="bevel"
          theme="outline"
          type="refresh"
          @click="storeIndex.loadTreeData()"
        />
      </div>

      <div class="mx-1 cursor-pointer">
        <q-tooltip>刷新目录</q-tooltip>
        <icon-park
          :size="iconSize"
          fill="#5f4364"
          strokeLinejoin="bevel"
          theme="outline"
          type="search"
          @click="storeIndex.loadTreeData()"
        />
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { IconPark } from "@icon-park/vue-next/es/all";

import { useIndexStore } from "@/store";
import { nextTick, onMounted, ref } from "vue";
import { FileSearchOutlined } from "@ant-design/icons-vue";

const iconSize = ref(24);
const storeIndex = ref();
onMounted(() => {
  nextTick(() => {
    storeIndex.value = useIndexStore();
  });
});
</script>
<style scoped></style>
