<template>
  <div :style="StyleNoDrag">
    <div class="flex justify-center">
      <a-radio-group v-model:value="activeKey" style="margin: 8px">
        <a-radio-button value="1">{{ lang("pageAbout.title") }}</a-radio-button>
        <a-radio-button value="2">{{
          lang("pageAbout.privacy")
        }}</a-radio-button>
        <a-radio-button value="3">{{
          lang("pageAbout.author")
        }}</a-radio-button>
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
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import AboutSoftware from "@/views/about/aboutSoftware.vue";
import AboutPrivacy from "@/views/about/aboutPrivacy.vue";
import AboutAuthor from "@/views/about/aboutAuthor.vue";
import { lang } from "../../utils/language";
import { StyleNoDrag } from "@/configs/cnts";

const activeKey = ref("1");
</script>
