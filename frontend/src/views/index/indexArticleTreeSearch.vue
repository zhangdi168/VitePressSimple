<template>
  <div>
    <a-input-search
      v-model:value="storeIndex.searchValue"
      style="margin-bottom: 8px"
      placeholder="Search"
    />
  </div>
</template>
<script setup lang="ts">
import { useIndexStore } from "@/store";
import { watch } from "vue";
import { TreeProps } from "ant-design-vue";

const storeIndex = useIndexStore();
</script>
<style scoped></style>
