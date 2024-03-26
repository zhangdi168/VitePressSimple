<template>
  <dy-add-k-v
    v-if="storeConfig.IsUseI18n"
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
  <select-setting-lang></select-setting-lang>
  <div class="flex justify-start items-center my-2 mx-8">
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['siteTitle']"
      tooltip="即themeConfig.siteTitle 站点的标题，在logo后显示"
      placeholder="请输入站点标题"
      label="站点标题"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['footer']['message']"
      tooltip="由于设计原因，仅当页面不包含侧边栏时才会显示页脚。"
      placeholder="请输入页脚提示信息"
      label="页脚信息"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['footer']['copyright']"
      tooltip="由于设计原因，仅当页面不包含侧边栏时才会显示页脚。"
      placeholder="请输入页脚版权信息"
      label="页脚版权信息"
    ></sim-input>

    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['langMenuLabel']"
      tooltip="自定义导航栏中语言切换按钮的 aria-label，仅当使用 i18n 时才使用此选项签"
      placeholder="请输入项目标题后缀"
      label="语言切换label"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['returnToTopLabel']"
      tooltip="自定义返回顶部按钮的标签，该标签仅在移动端视图中显示"
      placeholder=""
      label="返回顶部按钮的标签"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['sidebarMenuLabel']"
      tooltip="自定义侧边栏菜单标签，该标签仅在移动端视图中显示"
      placeholder=""
      label="自定义侧边栏菜单标签"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['darkModeSwitchLabel']"
      tooltip="自定义深色模式开关标签，该标签仅在移动端视图中显示"
      placeholder=""
      label="深色模式开关标签"
    ></sim-input>
    <sim-input
      v-model="
        storeConfig.currLangConfig['themeConfig']['lightModeSwitchTitle']
      "
      tooltip="自定义[悬停时]显示的浅色模式开关标题"
      placeholder=""
      label="悬停时浅色模式标题"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['darkModeSwitchTitle']"
      tooltip="自定义[悬停时]显示的深色模式开关标题"
      placeholder=""
      label="悬停时深色模式标题"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['editLink']['text']"
      tooltip="编辑链接可让显示链接以编辑 Git 管理服务 (例如 GitHub 或 GitLab) 上的页面"
      placeholder=""
      label="编辑链接文本"
    ></sim-input>
    <sim-input
      v-model="storeConfig.currLangConfig['themeConfig']['editLink']['pattern']"
      tooltip="编辑链接可让显示链接以编辑 Git 管理服务 (例如 GitHub 或 GitLab) 上的页面"
      placeholder=""
      label="编辑链接地址"
    ></sim-input>
    <sim-bool-input
      class="w-full"
      v-model="storeConfig.currLangConfig['themeConfig']['docFooter']['prev']"
      input-label="上一页文本"
      bool-tooltip="用于全局启用或禁用上一页链接"
      bool-label="是否启用上一页连接"
      input-tooltip="用于自定义出现在上一页链接上方的文本"
    ></sim-bool-input>
    <sim-bool-input
      class="w-full"
      v-model="storeConfig.currLangConfig['themeConfig']['docFooter']['next']"
      input-label="下一页文本"
      bool-tooltip="用于全局启用或禁用下一页链接"
      bool-label="是否启用下一页连接"
      input-tooltip="用于自定义出现在下一页链接上方的文本"
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
      保存配置
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
