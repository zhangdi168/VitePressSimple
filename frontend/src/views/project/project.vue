<template>
  <div v-if="!storeIndex.IsEmptyProject" style="overflow-y: auto">
    <!--        <div class="flex justify-center select-none items-center">-->
    <!--          <a-radio-group-->
    <!--            animated-->
    <!--            size="large"-->
    <!--            centered-->
    <!--            type="card"-->
    <!--            v-model:value="activeKey1"-->
    <!--            style="margin: 8px"-->
    <!--          >-->
    <!--            <a-radio-button value="1">基础配置</a-radio-button>-->
    <!--            <a-radio-button value="2">导航管理</a-radio-button>-->
    <!--            <a-radio-button value="3">侧栏管理</a-radio-button>-->
    <!--            &lt;!&ndash;        <a-radio-button value="4">构建相关</a-radio-button>&ndash;&gt;-->
    <!--          </a-radio-group>-->
    <!--        </div>-->
    <div v-if="activeKey1 === '1'">
      <div class="mt-1 select-none">
        <a-tabs
          style="min-height: 55vh"
          v-model:activeKey="activeKey2"
          :tab-position="'left'"
          type="line"
          animated
        >
          <a-tab-pane key="1-1" :tab="lang('pageProject.settingBase.title')">
            <setting-base></setting-base>
          </a-tab-pane>
          <a-tab-pane key="1-2" :tab="lang('pageProject.settingLang.title')">
            <setting-lang></setting-lang>
          </a-tab-pane>
          <a-tab-pane key="1-3" :tab="lang('pageProject.settingRewrite.title')">
            <dy-add-k-v
              :add-btn-text="lang('pageProject.settingRewrite.addRewriteRule')"
              :key-placeholder="lang('pageProject.settingRewrite.filePath')"
              :value-placeholder="lang('pageProject.settingRewrite.routePath')"
              v-model:obj="storeConfig.configData['rewrites']"
            ></dy-add-k-v>
            <hr class="my-2" />
            <div class="flex justify-center">
              <a-button
                @click="savePageConfig"
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
          </a-tab-pane>
          <a-tab-pane key="1-4" :tab="lang('pageProject.settingSocial.title')">
            <setting-social></setting-social>
          </a-tab-pane>
          <a-tab-pane key="1-5" :tab="lang('pageProject.settingSearch.title')">
            <setting-search></setting-search>
          </a-tab-pane>
        </a-tabs>
      </div>
    </div>
  </div>
  <empty-project></empty-project>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import SettingBase from "@/views/project/projectSettingBase.vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { IconPark } from "@icon-park/vue-next/es/all";
import DyAddKV from "@/components/dyAddKV.vue";

import SettingLang from "@/views/project/projectSettingLang.vue";
import SettingSocial from "@/views/project/projectSettingSocial.vue";
import { useIndexStore } from "@/store";
import EmptyProject from "@/components/emptyProject.vue";
import SettingSearch from "@/views/project/projectSettingSearch.vue";
import { lang } from "@/utils/language";

const activeKey1 = ref<string>("1");
const activeKey2 = ref("1-1");
const storeConfig = useVpconfigStore();
const storeIndex = useIndexStore();
onMounted(() => {
  storeConfig.readVpConfig();
});
//保存配置
const savePageConfig = () => {
  storeConfig.saveConfig();
};
</script>
<style scoped></style>
