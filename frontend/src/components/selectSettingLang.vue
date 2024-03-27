<template>
  <div
    class="flex justify-start items-center ml-2"
    v-if="storeConfig.IsUseI18n"
  >
    选择当前操作的语言：
    <a-radio-group
      @change="langChange"
      v-model:value="storeConfig.currSettingLang"
    >
      <a-radio-button :value="StringGlobalLang"
      >{{ StringGlobalLang }}
      </a-radio-button>
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
import { StringGlobalLang } from "@/configs/cnts";

const storeConfig = useVpconfigStore();
const langChange = () => {
  storeConfig.changeCurrLangConfig(storeConfig.currSettingLang);
};
</script>

<style scoped></style>
