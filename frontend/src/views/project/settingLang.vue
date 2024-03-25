<template>
  <div class="flex justify-start items-center">
    <sim-switch
      v-model="storeConfig.configData['themeConfig']['i18nRouting']"
      tooltip="将本地语言更改为 zh 会将 URL 从 /foo（或 /en/foo/）更改为 /zh/foo。可以通过将 themeConfig.i18nRouting 设置为 false 来禁用此行为"
      label="启用多语言"
    ></sim-switch>
  </div>

  <dy-add-k-v
    v-if="storeConfig.IsI18nRouting"
    add-btn-text="添加语言类型"
    key-placeholder="语言目录"
    value-placeholder="语言label"
    key-name="lang"
    value-name="label"
    ref="refAddLang"
    @remove-item="removeLangItem"
    v-model:arr="inputLangArray"
  >
  </dy-add-k-v>

  <hr class="my-2" />
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

  <div class="flex justify-center">
    <a-button
      @click="saveLangConfig"
      class="bg-blue-600 hover:bg-blue-500 text-white flex justify-center items-center"
    >
      <icon-park
        class="mr-1"
        strokeLinejoin="bevel"
        theme="outline"
        type="save"
      />
      保存配置
    </a-button>
  </div>
</template>
<script setup lang="ts">
import DyAddKV from "@/components/dyAddKV.vue";
import SimInput from "@/components/simInput.vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { IconPark } from "@icon-park/vue-next/es/all";
import SimSwitch from "@/components/simSwitch.vue";
import { onMounted, ref } from "vue";
import { IsEmptyValue } from "@/utils/utils";

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
  // console.log(inputLangArray.value, "inputLangArray -- console.log");
});
const saveLangConfig = () => {
  //将数据转换成vitepress所需要的格式
  // let resultData: any = {};
  let resultData = storeConfig.configData["locales"];
  for (let i = 0; i < inputLangArray.value.length; i++) {
    let item = inputLangArray.value[i];
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
  // console.log(inputLangArray.value, "inputLangArray.value -- console.log");
  storeConfig.saveConfig();
};
//移除一个元素后 刷新数据
const removeLangItem = (k: string, v: string, removeIndex: number) => {
  delete storeConfig.configData["locales"][k];
  //storeConfig.configData["locales"][k] = undefined;
};
</script>

<style scoped></style>
