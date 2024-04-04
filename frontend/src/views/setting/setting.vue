<template>
  <div class="mt-4 ml-2">
    <a-tabs
      style="min-height: 55vh"
      v-model:activeKey="activeKey"
      :tab-position="'left'"
      type="line"
      animated
    >
      <a-tab-pane key="1" tab="基础设置">
        <setting-base></setting-base>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import SettingBase from "@/views/setting/settingBase.vue";

const activeKey = ref("1");
</script>

<style scoped></style>
