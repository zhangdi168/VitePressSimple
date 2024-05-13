<template>
  <div>
    <sim-radio
      :tooltip="lang('pageProject.settingSearch.tooltips.tooltips')"
      v-model="storeConfig.configData['themeConfig']['search']['provider']"
      :label="lang('pageProject.settingSearch.searchProvider')"
      :items="[
        { label: 'local', value: 'local' },
        { label: 'algolia', value: 'algolia' },
      ]"
    ></sim-radio>

    <div
      class="flex justify-start"
      v-show="
        storeConfig.configData['themeConfig']['search']['provider'] == 'algolia'
      "
    >
      <sim-input
        v-model="
          storeConfig.configData['themeConfig']['search']['options']['appId']
        "
        :tooltip="lang('pageProject.settingSearch.tooltips.AlgoliaAppId')"
        label="AppId"
      >
      </sim-input>
      <sim-input
        v-model="
          storeConfig.configData['themeConfig']['search']['options']['apiKey']
        "
        :tooltip="lang('pageProject.settingSearch.tooltips.AlgoliaSearchKey')"
        label="apiKey"
      >
      </sim-input>
      <sim-input
        v-model="
          storeConfig.configData['themeConfig']['search']['options'][
            'indexName'
          ]
        "
        :tooltip="lang('pageProject.settingSearch.tooltips.AlgoliaIndexName')"
        label="indexName"
      >
      </sim-input>
    </div>

    <div class="flex justify-center">
      <a-button
        @click="saveBaseConfig"
        class="bg-blue-600 hover:bg-blue-500 text-white flex justify-center items-center"
      >
        <icon-park
          class="mr-1"
          strokeLinejoin="bevel"
          theme="outline"
          type="save"
        />
        {{ lang("common.saveConfig") }}
      </a-button>
    </div>
  </div>
</template>
<script setup lang="ts">
import SimRadio from "@/components/simRadio.vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { onBeforeMount } from "vue";
import { IsEmptyValue } from "@/utils/utils";
import { defaultShareConfigValue } from "@/configs/defaultShareConfig";
import SimInput from "@/components/simInput.vue";
import { lang } from "@/utils/language";

const storeConfig = useVpconfigStore();
onBeforeMount(() => {
  //监测语言
  if (IsEmptyValue(storeConfig.configData["themeConfig"]["search"])) {
    storeConfig.configData["themeConfig"]["search"] =
      defaultShareConfigValue.themeConfig.search;
  }

  checkSearchKey1("provider", "local");
  checkSearchKey1(
    "options",
    defaultShareConfigValue.themeConfig.search.options,
  );
  checkSearchKey2("options", "appId", "");
  checkSearchKey2("options", "apiKey", "");
  checkSearchKey2("options", "indexName", "");
});

const checkSearchKey1 = (key1: string, defaultValue: any) => {
  if (IsEmptyValue(storeConfig.configData["themeConfig"]["search"][key1])) {
    storeConfig.configData["themeConfig"]["search"][key1] = defaultValue;
  }
};
const checkSearchKey2 = (key1: string, key2: string, defaultValue: any) => {
  checkSearchKey1(key1, {});
  if (
    IsEmptyValue(storeConfig.configData["themeConfig"]["search"][key1][key2])
  ) {
    let base = (storeConfig.configData["themeConfig"]["search"][key1][key2] =
      defaultValue);
  }
};
const saveBaseConfig = () => {
  storeConfig.saveConfig();
};
</script>
<style scoped></style>
