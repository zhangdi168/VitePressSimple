<template>
  <div class="m-3">
    <a-alert
      type="success"
      closable
      :show-icon="false"
      :description="
        lang('pageProject.settingLang.tooltips.multiLanguageEnabled')
      "
    ></a-alert>
  </div>
  <dy-add-k-v
    :add-btn-text="lang('pageProject.settingLang.labels.addLangType')"
    :key-placeholder="
      lang('pageProject.settingLang.placeholders.langDirectory')
    "
    :value-placeholder="lang('pageProject.settingLang.placeholders.langLabel')"
    key-name="lang"
    value-name="label"
    ref="refAddLang"
    :remove-confirm="true"
    :remove-confirm-text="lang('pageProject.settingLang.labels.removeConfirm')"
    @remove-item="removeLangItem"
    v-model:objs="inputLangArray"
  >
  </dy-add-k-v>

  <hr class="my-2" />
  <select-setting-lang></select-setting-lang>
  <div class="flex justify-start items-center my-2 mx-8">
    <sim-switch
      v-model="storeConfig.currLangConfig['themeConfig']['i18nRouting']"
      :tooltip="lang('pageProject.settingLang.tooltips.i18nRouting')"
      :label="lang('pageProject.settingLang.labels.i18nRouting')"
    ></sim-switch>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['siteTitle']"
      :tooltip="lang('pageProject.settingLang.tooltips.siteTitleInfo')"
      :placeholder="lang('pageProject.settingLang.placeholders.siteTitle')"
      :label="lang('pageProject.settingLang.labels.siteTitle')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['footer']['message']"
      :tooltip="lang('pageProject.settingLang.tooltips.footerMessageInfo')"
      :placeholder="lang('pageProject.settingLang.placeholders.footerMessage')"
      :label="lang('pageProject.settingLang.labels.footerMessage')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['footer']['copyright']"
      :tooltip="lang('pageProject.settingLang.tooltips.copyrightInfo')"
      :placeholder="lang('pageProject.settingLang.placeholders.copyright')"
      :label="lang('pageProject.settingLang.labels.copyright')"
    ></sim-input>

    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['langMenuLabel']"
      :tooltip="lang('pageProject.settingLang.tooltips.langMenuLabelInfo')"
      :placeholder="
        lang('pageProject.settingLang.placeholders.langSwitchLabel')
      "
      :label="lang('pageProject.settingLang.labels.langSwitchLabel')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['returnToTopLabel']"
      :tooltip="lang('pageProject.settingLang.tooltips.returnToTop')"
      placeholder=""
      :label="lang('pageProject.settingLang.labels.returnToTop')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['sidebarMenuLabel']"
      :tooltip="lang('pageProject.settingLang.tooltips.sidebarMenuLabelInfo')"
      :placeholder="lang('pageProject.settingLang.placeholders.sidebarMenu')"
      :label="lang('pageProject.settingLang.labels.sidebarMenu')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['darkModeSwitchLabel']"
      :tooltip="
        lang('pageProject.settingLang.tooltips.darkModeSwitchLabelInfo')
      "
      :placeholder="lang('pageProject.settingLang.placeholders.darkModeSwitch')"
      :label="lang('pageProject.settingLang.labels.darkModeSwitch')"
    ></sim-input>
    <sim-input
      v-model="
        storeConfig.currLangConfig['themeConfig']['lightModeSwitchTitle']
      "
      :tooltip="
        lang('pageProject.settingLang.tooltips.lightModeSwitchTitleInfo')
      "
      :placeholder="
        lang('pageProject.settingLang.placeholders.lightModeSwitchTitle')
      "
      :label="lang('pageProject.settingLang.labels.lightModeSwitchTitle')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['darkModeSwitchTitle']"
      :tooltip="
        lang('pageProject.settingLang.tooltips.darkModeSwitchTitleInfo')
      "
      :placeholder="
        lang('pageProject.settingLang.placeholders.darkModeSwitchTitle')
      "
      :label="lang('pageProject.settingLang.labels.darkModeSwitchTitle')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['editLink']['text']"
      :tooltip="lang('pageProject.settingLang.tooltips.editLinkTextInfo')"
      :placeholder="lang('pageProject.settingLang.placeholders.editLinkText')"
      :label="lang('pageProject.settingLang.labels.editLinkText')"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['editLink']['pattern']"
      :tooltip="lang('pageProject.settingLang.tooltips.editLinkPatternInfo')"
      :placeholder="
        lang('pageProject.settingLang.placeholders.editLinkPattern')
      "
      :label="lang('pageProject.settingLang.labels.editLinkPattern')"
    ></sim-input>
    <sim-bool-input
      class="w-full"
      v-model="storeConfig.currLangConfig['themeConfig']['docFooter']['prev']"
      :input-label="lang('pageProject.settingLang.labels.prevButtonText')"
      :bool-tooltip="lang('pageProject.settingLang.tooltips.prevNextPageInfo')"
      :bool-label="lang('pageProject.settingLang.labels.prevButton')"
      :input-tooltip="
        lang('pageProject.settingLang.tooltips.preButtonTextInfo')
      "
    ></sim-bool-input>
    <sim-bool-input
      class="w-full"
      v-model="storeConfig.currLangConfig['themeConfig']['docFooter']['next']"
      :input-label="lang('pageProject.settingLang.labels.nextButtonText')"
      :bool-tooltip="lang('pageProject.settingLang.tooltips.prevNextPageInfo')"
      :bool-label="lang('pageProject.settingLang.labels.nextButton')"
      :input-tooltip="
        lang('pageProject.settingLang.tooltips.nextButtonTextInfo')
      "
    ></sim-bool-input>
  </div>
  <div class="flex justify-center">
    <a-button
      @click="saveLangConfig()"
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
</template>
<script setup lang="ts">
import DyAddKV from "@/components/dyAddKV.vue";
import SimInput from "@/components/simInput.vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { IconPark } from "@icon-park/vue-next/es/all";
import { onMounted, ref } from "vue";
import { IsEmptyValue } from "@/utils/utils";
import SelectSettingLang from "@/components/selectSettingLang.vue";
import SimBoolInput from "@/components/simBoolInput.vue";
import SimSwitch from "@/components/simSwitch.vue";
import { StringGlobalLang } from "@/configs/cnts";
import { lang } from "@/utils/language";

const storeConfig = useVpconfigStore();
const inputLangArray = ref<any[]>([]);
const refAddLang = ref();
onMounted(() => {
  //获取原始数据
  let langData = storeConfig.configData["locales"];

  let arrData: any = [];
  for (const langDataKey in langData) {
    let item: any = {
      lang: langDataKey,
      label: langData[langDataKey].label,
    };
    arrData.push(item);
  }
  refAddLang.value.setArrayValue(arrData);
});
const saveLangConfig = () => {
  //将数据转换成vitepress所需要的格式
  // let resultData: any = {};
  let resultData = storeConfig.configData["locales"];

  for (let i = 0; i < inputLangArray.value.length; i++) {
    let item = inputLangArray.value[i];
    console.log(item, "item -- console.log");
    if (IsEmptyValue(resultData[item.lang])) {
      resultData[item.lang] = {
        label: item.label,
        lang: item.lang,
      };
    } else {
      resultData[item.lang].label = item.label;
      resultData[item.lang].lang = item.lang;
    }
  }
  storeConfig.configData["locales"] = resultData;
  if (storeConfig.IsUseI18n) {
    if (
      storeConfig.currSettingLang == "" ||
      storeConfig.currSettingLang == StringGlobalLang
    ) {
      storeConfig.currSettingLang = inputLangArray.value[0]["lang"];
    }
  }
  storeConfig.saveConfig();
};
//移除一个元素后 刷新数据
const removeLangItem = (k: string, v: string, removeIndex: number) => {
  delete storeConfig.configData["locales"][k];

  //storeConfig.configData["locales"][k] = undefined;
};
</script>

<style scoped></style>
