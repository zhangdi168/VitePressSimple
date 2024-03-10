<script setup lang="ts">

import {IconPark} from "@icon-park/vue-next/es/all";
</script>

<template>
  <div class="px-2  py-1  border-t">
    <div class="flex items-center justify-end">
      <a-tooltip>
        <template #title>新建文件(根目录)</template>
        <icon-park
            :size="20"
            type="add-one"
            fill="#333"
            strokeLinejoin="bevel"
            theme="outline"
        />
      </a-tooltip>

      <a-tooltip>
        <template #title>新建文件夹(根目录)</template>
        <icon-park
            :size="20"
            type="folder-plus"
            fill="#333"
            strokeLinejoin="bevel"
            theme="outline"
        />
      </a-tooltip>
    </div>
  </div>
</template>

<style scoped>

</style>