<template>
  <div class="flex justify-center">
    <a-radio-group v-model:value="tabPosition" style="margin: 8px">
      <a-radio-button value="top">关于</a-radio-button>
      <a-radio-button value="bottom">隐私协议</a-radio-button>
      <a-radio-button value="left">开源软件</a-radio-button>
    </a-radio-group>
  </div>
  <div>
    <about-software></about-software>
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import { TabsProps } from "ant-design-vue/es/tabs";
import AboutSoftware from "@/views/about/aboutSoftware.vue";

const tabPosition = ref<TabsProps["tabPosition"]>("top");
const activeKey = ref("1");
</script>
