<script setup lang="ts">
import { IconPark } from "@icon-park/vue-next/es/all";
import { lang } from "../../utils/language";
</script>

<template>
  <div class="px-2 py-1 border-t">
    <div class="flex items-center justify-end">
      <a-tooltip>
        <template #title>{{ lang("pageIndex.newFileInRoot") }}</template>
        <icon-park
          :size="20"
          type="add-one"
          fill="#333"
          strokeLinejoin="bevel"
          theme="outline"
        />
      </a-tooltip>

      <a-tooltip>
        <template #title>{{ lang("pageIndex.newFolderInRoot") }}</template>
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

<style scoped></style>
