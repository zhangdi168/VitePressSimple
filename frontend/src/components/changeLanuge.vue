<template>
  <div
    class="flex justify-start items-center ml-2"
    v-if="storeConfig.IsI18nRouting"
  >
    选择当前操作的语言：
    <a-radio-group v-model:value="storeConfig.currSettingLang">
      <a-radio-button
        v-for="(item, index) in storeConfig.GetLangArray"
        :key="index"
        :value="item"
        >{{ item }}
      </a-radio-button>
    </a-radio-group>
  </div>
</template>
<script setup lang="ts">
import { useVpconfigStore } from "@/store/vpconfig";

const storeConfig = useVpconfigStore();
</script>

<style scoped></style>
