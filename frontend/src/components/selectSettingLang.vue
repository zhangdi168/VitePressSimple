<template>
  <div
    class="flex justify-start items-center ml-2"
    v-if="storeConfig.IsUseI18n"
  >
    <icon-park strokeLinejoin="bevel" theme="outline" type="translate" />
    当前操作的语言：
    <a-radio-group
      @change="langChange"
      v-model:value="storeConfig.currSettingLang"
    >
      <a-radio-button v-show="!storeConfig.IsUseI18n" :value="StringGlobalLang">
        <q-tooltip
          >配置将写入根themeConfig，该选项下仅对单语言模式生效(即语言列表为空时该选项配置生效)
        </q-tooltip>
        {{ StringGlobalLang }}
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
import { IconPark } from "@icon-park/vue-next/es/all";
import { ToastInfo } from "@/utils/Toast";

const storeConfig = useVpconfigStore();
const emits = defineEmits(["changeLang"]);
const langChange = () => {
  storeConfig.changeCurrLangConfig(storeConfig.currSettingLang);
  if (storeConfig.currSettingLang == StringGlobalLang) {
    ToastInfo("已切换到global配置，global配置仅对单语言模式生效");
  }
  emits("changeLang");
};
</script>

<style scoped></style>
