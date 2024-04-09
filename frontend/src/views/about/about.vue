<template>
  <div class="flex justify-center">
    <a-radio-group v-model:value="activeKey" style="margin: 8px">
      <a-radio-button value="1">关于</a-radio-button>
      <a-radio-button value="2">隐私协议</a-radio-button>
      <a-radio-button value="3">关于作者</a-radio-button>
    </a-radio-group>
  </div>
  <div v-show="activeKey == '1'">
    <about-software></about-software>
  </div>
  <div v-show="activeKey == '2'">
    <about-privacy></about-privacy>
  </div>
  <div v-show="activeKey == '3'">
    <about-author></about-author>
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import AboutSoftware from "@/views/about/aboutSoftware.vue";
import AboutPrivacy from "@/views/about/aboutPrivacy.vue";
import AboutAuthor from "@/views/about/aboutAuthor.vue";

const activeKey = ref("1");
</script>
