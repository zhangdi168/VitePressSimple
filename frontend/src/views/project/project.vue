<template>
  <div style="overflow-y: auto">
    <div class="flex justify-center items-center">
      <a-radio-group
        animated
        size="large"
        centered
        type="card"
        v-model:value="activeKey1"
        style="margin: 8px"
      >
        <a-radio-button value="1">基础配置</a-radio-button>
        <a-radio-button value="2">导航管理</a-radio-button>
        <a-radio-button value="3">侧栏管理</a-radio-button>
        <a-radio-button value="4">构建相关</a-radio-button>
      </a-radio-group>
    </div>
    <div v-if="activeKey1 === '1'">
      <div class="mt-1">
        <a-tabs
          style="min-height: 55vh"
          v-model:activeKey="activeKey2"
          :tab-position="'left'"
          type="line"
          animated
        >
          <a-tab-pane key="1-1" tab="基础设置">
            <setting-base></setting-base>
          </a-tab-pane>
          <a-tab-pane key="1-2" tab="多语言">
            <setting-lang></setting-lang>
          </a-tab-pane>
          <a-tab-pane key="1-3" tab="路由重写">
            <dy-add-k-v
              add-btn-text="添加路由重写规则"
              key-placeholder="文件路径"
              value-placeholder="路由路径"
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
                保存配置
              </a-button>
            </div>
          </a-tab-pane>
          <a-tab-pane key="1-4" tab="社交帐户">
            <setting-social></setting-social>
          </a-tab-pane>
        </a-tabs>
      </div>
    </div>
    <div v-if="activeKey1 === '2'">
      <div class="mt-1">
        <setting-nav></setting-nav>
      </div>
    </div>
    <div v-if="activeKey1 === '3'">
      <div class="mt-1">
        <setting-sidebar></setting-sidebar>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref } from "vue";
import SettingBase from "@/views/project/settingBase.vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { IconPark } from "@icon-park/vue-next/es/all";
import DyAddKV from "@/components/dyAddKV.vue";
import SettingNav from "@/views/project/settingNav.vue";

import { ToastError, ToastInfo, ToastSuccess } from "@/utils/Toast";
import SimInput from "@/components/simInput.vue";
import SettingLang from "@/views/project/settingLang.vue";
import SettingSocial from "@/views/project/settingSocial.vue";
import SettingSidebar from "@/views/project/settingSidebar.vue";

const activeKey1 = ref<string>("1");
const activeKey2 = ref("1-1");
const storeConfig = useVpconfigStore();
onMounted(() => {
  storeConfig.readVpConfig();
});
//保存配置
const savePageConfig = () => {
  storeConfig.saveConfig();
};
</script>
<style scoped></style>
